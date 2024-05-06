package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// to run the symmetric spider web gif on local host 8080
	http.HandleFunc("/", ToRunSpiderWeb)
	http.ListenAndServe(":8080", nil)

}

func ToRunSpiderWeb(gifout http.ResponseWriter, getspider *http.Request) {

	gifout.WriteHeader(http.StatusOK)
	gifout.Header().Set("Content-Type", "image/gif")
	rand.Seed(time.Now().UTC().UnixNano())

	rightfile, wrongfile := ioutil.ReadFile("../webby/symmetric.out.gif")
	if wrongfile != nil {
		panic(wrongfile)
	}
	gifout.Write(rightfile)

}
