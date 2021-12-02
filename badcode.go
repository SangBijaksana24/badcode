package main

import "errors"

type Item struct {
	Number int8
	Active bool
}

var (
	EmptyItem = Item{}
)

func badcode(itemName string) (Item, error) {

	if itemName, ok := some_function1(itemName); ok {
		if itemName, ok := some_function2(itemName); ok {

			if itemName, ok := some_function3(itemName); ok {

				if itemName, ok := some_function1(itemName); ok {
					if itemName, ok := some_function2(itemName); ok {

						if itemName, ok := some_function3(itemName); ok {
							return Item{
								Number: 1,
								Active: itemName == "",
							}, nil
						} else {
							return EmptyItem, errors.New("function 3 isn't ok")
						}
					} else {
						return EmptyItem, errors.New("function 2 isn't ok")
					}
				} else {
					return EmptyItem, errors.New("function 1 isn't ok")
				}

			} else {
				return EmptyItem, errors.New("function 3 isn't ok")
			}
		} else {
			return EmptyItem, errors.New("function 2 isn't ok")
		}
	} else {
		return EmptyItem, errors.New("function 1 isn't ok")
	}

	return EmptyItem, errors.New("all function isn't ok")
}

func some_function1(itemName string) (string, bool) {
	return itemName, true
}

func some_function2(itemName string) (string, bool) {
	return itemName, true
}

func some_function3(itemName string) (string, bool) {
	return itemName, true
}

func some_function4(itemName string) (string, bool) {
	return itemName, true
}
