package application

import (
)

const base_url string = "https://api.clarifai.com"
const version string = "v2"
	
type Application struct{
	api Client
	models_cache []map[string]string
	
}

//initialize a new application by app key
func NewApplication(app_key string) *Application{
	var api Client
	api.Set("", "", base_url, app_key, version)
	app:=new(Application)
	app.api = api
	app.models_cache = api.Get_all_models()
	return app
}

//get model by name
func (p_app *Application) Get(model_name string, model_type string) *Model{
	model := new(Model)
	for _,m := range p_app.models_cache{
		if m["name"] == model_name && m["type"] == model_type {
			model.api = p_app.api
			model.model_id = m["id"]
			model.model_name = m["name"]
			model.model_type = m["type"]
			break
		}
	}
	
	return model
}
