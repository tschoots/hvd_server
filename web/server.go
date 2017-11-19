package main 

import (
	"net/http"
	"fmt"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./html")))
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
		os.Exit(1)
	}

}

