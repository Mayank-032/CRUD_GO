package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux" // we are using gorilla/mux instead of net/http is as similar as using express js instead of http module of nodejs
)






type Article struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article;







// ResponseWriter is an interface
// http.Request is a struct
func homePage(w http.ResponseWriter, r *http.Request){	// this is the function which will be passed as a landing page
	fmt.Fprintf(w, "Welcome to our homepage")
	fmt.Println("Landing Page EndPoint")
}







func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	
	json.NewEncoder(w).Encode(Articles);	// this will encode my array data into json format and serve it to client as a response
}





func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if(article.ID == key){
			json.NewEncoder(w).Encode(article);
		}
	}
}





func createNewArticle(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)


	json.NewEncoder(w).Encode(article)
}







func handleRequests() {	// As its name suggests, this function is used to handle requests made at particular path, i.e. which function to be executed for a particular path
	// 
	myRouter := mux.NewRouter().StrictSlash(true)
	
	
	
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/all", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/all/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))	
}







func main() {
	Articles = []Article {
		{ID: "1", Title: "Hello1", Desc: "Article_Description", Content: "Article Content"},
		{ID: "2", Title: "Hello2", Desc: "Article_Description", Content: "Article Content"},
		{ID: "3", Title: "Hello3", Desc: "Article_Description", Content: "Article Content"},
		{ID: "4", Title: "Hello4", Desc: "Article_Description", Content: "Article Content"},
		{ID: "5", Title: "Hello5", Desc: "Article_Description", Content: "Article Content"},
		{ID: "6", Title: "Hello6", Desc: "Article_Description", Content: "Article Content"},
	}


	handleRequests()
}