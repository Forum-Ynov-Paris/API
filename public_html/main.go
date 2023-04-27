package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Title      string       `json:"title"`
	Tag        string       `json:"tag"`
	Content    string       `json:"text"`
	UUID       string       `json:"uuid"`
	Commentary []Commentary `json:"commentaire"`
}

type Commentary struct {
	Content string `json:"text"`
	UUID    string `json:"uuid"`
}

var Data []Message

func Init() {
	Data = []Message{
		{
			Title:   "Hello",
			Tag:     "Golang",
			Content: "Hello World",
			UUID:    "1",
			Commentary: []Commentary{
				{
					Content: "Hello World",
					UUID:    "1",
				},
			},
		},
	}
}

func main() {

	Init()

	http.HandleFunc("/Data", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(Data)
		if err != nil {
			return
		}
	})

	//fmt.Println("localhost:8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
