package common

import "errors"

func AppendError(msg string, err error) error {
	return errors.New(msg + "\n" + err.Error())
}
