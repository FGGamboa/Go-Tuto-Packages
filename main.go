package main

import (
	"packages/countries"
)

func main() {
	countries.Add("U.S.A.")
	countries.Add("México")
	countries.Add("Japan")
	countries.Add("Spain")

	err := countries.SetMyCountry("México")
	if err != nil {
		panic(err)
	}

	countries.List()
}
