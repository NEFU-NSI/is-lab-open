package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func ping(host string) bool {
	out, err := exec.Command("ping", "-c", "3", host).Output()
	if err != nil {
		return false
	}
	outputStr := string(out)
	return strings.Contains(outputStr, "3 received")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>")
	if ping("219.217.199.200") {
		fmt.Fprintf(w, "<h1><font color='green'>Host is online</h1>")
	} else {
		fmt.Fprintf(w, "<h1><font color='red'>Host is offline</h1>")
	}
	fmt.Fprintf(w, "</body></html>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on port 8080...\n")
	http.ListenAndServe(":8080", nil)
}
