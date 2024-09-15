package apperrors

import "errors"

var (
	NotFound         = errors.New("not_found")
	IllegalOperation = errors.New("illegal_operation")
	InavlidInput     = errors.New("invalid_input")
	Internal         = errors.New("internal")
)
