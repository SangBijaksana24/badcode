package main

func main() {
	tempBadcode("something")
}

func tempBadcode(itemName string) (Item, error) {
	return Item{Number: 1, Active: true}, nil
}
