package main

import (
	"bufio"
	"fmt"
	"os"
    "net/http"
) 

func main() {
    fmt.Println("Introduce una dirección objetivo:")
    scanner := bufio.NewReader(os.Stdin)
    target, _ := scanner.ReadString('\n')    
    target = target[:len(target)-1] 

    resp := Ping(target)
    fmt.Printf("is-alive: %t",resp)
}

func Ping(url string) (bool){
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error al hacer la solicitud: %v\n", err)
		return false
	}
	defer resp.Body.Close()
    return true
}
