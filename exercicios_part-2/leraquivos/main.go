package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func lerArquivos() ([]string, error) {

	var sites []string
	arquivo, err := os.Open("/Users/pedrobsouza/Documents/repos-jp/learn-go/exercicios_part-2/leraquivos/lista.txt")
	if err != nil {
		return sites, err
	}

	read := bufio.NewReader(arquivo)

	for {
		line, err := read.ReadString('\n') // a função readString fo pacote bufio delemita até onde será lido o byte que deseja
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)

	return sites, nil
}
func main() {
	lerArquivos()
}
