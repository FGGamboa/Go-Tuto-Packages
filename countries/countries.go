package countries

import (
	"errors"
	"fmt"
)

var (
	myCountry   string
	collection  []string
	errNotFound error = errors.New("country not found")
)

func Add(country string) {
	collection = append(collection, country)
}

func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errNotFound
	}
	myCountry = country
	return nil
}

func List() {
	for i, c := range collection {
		myCountryMsg := ""

		if c == myCountry {
			myCountryMsg = "[my country]"
		}

		fmt.Printf("%d: %s %s\n", i, c, myCountryMsg)
	}
}

func isInCollection(country string) bool {
	for _, c := range collection {
		if c == country {
			return true
		}
	}
	return false
}
