package helpers

import m "pogo/api/models"

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Error(message string) m.Error {
	var error m.Error
	error.Error = true
	error.Message = message

	return error
}
