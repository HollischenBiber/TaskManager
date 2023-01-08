package auth

import "context"

func SignIn(ctx context.Context, authData *authData) (string, error)
