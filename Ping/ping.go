package ping

import (
	"fmt"
	"io"
	"net/http"
)

var responseBody string = "Ping!\n"

func Run() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Printf("\nERROR: %v", err)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Ping: got / request\n")
	io.WriteString(w, responseBody)
}
