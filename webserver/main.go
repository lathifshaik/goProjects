package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w "Post request sucessful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)
	
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w , "$04 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return

	}
	fmt.Fprint(w,"hello")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", formHandler)

	fmt.Printf("String server at port 8080\n")
	if err := http.ListenAndServer(":8080". nil); err !=nil {
		log.Fatal(err)
	}
	
	
}