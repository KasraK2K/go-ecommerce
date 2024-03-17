package user

import (
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"

	"app/common"
	"app/model"
	"app/pkg/storage/pg"
)

func List(filter model.UserFilter, omits ...string) ([]model.User, common.Status, error) {
	var user []model.User

	result := pg.Gorm.DB.Omit(omits...).Model(&model.User{}).Find(&user, filter)
	if result.Error != nil {
		return []model.User{}, http.StatusInternalServerError, result.Error
	}

	return user, http.StatusOK, nil
}

func Insert(user model.User) (model.User, common.Status, error) {
	result := pg.Gorm.DB.Model(&model.User{}).Create(&user)
	if result.Error != nil {
		return model.User{}, http.StatusInternalServerError, result.Error
	}

	user.Password = ""
	return user, http.StatusOK, nil
}

func Update(filter model.UserFilter, update model.User) (model.User, common.Status, error) {
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
