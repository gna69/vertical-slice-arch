package login

import "context"

func (f *Flow) GetUserByLogin(ctx context.Context, login string) (*User, error) {
	return new(User), nil
}
