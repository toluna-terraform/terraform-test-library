/* This package should include general common functions for testing*/
package tolunacommons

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/mod/modfile"
)

// Returns the module name being tested
func GetModName() string {
	modcontent, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalln(err)
	}

	modulename := fmt.Sprintf("%s", modfile.ModulePath(modcontent))
	return string(modulename)
}

// Verifies an Array of type String contains a given String , Returns boolean [true|false]
func ListContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Verifies an Array of type bool contains a given bool [true|false], Returns boolean [true|false]
func ListBoolContains(s []bool, str bool) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
