package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"strconv"
	// "text/template"
	// "reflect"
)

type Observation struct {
	Treats, Cuddles bool
}

type cat_friend struct {
	Name string
	Age uint64
	Observations Observation
}

var Cat_friends []cat_friend

// func init() {
// 	Cat_friends = []cat_friend{
// 		cat_friend{
// 			Name:"Kenei",
// 			Age: 29,
// 			Observations: Observation{ Treats: true, Cuddles: true },
// 		},
// 		cat_friend{
// 			Name:"Miki",
// 			Age: 11,
// 			Observations: Observation{ Treats: true, Cuddles: false },
// 		},
// 	}
// }

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func page_title(w http.ResponseWriter) {
	title := "Lovely humans - meow meow..."
	fmt.Fprintf(w, "<h1>%s</h1><button><a href=\"/view\">Add friend</a></button>", title)
}

func form(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
	// "<form>"+
	// 	"Name: <input name=\"name\"></input>\n"+
	// 	"Age: <input type=\"number\" name=\"age\"></input>\n"+
	// 	"Treats: <select name=\"treats\">"+
	// 		"<option value=\"true\">Yes</option>"+
	// 		"<option value=\"false\">No</option>"+
	// 	"</select>\n"+
	// 	"Cuddles: <select name=\"cuddles\">"+
	// 		"<option value=\"true\">Yes</option>"+
	// 		"<option value=\"false\">No</option>"+
	// 	"</select>\n"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// "</form>")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page_title(w)
	// form(w)
	fmt.Println(Cat_friends)
	

	for _, friend := range Cat_friends {

		// var treats string
		// if friend.Observations.treats {
		// 	treats = "yessir!"
		// } else {
		// 	treats = "no way jose"
		// }

		// var cuddles string
		// if friend.Observations.cuddles {
		// 	cuddles = "yessir!"
		// } else {
		// 	cuddles = "no way jose"
		// }
		// fmt.Println(friend)
		t, err := template.ParseFiles("/templates/view.html")
		if err != nil {
			panic(err)
		}
		error := t.Execute(w, friend)
		if error != nil {
			panic(error)
		}
	}
}

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("form.html")
// 	t.Execute(w, nil)
// }

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/view", index)
	http.HandleFunc("/process", processor)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")

	if r.Method != "POST" {
		http. Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")
	treats := r.FormValue("treats")
	cuddles := r.FormValue("cuddles")
	
	conv_age, _ := strconv.ParseUint(age, 0, 0)
	conv_treats, _ := strconv.ParseBool(treats)
	conv_cuddles, _ := strconv.ParseBool(cuddles)

	o := Observation{
		Treats: conv_treats,
		Cuddles: conv_cuddles,
	}

	d := cat_friend{
		Name: name,
		Age: conv_age,
		Observations: o,
	}

	// Cat_friends.append(Cat_friends, d)
	// Cat_friends = []cat_friend{
	// 	cat_friend{
	// 		Name:"Ken",
	// 		Age: 3,
	// 		Observations: Observation{ Treats: false, Cuddles: false },
	// 	},
	// }

	Cat_friends = append(Cat_friends, d)

	fmt.Println(Cat_friends)

	http. Redirect(w, r, "/", http.StatusSeeOther)
	// tpl.ExecuteTemplate(w, "index.html", nil)
}

// func (c *Cat_friends) save(new_friend) {
// 	c.append(c, new_friend)
// }