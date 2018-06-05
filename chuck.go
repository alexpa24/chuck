package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"encoding/json"
	"bufio"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Value struct {
	Id int `json:"id"`
	Joke string `json:"joke"`
	Categories string `json:"categories"`
}

type Recv struct {
	Type string `json:"type"`
	Value Value `json:"value"`
}
	
func request(c chan string){
	var recv Recv
    response, err := http.Get("http://api.icndb.com/jokes/random")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        json.Unmarshal(data, &recv)
    }
	c <- recv.Value.Joke
}

func main() {
	var i = 0
	var alive = true
	jokes := make(map[int]string)
	c := make(chan string)
	f, err := os.Create("./norris.txt")
	check(err)
	defer f.Close()

	for alive == true {
			timer := time.NewTimer(60 * time.Second)
			go func() {
				<- timer.C
				alive = false
			}()

			go request(c)
			jokes[i] = <-c
			i++;
			time.Sleep(3000 * time.Millisecond)
    }	
	fmt.Println("map:", jokes)

	for k, v := range jokes {
		f.WriteString(v)
		f.WriteString("\n")
		k = k
    }
	f.Sync()
	w := bufio.NewWriter(f)
	w.Flush()
}
