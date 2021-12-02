package main

import "errors"

type Item struct {
	Number int8
	Active bool
}

var (
	EmptyItem = Item{}
)

func goodCode(itemName string) (Item, error) {
	var isError bool

	if itemName, isError = someFunction1(itemName); !isError {
		return EmptyItem, errors.New("function 1 isn't ok")
	}
	if itemName, isError = someFunction2(itemName); !isError {
		return EmptyItem, errors.New("function 2 isn't ok")
	}

	if itemName, isError = someFunction3(itemName); !isError {
		return EmptyItem, errors.New("function 3 isn't ok")
	}
	if itemName, isError = someFunction4(itemName); !isError {
		return EmptyItem, errors.New("function 4 isn't ok")
	}

	return Item{Number: 1, Active: itemName == ""}, nil
}

func someFunction1(itemName string) (string, bool) {
	return itemName, true
}

func someFunction2(itemName string) (string, bool) {
	return itemName, true
}

func someFunction3(itemName string) (string, bool) {
	return itemName, true
}

func someFunction4(itemName string) (string, bool) {
	return itemName, true
}
