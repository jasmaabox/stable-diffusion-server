package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const DEFAULT_PORT = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("unable to load .env file")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/api/v1/txt2img", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Sending txt2img request to worker...")
		resp, err := http.PostForm(fmt.Sprintf("%s/api/v1/txt2img", os.Getenv("WORKER_URL")), r.PostForm)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Worker sucessfully processed txt2img request.")

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if resp.StatusCode == 200 {
			io.WriteString(w, string(body))
		} else {
			http.Error(w, "an error occured", http.StatusInternalServerError)
			return
		}
	})

	log.Printf("Starting server on %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}
