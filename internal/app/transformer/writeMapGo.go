package transformer

import (
	"os"
)

func write(pathFileDestiny string, content string) error {

	f, err := os.Create(pathFileDestiny)
	if err != nil {
		return err
	}

	_, err = f.WriteString(content)
	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
