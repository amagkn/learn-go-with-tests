package maps

import (
	"errors"
)

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = value
		return nil
	case err == nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
