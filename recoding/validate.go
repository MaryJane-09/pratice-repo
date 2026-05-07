package main

import "errors"

func Validate(input string) (rune, error) {
	for _, ascii := range input{
		if ascii < ' ' || ascii > '~' {
			return ascii, errors.New("invalid ascii character")
		}
	}
	return  0, nil
}