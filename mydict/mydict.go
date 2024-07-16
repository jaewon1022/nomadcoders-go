package mydict

import (
	"errors"
)

type Dictionary map[string]string

var errNotFound = errors.New("dict not found")
var errCannotUpdate = errors.New("cannot update non-existing data")
var errCannotDelete = errors.New("cannot delete non-existing data")
var errKeyAlreadyExist = errors.New("you are trying to add value in existing key. you should use update function for change value")

func (dict Dictionary) Search(word string) (string, error) {
	value, exists := dict[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

func(dict Dictionary) Add(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case errNotFound:
		dict[key] = value
		return nil
	case nil:
		return errKeyAlreadyExist
	default:
		return nil
	}
}

func (dict Dictionary) Info() Dictionary {
	return dict
}

func (dict Dictionary) Update(key, value string) error {
	_, err := dict.Search(key)

	switch err {
	case nil:
		dict[key] = value 
		return nil
	case errNotFound:
		return errCannotUpdate
	default:
		return nil
	}
}

func (dict Dictionary) Delete(key string) {
	delete(dict, key)
}