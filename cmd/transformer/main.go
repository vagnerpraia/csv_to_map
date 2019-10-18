package main

import (
	"flag"
	"fmt"

	. "github.com/vagnerpraia/transform_format/internal/app/transformer"
)

func main() {
	pathFileSource := flag.String("fileSource", "bisac.csv", "Endereço do arquivo a ser transformado.")
	formatSource := flag.String("formatSource", "csv", "Formato de origem dos dados. Opções: csv.")
	pathFileDestiny := flag.String("fileDestiny", "bisac.go", "Endereço do arquivo de destino da transformação.")
	formatDestiny := flag.String("formatDestiny", "map_go", "Formato de destino dos dados. Opções: map_go.")
	separator := flag.String("separator", ";", "Endereço do separador do CSV.")
	quote := flag.String("quote", "\"", "Caracteres de quote entre os dados do CSV.")
	resultForm := flag.String("resultForm", "print", "Forma de retorno do resultado. Opções: print / file.")

	flag.Parse()

	err := Execute(*pathFileSource, *formatSource, *pathFileDestiny, *formatDestiny, *separator, *quote, *resultForm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Transformação realizada com sucesso.")
	}
}
