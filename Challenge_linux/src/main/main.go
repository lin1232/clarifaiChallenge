package main

import (
	"util"
    "fmt"
    "webpage"
    "net/http"
    "encoding/json"
    "log"
    "search"
    "os"
)

func main(){
	fmt.Print("Start initialization\n")
	data := make(map[string]map[string]float64)
	//read images and tags in data.json
	if len(os.Args) == 1 {
		b,_:=util.Read_json("./data.json")
		json.Unmarshal(b, &data)
	}else if len(os.Args) == 2{
		//tags all images
		key := os.Args[1]
		model_name := "general-v1.3"
		model_type := ""
		data = search.Initialize("https://s3.amazonaws.com/clarifai-data/backend/api-take-home/images.txt",
		 "Key "+key, model_name, model_type)
	}else if len(os.Args) == 4{
		//tags all images
		key := os.Args[1]
		model_name := os.Args[2]
		model_type := os.Args[3]
		data = search.Initialize("https://s3.amazonaws.com/clarifai-data/backend/api-take-home/images.txt",
		 "Key "+key, model_name, model_type)
	}else{
		fmt.Print("bad parameter")
		return 
	}
	
	
	var searchAgent = webpage.SearchAgent{Tfidf: search.Convert_tfidf_format(data)}
	fmt.Print("Initializing is done\n")
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/webpage", searchAgent.Webpage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}