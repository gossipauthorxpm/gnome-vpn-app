package errors

import "fmt"

type Error struct {
	errorType Type // todo: Скорее всего типы ошибок будут рендериться где-то в приложении чем то похожим на DI. Singleton readonly переменные.
	message   string
	previous  *Error
	baseError error
}

type Type struct {
	description string
	parent      *Type
}

func CreateError(message string, errorType *Type) *Error {
	return &Error{
		errorType: *errorType,
		message:   message,
		baseError: fmt.Errorf("%s", message),
	}
}

func CreateErrorWithPrevious(message string, errorType *Type, previous *Error) *Error {
	return &Error{
		errorType: *errorType,
		message:   message,
		previous:  previous,
		baseError: fmt.Errorf("%s", message),
	}
}

func CreateErrorType(description string, parent *Type) *Type {
	return &Type{
		description: description,
		parent:      parent,
	}
}

func (self *Error) Error() string {
	return self.message
}

func (self *Error) GetType() Type {
	return self.errorType
}

func (self *Error) GetParantType() *Type {
	return self.errorType.parent
}

func (self *Error) GetPrevious() *Error {
	return self.previous
}

func (self *Error) SetPrevious(previous *Error) {
	self.previous = previous
}

func (self *Error) HasAnyParentType(errorType *Type) bool {
	for value := self.previous; value != nil; value = value.previous {
		if value.GetType() == *errorType {
			return true
		}
	}
	return false
}
