package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/",fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Servidor iniciado")

	err := http.ListenAndServe(":8084",nil)
	if err !=nil{
		log.Fatal(err)
	}
}
/*
	writter =  es algo que envía el servidor
	request = es algo que pide el usuario
*/
func formHandler(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()

	if err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v" ,err)
		return
	}
	fmt.Fprintf(writer, "POST request successful")
	nombre := request.FormValue("nombre")
	direccion := request.FormValue("direccion")

	fmt.Fprintf(writer, "nombre: %s\n", nombre )
	fmt.Fprintf(writer, "direccion: %s\n", direccion )
}

func helloHandler(writer http.ResponseWriter, request *http.Request){
	if request.URL.Path != "/hello"{
		http.Error(writer, "404 Not Found",http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "This method is nor supported",http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer,"Hola, soy la funcion Hola!")
}

