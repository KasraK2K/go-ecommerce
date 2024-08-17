package user

import (
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
	user = sanitiseData(user)
	result := pg.Gorm.DB.Model(&model.User{}).Create(&user)
	if result.Error != nil {
		return model.User{}, http.StatusInternalServerError, result.Error
	}

	user.Password = ""
	return user, http.StatusCreated, nil
}

func Update(filter model.UserFilter, update model.User) (model.User, common.Status, error) {
	filter = sanitiseFilter(filter)
	update = sanitiseData(update)
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

func sanitiseData(data model.User) model.User {
	if len(data.Email) > 0 {
		data.Email = strings.ToLower(data.Email)
	}
	return data
}
