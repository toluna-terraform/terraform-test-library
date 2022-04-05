package tolunacommons

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/mod/modfile"
)

func GetModName() string {
	modcontent, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalln(err)
	}

	modulename := fmt.Sprintf("%s", modfile.ModulePath(modcontent))
	return string(modulename)
}

func ListContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ListBoolContains(s []bool, str bool) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
