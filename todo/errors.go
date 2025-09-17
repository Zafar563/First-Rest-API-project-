package todo

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlredyExist = errors.New("task alredy exist")
