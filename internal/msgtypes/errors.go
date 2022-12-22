package msgtypes

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNotFoundKeyUser = errors.New("нет ключа от сообщения в базе пользователя")
var ErrNotModified = errors.New("not modified")
