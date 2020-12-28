package errors

import "errors"

var (
	//Error represents not found any product
	ErrNotFound = errors.New("error not found")
	//Error represents not found any product
	ErrCreateFail = errors.New("error create fail")
)
