package auth

import (
	"context"
)

var Service Auth

type Auth struct{}

func (a *Auth) SignIn(ctx context.Context, login string, password string) (string, error) {

}
