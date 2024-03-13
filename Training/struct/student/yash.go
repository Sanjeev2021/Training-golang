package main

import "errors"

type Yash struct {
	firstName string
	lastName  string
	fullName  string
}

func newYash(firstName, lastName string) (*Yash, error) {
	if firstName == "" {
		return nil, errors.New("Firstname missing")
	}
	if lastName == "" {
		return nil, errors.New("last Name missing")
	}

	fullName := firstName + lastName
	return &Yash{
		firstName: firstName,
		lastName:  lastName,
		fullName:  fullName,
	}, nil
}
