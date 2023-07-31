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

const monitoring = 3
const delay = 5

func main() {

    
	showIntro()

	for {

		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Viewing Logs...")
			printLogs()
		case 0:
			fmt.Println("Leaving the Program...")
			os.Exit(0)
		default:
			fmt.Println("I don't know this command!")
			os.Exit(-1)

		}
	}

}


func showIntro() {

	version := 1.1
	fmt.Println("This program is in version", version)
}

func readCommand() int {

	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Println("The chosen command was", commandRead)
	fmt.Println("")

	return commandRead

}

func displayMenu() {

	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - View Logs")
	fmt.Println("0 - Leave")

}

func startMonitoring() {
	
  fmt.Println("Monitoring...")
	sites := readSitesFromFiles()

	for i := 0; i < monitoring; i++ {

		for i, site := range sites {
			fmt.Println("Testing website", i,
				":", site)
			testSite(site)

		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testSite(site string) {

	resp, err := http.Get(site)

    if err !=nil {

        fmt.Println("An error has occurred:", err)
    }

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, " has been uploaded successfully!")
		registerLog(site,true)
	} else {
		fmt.Println("Site:", site, "it has problems. Status Code:", resp.StatusCode)
		registerLog(site,false)
	}
}

func readSitesFromFiles() []string {
	
  var sites []string
	file, err := os.Open("sites.txt")
    
    if err != nil {
        fmt.Println("An error has occurred:", err)
    }
	
	defer file.Close()
    reader := bufio.NewScanner(file)
    
    for reader.Scan(){ 
        
        line := strings.TrimSpace(reader.Text())
        sites = append(sites, line)

        if err := reader.Err(); err != nil {
			fmt.Println("An error occurred while reading:", err)
        }
    }

    file.Close()
    
	return sites
}

func registerLog (site string, status bool) {
	file, err := os.OpenFile("log.txt",os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err !=nil {
		fmt.Println ("An error has occurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs(){
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	fmt.Println(string(file))
}
