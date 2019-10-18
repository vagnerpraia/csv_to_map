package load_publisher

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func readCsv(pathFile string, separator string, quote string) ([]string, map[string][]string) {
	csvFile, err := os.Open(pathFile)
	defer csvFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(csvFile)

	var lines []string
	for {

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		lines = append(lines, line)
	}

	var header []string
	flagHeader := true
	csvMap := make(map[string][]string)

	for _, line := range lines {

		re := regexp.MustCompile(`\r?\n`)
		line = re.ReplaceAllString(line, " ")
		line = strings.TrimSpace(line)

		pattern := "(" + quote + "[^" + quote + ",]+),([^" + quote + "]+" + quote + ")"
		re = regexp.MustCompile(pattern)
		stringsCommaRegex := re.FindStringSubmatch(line)

		for _, s := range stringsCommaRegex {

			stringAdjusted := strings.Replace(s, separator, "|||", 1)
			line = strings.Replace(line, s, stringAdjusted, 1)
		}

		if flagHeader {

			header = strings.Split(line, separator)

			for index, data := range header {
				data = strings.Trim(data, quote)
				data = strings.Replace(data, "|||", separator, 1)

				header[index] = data
			}

			flagHeader = false
		} else {

			content := strings.Split(line, separator)

			reKey, err := regexp.Compile("[^0-9]+")
			if err != nil {
				log.Fatal(err)
			}
			key := reKey.ReplaceAllString(content[0], "")

			for _, data := range content {
				data = strings.Trim(data, quote)
				data = strings.Replace(data, "|||", separator, 1)

				csvMap[key] = append(csvMap[key], data)
			}
		}
	}

	return header, csvMap
}
