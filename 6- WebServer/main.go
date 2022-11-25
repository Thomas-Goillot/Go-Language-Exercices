package main

import (
	"fmt"
	"net/http"
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

		//check if there is a name surname and adrr field
		if _, ok := req.PostForm["name"]; !ok {
			fmt.Println("Name field is missing")
			fmt.Fprintln(w, "Name field is missing")
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
		fmt.Fprintf(w, "<p>%s</p>", req.PostForm["adrr"][0])
	}
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9000", nil)
}

//nom prenom addresse utilisateur
