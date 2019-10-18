package main

import (
	"flag"
	"fmt"

	. "gitlab.b2w/team/alpha/presal/csv-to-map/internal/app/csv_to_map"
)

func main() {
	pathCsvFile := flag.String("file", "publisher.csv", "Endereço do arquivo CSV para o mapeamento.")
	separator := flag.String("separator", ";", "Endereço do separador do CSV.")
	quote := flag.String("quote", "\"", "Caracteres de quote entre os dados do CSV.")

	flag.Parse()

	err := Execute(*pathCsvFile, *separator, *quote)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mapeamento realizado com sucesso.")
	}
}
