package user

import (
	"app/pkg"
	"errors"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"

	"app/common"
	"app/model"
	"app/pkg/storage/pg"
)

func List(filter model.UserFilter, omits ...string) ([]model.User, common.Status, error) {
	var user []model.User
	filter = sanitiseFilter(filter)

	result := pg.Gorm.DB.Omit(omits...).Model(&model.User{}).Find(&user, filter)
	if result.Error != nil {
		return []model.User{}, http.StatusInternalServerError, result.Error
	}

	return user, http.StatusOK, nil
}

func Insert(user model.User) (model.User, common.Status, error) {
	user, err := sanitiseData(user)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	result := pg.Gorm.DB.Model(&model.User{}).Create(&user)
	if result.Error != nil {
		return model.User{}, http.StatusInternalServerError, result.Error
	}

	user.Password = ""
	return user, http.StatusCreated, nil
}

func Update(filter model.UserFilter, update model.User) (model.User, common.Status, error) {
	filter = sanitiseFilter(filter)
	update, err := sanitiseData(update)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	result := pg.Gorm.DB.Model(&model.User{}).Where(filter).Updates(&update).Scan(&update)
	if result.Error != nil {
		return model.User{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return model.User{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	update.Password = ""
	return update, http.StatusOK, nil
}

func Archive(filter model.UserFilter) (model.UserFilter, common.Status, error) {
	filter = sanitiseFilter(filter)
	updates := map[string]interface{}{
		"IsArchive": true,
		"ArchiveAt": gorm.DeletedAt{Time: time.Now(), Valid: true},
	}

	result := pg.Gorm.DB.Model(&model.User{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return model.UserFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return model.UserFilter{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return filter, http.StatusOK, nil
}

func Restore(filter model.UserFilter) (model.UserFilter, common.Status, error) {
	filter = sanitiseFilter(filter)
	updates := map[string]interface{}{
		"IsArchive": false,
		"ArchiveAt": gorm.DeletedAt{},
	}

	result := pg.Gorm.DB.Unscoped().Model(&model.User{}).Where(filter).Updates(updates)
	if result.Error != nil {
		return model.UserFilter{}, http.StatusInternalServerError, result.Error
	}

	if result.RowsAffected == 0 {
		return model.UserFilter{}, http.StatusNotFound, errors.New("can't find any user with this filter")
	}

	return filter, http.StatusOK, nil
}

func sanitiseFilter(filter model.UserFilter) model.UserFilter {
	if len(filter.Email) > 0 {
		filter.Email = strings.ToLower(filter.Email)
	}
	return filter
}

func sanitiseData(data model.User) (model.User, error) {
	if len(data.Email) > 0 {
		data.Email = strings.ToLower(data.Email)
	}

	// Hash password
	if len(data.Password) > 0 {
		hash, err := pkg.HashPassword(data.Password)
		if err != nil {
			return data, err
		}
		data.Password = hash
	}

	return data, nil
}
