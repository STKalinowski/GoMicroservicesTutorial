package main
import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"./handlers"
	"time"
	"context"
)

func main(){
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	
	//Handlers
	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	
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
Epsiode4 Notes:
RESTful principles.
REST is useful for microservices.

POST new resources
PUT to update resources


*/