package internal

import "strconv"

type (
	EmptyFieldError          string
	EmptyInputError          string
	EmptyOrInvalidFieldError string
	EmptyOrInvalidInputError string
	FieldNotFound            string
	NullInputError           string
	NullStructError          string
	UndefinedVersionError    string
)

func (e EmptyFieldError) Error() string {
	return strconv.Quote(string(e)) + " field is empty"
}

func (e EmptyInputError) Error() string {
	return strconv.Quote(string(e)) + " function has received an empty input"
}

func (e EmptyOrInvalidFieldError) Error() string {
	return strconv.Quote(string(e)) + " field is empty or invalid"
}

func (e EmptyOrInvalidInputError) Error() string {
	return strconv.Quote(string(e)) + " function has received an empty or invalid input"
}

func (e FieldNotFound) Error() string {
	return strconv.Quote(string(e)) + " field doesn't exist"
}

func (e NullInputError) Error() string {
	return strconv.Quote(string(e)) + " function has received a null input"
}

func (e NullStructError) Error() string {
	return strconv.Quote(string(e)) + " struct is null"
}

func (e UndefinedVersionError) Error() string {
	return strconv.Quote(string(e)) + " version is undefined"
}
