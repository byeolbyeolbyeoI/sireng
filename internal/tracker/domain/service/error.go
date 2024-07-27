package service

import "errors"

var ErrUserAlreadyInSession = errors.New("User is already in an active session")
