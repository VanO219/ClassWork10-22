package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"strconv"

// 	"github.com/pkg/errors"
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	page, err := template.ParseFiles(
// 		"/home/ivan/go/src/ClassWork/home.html",
// 		"/home/ivan/go/src/ClassWork/base.html",
// 		"/home/ivan/go/src/ClassWork/footer.html")
// 	if err != nil {
// 		err = errors.Wrap(err, "Ошибка при парсинге файла")
// 		fmt.Println(err)
// 		http.Error(w, "Internal Server Error", 500)
// 		return
// 	}
// 	err = page.Execute(w, nil)
// 	if err != nil {
// 		err = errors.Wrap(err, "Ошибка при записи ответа")
// 		fmt.Println(err)
// 		http.Error(w, "Internal Server Error", 500)
// 		return
// 	}
// 	// w.Write([]byte("Hello, World. Welcome to home"))
// }

// func showSnippet(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprintf(w, "Заметка с ID %d", id)
// }

// func createSnippet(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		http.Error(w, "Метод запрещен", 405)
// 		return
// 	}
// 	w.Write([]byte("Create snippet"))
// }
