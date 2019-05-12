package main

import (
	"fmt"
	"html/template"
	"net/http"

	models "./models"
	// "github.com/wolfusflow/goBlog/models"
)

var posts map[string]*models.Post

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>asdasads</h1>")
	//parse the whole folder in future
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Println(posts)

	t.ExecuteTemplate(w, "index", nil)
}

func writeDown(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/writeDown.html",
		"templates/header.html",
		"templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func savePost(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Content = content
	} else {
		id = GenerateId()
		post := models.NewPost(id, title, content)
		post[post.Id] = post
	}

	post := models.NewPost(id, title, content)
	posts[post.Id] = post

	http.Redirect(w, r, "/", 302)
}

func main() {
	fmt.Println("List prt :8000")
	posts = make(map[string]*models.Post, 0)

	//asd/assets/js~css/app.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/writeDown", writeDown)
	http.HandleFunc("/savePost", savePost)
	// http.HandleFunc("/linkFunk", func)
	// http.HandleFunc("/linkFunk", func)

	http.ListenAndServe(":8000", nil)

}
