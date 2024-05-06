package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	// to run the random spider web gif on local host 8081
	http.HandleFunc("/", ToRunSpiderWeb)
	http.ListenAndServe(":8081", nil)

}

func ToRunSpiderWeb(gifout http.ResponseWriter, getspider *http.Request) {

	gifout.WriteHeader(http.StatusOK)
	gifout.Header().Set("Content-Type", "image/gif")
	rand.Seed(time.Now().UTC().UnixNano())

	rightfile, wrongfile := ioutil.ReadFile("../webby/random.out.gif")
	if wrongfile != nil {
		panic(wrongfile)
	}
	gifout.Write(rightfile)

}
