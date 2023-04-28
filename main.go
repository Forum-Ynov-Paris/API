package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func main() {
	Get()
	Post()
	Get()
}

func Get() {
	url := "https://forum-ynov-paris.github.io/API/data.json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête :", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer resp.Body.Close()

	var articles []Article
	err = json.NewDecoder(resp.Body).Decode(&articles)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse :", err)
		return
	}

	fmt.Println(articles[0].Tag)
}

func Post() {
	url := "https://forum-ynov-paris.github.io/API/data.json"
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

	articleJson, err := json.Marshal(article)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(articleJson))
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête :", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
