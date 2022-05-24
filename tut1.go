package main
import (
	"net/http"
	"log"
	"io/ioutil"
)

func main(){
	//Basic handler
	/*
	HandleFunc 
	default serv monks?
	*/
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			//Still need return because http.Error does not stop flow of request
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye World")
	})

	//Every IP :9090, if specific 127.0.0.1:9090
	http.ListenAndServe(":9090", nil)

}