package helper

import (
	"errors"
	"log"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewError(errstr string) error {
	return errors.New(errstr)
}
