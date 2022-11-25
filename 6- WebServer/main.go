package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "<h1>Hello world</h1>")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}

		if _, ok := req.PostForm["name"]; !ok {
			fmt.Println("Name field is missing")
			fmt.Fprintln(w, "Name field is missing")
			return
		}

		if _, ok := req.PostForm["passion"]; !ok {
			fmt.Println("Passion field is missing")
			fmt.Fprintln(w, "Passion field is missing")
			return
		}

		if _, ok := req.PostForm["surname"]; !ok {
			fmt.Println("Surname field is missing")
			fmt.Fprintln(w, "Surname field is missing")
			return
		}

		if _, ok := req.PostForm["adrr"]; !ok {
			fmt.Println("Adrr field is missing")
			fmt.Fprintln(w, "Adrr field is missing")
			return
		}

		fmt.Println("Name:", req.PostForm["name"][0])
		fmt.Println("Surname:", req.PostForm["surname"][0])
		fmt.Println("Addr:", req.PostForm["adrr"][0])

		fmt.Fprintf(w, "<h1>%s</h1>", req.PostForm["name"][0])
		fmt.Fprintf(w, "<h2>%s</h2>", req.PostForm["surname"][0])
		fmt.Fprintf(w, "<h2>%s</h2>", req.PostForm["passion"][0])
		fmt.Fprintf(w, "<p>%s</p>", req.PostForm["adrr"][0])
	}
	saveFile(w, req)
}

func saveFile(w http.ResponseWriter, req *http.Request) {
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stringreturn := "Name: " + req.PostForm["name"][0] + " Surname: " + req.PostForm["surname"][0] + " Passion: " + req.PostForm["passion"][0] + " Address: " + req.PostForm["adrr"][0]
	fmt.Fprintf(file, "%s\n", stringreturn)
}

func deleteLine(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}

		if _, ok := req.PostForm["name"]; !ok {
			fmt.Println("Name field is missing")
			fmt.Fprintln(w, "Name field is missing")
			return
		}

	default:
		fmt.Println("Method not allowed")
		fmt.Fprintln(w, "Method not allowed")
	}

}

func main() {
	http.HandleFunc("/hello/delete", deleteLine)
	http.HandleFunc("/hello/save", helloHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9000", nil)
}
