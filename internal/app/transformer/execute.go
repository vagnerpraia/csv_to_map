package transformer

import (
	"fmt"
	"sort"
	"strings"
)

func Execute(pathFileSource string, formatSource string, pathFileDestiny string, formatDestiny string, separator string, quote string, resultForm string) error {

	var mapSource map[string][]string
	var stringDestiny string

	switch formatSource {
	case "csv":
		_, mapSource = readCsv(pathFileSource, separator, quote)
		break
	}

	var keys []string
	for key := range mapSource {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	switch formatDestiny {
	case "map_go":
		stringDestiny = getMapGo(keys, mapSource)
		break

	case "json":
		stringDestiny = getJson(keys, mapSource)
		break

	}

	if resultForm == "print" {
		fmt.Println(stringDestiny)
	} else if resultForm == "file" {
		write(pathFileDestiny, stringDestiny)
	}

	return nil
}

func getMapGo(keys []string, valuesMap map[string][]string) string {

	result := "map[string]string{\n"

	for _, key := range keys {
		value := valuesMap[key][1]
		result += "\t\"" + key + "\": \"" + value + "\",\n"
	}

	result += "}\n"

	return result
}

func getJson(keys []string, valuesMap map[string][]string) string {

	result := "{\n"

	for _, key := range keys {
		value := valuesMap[key][1]
		result += "\t\"" + key + "\": \"" + value + "\",\n"
	}

	result = strings.TrimRight(result, ",\n")

	result += "\n}\n"

	return result
}
