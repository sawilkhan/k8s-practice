package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const PORT = 3000

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/nginx", nginxHandler)
	fmt.Printf("Web server is listening at port %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", PORT), nil)
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil{
		fmt.Printf("Error getting hostname")
	}
	helloMsg := fmt.Sprintf("VERSION 2: Hello from the %s", hostname)
	res, _ := json.Marshal(helloMsg)
	w.WriteHeader(200)
	w.Write(res)
}

func nginxHandler(w http.ResponseWriter, r *http.Request){
	url := "http://nginx"

	res, err := http.Get(url)
	if err != nil{
		fmt.Printf("Failed to connect tp nginx")
	}
	defer res.Body.Close()

	w.WriteHeader(200)
	body, _ := io.ReadAll(res.Body)
	w.Write(body)
}