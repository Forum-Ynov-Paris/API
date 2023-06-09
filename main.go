package main

//replace the package main by the name of your package and delete main function

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Commentaire struct {
	Content string `json:"content"`
	Uuid    string `json:"uuid"`
}

type Article struct {
	Title       string        `json:"title"`
	Tag         string        `json:"tag"`
	Content     string        `json:"content"`
	Upvote      int           `json:"upvote"`
	Date        string        `json:"date"`
	Uuid        string        `json:"uuid"`
	Commentaire []Commentaire `json:"commentaire"`
}

var (
	Path     = "./data.json"
	articles []Article
)

func main() {
	article := Article{
		Title:   "Nouvel article",
		Tag:     "test",
		Content: "Contenu de l'article",
		Upvote:  0,
		Date:    "01/01/2020",
		Uuid:    "123456",
		Commentaire: []Commentaire{
			{Content: "Commentaire 1", Uuid: "qwerty"},
			{Content: "Commentaire 2", Uuid: "asdfgh"},
			{Content: "Commentaire 3", Uuid: "zxcvbn"},
		},
	}
	//Get()
	PostArticle(article)
	//Get()
}

func Get() {
	//open ./data.json and unmarshall it
	file, err := os.Open(Path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//unmarshall the json file
	err = json.NewDecoder(file).Decode(&articles)
	if err != nil {
		log.Fatal(err)
	}
}

func PostArticle(article Article) {
	Get()
	articles = append(articles, article)

	jsonData, err := json.Marshal(articles)
	if err != nil {
		panic(err)
	}

	//open ./data.json and write the new json
	file, err := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
}
