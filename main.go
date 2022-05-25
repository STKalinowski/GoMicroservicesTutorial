package main
import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"./handlers"
	"time"
	"context"
	"github.com/gorilla/mux"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	//Handlers
	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)
	
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	
	l.Println("Recieved terminate, graceful shutdown", sig)
	tc,_ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
/*
Epsiode5 Notes:

Using Gorilla framework. 
It should help with making the Mux and handling requests.


*/