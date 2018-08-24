package webpage

import (
	"fmt"
	"html/template"
	"util"
	"net/http"
	"search"
)

type web struct {
	Title string
	Result []util.Pair
	Key_word []string
}

type SearchAgent struct{
	Tfidf *search.Tfidf_format
}

//get data form html or output data to html
func (searchAgent *SearchAgent)Webpage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.ParseFiles("./webpage.html")
	var web_data web = web{
        Title: "Challenge",
        Key_word:[]string{
        },
        Result: []util.Pair{
        },
    }
	if r.Method == "GET" {
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, web_data)
	} else {
		key_word := r.Form["key_word"][0]
		tags:=util.Split(key_word, "&")
		web_data.Key_word = r.Form["key_word"]
		web_data.Result = searchAgent.Tfidf.Search(tags)
		t.Execute(w, web_data)
	}
}


