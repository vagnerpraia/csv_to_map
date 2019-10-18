package load_publisher

import "fmt"

func Execute(pathCsvFile string, separator string, quote string) error {
	_, csvMap := readCsv(pathCsvFile, separator, quote)

	fmt.Println(csvMap)

	return nil
}
