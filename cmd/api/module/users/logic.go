package users

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"

	"app/common"
	userData "app/data/user"
	"app/model"
	"app/pkg"
)

type logic struct{}

var Logic logic

func (l *logic) List(filter model.UserFilter) ([]model.User, common.Status, error) {
	results, status, err := userData.List(filter, []string{"password"}...)
	if err != nil {
		return []model.User{}, status, err
	}

	return results, status, nil
}

func (l *logic) Insert(user model.User) (model.User, common.Status, error) {
	result, status, err := userData.Insert(user)
	if err != nil {
		return model.User{}, status, err
	}

	return result, status, nil
}

func (l *logic) Update(filter model.UserFilter, update model.UserUpdate) (model.User, common.Status, error) {
	var user model.User
	err := mapstructure.Decode(update, &user)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}

	result, status, err := userData.Update(filter, user)
	if err != nil {
		return model.User{}, status, err
	}

	return result, status, nil
}

func (l *logic) Archive(filter model.UserFilter) (model.UserFilter, common.Status, error) {
	result, status, err := userData.Archive(filter)
	if err != nil {
		return model.UserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Restore(filter model.UserFilter) (model.UserFilter, common.Status, error) {
	result, status, err := userData.Restore(filter)
	if err != nil {
		return model.UserFilter{}, status, err
	}

	return result, status, nil
}

func (l *logic) Login(payload model.UserLoginPayload) (string, common.Status, error) {
	filter := model.UserFilter{Email: payload.Email}

	results, status, err := userData.List(filter)
	if err != nil {
		return "", status, err
	}

	if len(results) == 0 {
		return "", http.StatusNotFound, errors.New("email or password is wrong")
	}

	user := results[0]
	if !pkg.ComparePassword(user.Password, payload.Password) {
		return "", http.StatusNotFound, errors.New("email or password is wrong")
	}

	payloadClaims := pkg.PayloadClaims{
		ID:         user.ID,
		Permission: user.Permission,
		Platform:   payload.Platform,
		UserType:   pkg.Token.UserType.User,
	}
	token, err := pkg.Token.CreateToken(payloadClaims)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func (l *logic) ForgotPassword(email string) (string, common.Status, error) {
	filter := model.UserFilter{Email: email}
	password := pkg.RandomString(30)
	update := model.UserUpdate{Password: password}

	_, status, err := l.Update(filter, update)
	if err != nil {
		return "", status, err
	}

	body := "<html><body>Your password is changed and your new password is <h3 style=\"display:inline\">%s</strong></body></html>"

	payload := pkg.EmailPayload{
		Recipients: []string{email},
		Body:       fmt.Sprintf(body, password),
		Subject:    "Change Password - Go Ecommerce",
	}
	_, _, err = pkg.SendEmail(payload)
	if err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("your password has been changed but we have a problem on sending email. error: %s", err.Error())
	}

	return "password successfully changed and sent to your email", status, nil
}
