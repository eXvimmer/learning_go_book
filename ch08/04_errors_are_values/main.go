package main

import (
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

// A fake login function
func login(uid, pwd string) error {
	return nil
}

// A fake getData function
func getData(file string) ([]byte, error) {
	return []byte{}, nil
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}

	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func main() {

}
