package main

import (
	"fmt"

	"log"

	"net/http"
)

func main() {

	hello := func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("This is version V1 !!!")

		fmt.Fprintln(w, "This is version V1 \n")

	}

	fmt.Println("Hello World!!!")

	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":31003", nil))

}
