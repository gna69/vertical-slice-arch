package login

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"vertical-slice-arch/internal/pkg/postgres"
)

type (
	Flow struct {
		postgresClient *postgres.Client
	}

	User struct {
		Credentians Credentials
		Profile     Profile
	}

	Credentials struct {
		Login    string
		Password string
	}

	Profile struct {
		Username string
		Age      uint8
	}
)

var (
	ErrorInvalidCredentials = errors.New("invalid login or password")
)

func NewFlow(postgresClient *postgres.Client) *Flow {
	return &Flow{
		postgresClient: postgresClient,
	}
}

func (f *Flow) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	/*
		unmarshal body to dto
		validate input
		convert dto to entities
	*/

	if err := f.login(ctx, "login_from_body", "password_from_body"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (f *Flow) login(ctx context.Context, login, password string) error {
	gotCredentials := Credentials{
		Login:    login,
		Password: password,
	}

	user, err := f.GetUserByLogin(ctx, login)
	if err != nil {
		return fmt.Errorf("failed to get user by login from storage: %w", err)
	}

	if !f.isCredentialsEqual(gotCredentials, user.Credentians) {
		return ErrorInvalidCredentials
	}

	return nil
}

func (f *Flow) isCredentialsEqual(firstCreds, secondCreds Credentials) bool {
	if firstCreds.Login != secondCreds.Login {
		return false
	}

	if firstCreds.Password != secondCreds.Password {
		return false
	}

	return true
}
