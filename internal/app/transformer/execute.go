package transformer

import (
	"fmt"
	"sort"
)

func Execute(pathFileSource string, formatSource string, pathFileDestiny string, formatDestiny string, separator string, quote string, resultForm string) error {

	var mapSource map[string][]string
	var stringDestiny string

	if formatSource == "csv" {
		_, mapSource = readCsv(pathFileSource, separator, quote)
	}

	var keys []string
	for key := range mapSource {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	if formatDestiny == "map_go" {
		stringDestiny += "map[string]string{\n"

		for _, key := range keys {
			value := mapSource[key][1]
			stringDestiny += "\t\"" + key + "\": \"" + value + "\",\n"
		}

		stringDestiny += "}\n"
	}

	if resultForm == "print" {
		fmt.Println(stringDestiny)
	} else if resultForm == "file" {
		write(pathFileDestiny, stringDestiny)
	}

	return nil
}
