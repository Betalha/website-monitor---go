package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5
const numTests = 5

func main() {
	for {
		initiate()
		comand := scanInt()
		switch comand {
		case 1:
			monitorando()
		case 2:
			fmt.Println("exibindo logs...")
			printLogs()
		case 0:
			fmt.Println("saindo...")
			os.Exit(0)
		default:
			fmt.Println("comando não existe")
			os.Exit(-1)
		}
	}
}
func initiate() {
	var coms = [3]string{"0 - sair do programa", "1 - Iniciar monnitoramento", "2 - Exibir logs"}
	for i := 0; i < 3; i++ {
		fmt.Println(coms[i])
	}
}
func scanInt() int {
	var i int
	fmt.Scan(&i)

	return i
}
func monitorando() {
	fmt.Println("monitorando...")
	pages := readFile("./pages.txt")
	for i := 0; i < numTests; i++ {
		fmt.Println("teste", (i + 1))
		for _, page := range pages {
			resp, err := http.Get(page)

			if err != nil {
				fmt.Println("error:", err)
				return
			}
			if resp.StatusCode == 200 {
				fmt.Println("A pagina:", page, "foi carregada com sucesso.")
				logsFile(page, true)
			} else {
				fmt.Println("A pagina:", page, "esta com o erro: ", resp.StatusCode)
				logsFile(page, false)
			}
		}
		if i != (numTests - 1) {
			time.Sleep(delay * time.Second)
		}
	}
	fmt.Println("")
}
func readFile(pagesFile string) []string {
	var pages []string

	file, err := os.Open(pagesFile)
	if err != nil {
		fmt.Println("error:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		pages = append(pages, line)

		if err == io.EOF {
			break
		}
	}
	file.Close()
	return pages
}
func logsFile(page string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("error", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + page + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
func printLogs() {
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(file))
}
