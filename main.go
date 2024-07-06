package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
    "time"
) 

func main() {
    //mux := http.NewServeMux()
    fmt.Println("Introduce una dirección objetivo:")
    scanner := bufio.NewReader(os.Stdin)
    target, _ := scanner.ReadString('\n')    
    isAlive := pingURL2(target)
    fmt.Printf("Funciona %v\n",isAlive)
}

func pingURL(target string) bool {
    client := http.Client{
        Timeout: 5 * time.Second,
    }
    resp, err := client.Get(target)
    if err != nil {
        return false
    }
    defer resp.Body.Close()
    return resp.StatusCode == http.StatusOK
}

func pingURL2(url string) (bool) {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return true 
}
