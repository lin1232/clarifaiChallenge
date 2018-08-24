package application

import (
	//"webClient"
	"util"
	"fmt"
)

type Model struct{
	model_id string
	model_name string
	model_type string
	api Client
}

func (p_model *Model) Pridict_url(url string) interface{}{
	//fmt.Printf("%s%s", p_model.model_id, p_model.model_name)
	input:= []interface{} {util.Image(url)}
	response := p_model.api.Predict_model(p_model.model_id, input)
	//fmt.Print(response)
	return response
}

//tag images by url, the maximum number of url is 128
func (p_model *Model) Tag_urls(urls []string) interface{}{
	if len(urls) > 128{
		fmt.Println("error: too many url")
		return "error"
	}
	var input []interface{}
	for _, url := range urls{
		input = append(input, util.Image(url))
	}
	response := p_model.api.Predict_model(p_model.model_id, input)
	//fmt.Print(response, "\n")
	return response
}

//initialize a new model
func New_model(model_id string, model_name string, 
	api Client, model_type string) *Model{
	p_model := new(Model)
	p_model.model_id = model_id
	p_model.model_name = model_name
	p_model.api = api
	p_model.model_type = model_type
	return p_model
}

/*
func (p_model *Model) Set(model_id string, model_version map[string]interface{}, model_name string, 
	api Client, model_type string, model_output_info map[string]interface{}){
	p_model.model_id = model_id
	p_model.model_version = model_version
	p_model.model_name = model_name
	p_model.api = api
	p_model.model_type = model_type
	p_model.model_output_info = model_output_info
}
*/