package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

    
	exibeIntroducao()

	for {

		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando!")
			os.Exit(-1)

		}
	}

}


func exibeIntroducao() {

	nome := "Robson"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {

	var comandoLido int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")

	return comandoLido

}

func exibeMenu() {

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

}

func iniciarMonitoramento() {
	
  fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site", i,
				":", site)
			testaSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string) {

	resp, err := http.Get(site)

    if err !=nil {

        fmt.Println("Ocorreu um erro:", err)
    }

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " foi carregado com sucesso!")
		registraLog(site,true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site,false)
	}
}

func leSitesDoArquivo() []string {
	
  var sites []string
	arquivo, err := os.Open("sites.txt")
    
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
	
	defer arquivo.Close() // Usamos defer para garantir que o arquivo seja fechado ao final da função.
    leitor := bufio.NewScanner(arquivo)
    
    for leitor.Scan(){ // Continuamos lendo enquanto houver linhas no arquivo.
        
        linha := strings.TrimSpace(leitor.Text())
        sites = append(sites, linha)

        if err := leitor.Err(); err != nil {
			fmt.Println("Ocorreu um erro durante a leitura:", err)
        }
    }

    arquivo.Close()
    
	return sites
}

func registraLog (site string, status bool) {
	arquivo, err := os.OpenFile("log.txt",os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err !=nil {
		fmt.Println ("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs(){
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}