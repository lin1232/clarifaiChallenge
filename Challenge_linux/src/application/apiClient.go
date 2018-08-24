package application

import (
	"webclient"
	//"fmt"
	"encoding/json"
	//"reflect"
	//"strings"
)

type Client struct{
	app_id string
    app_secret string
    base_url string
    app_key string
    version string
}

//you can set the value of the Client
func (p_client* Client) Set(app_id, app_secret, app_url, app_key, version string){
	p_client.app_id = app_id
	p_client.app_secret = app_secret
	p_client.base_url = app_url
	p_client.app_key = app_key
	p_client.version = version
}

//get all models in the application, in this program, only public models is supported
func (p_client* Client) Get_all_models() []map[string]string{
	resource := "models"
	url := p_client.base_url+"/"+p_client.version+"/"+resource
	header := make(map[string]string)
	
	if p_client.app_key != ""{
		header["Authorization"] = p_client.app_key
	}
	
	response, _ := webclient.Get(url, header)

	json_resp := make(map[string]interface{})
	
	err := json.Unmarshal([]byte(response), &json_resp)
	
	if err != nil{
		panic(err)
	}
	
	cache:= json_resp["models"].([]interface{})
	res := make([] map[string]string, 0)
	
	for _,m := range cache{
		mmap := m.(map[string]interface{})
		model := make(map[string]string)
		model["id"] = mmap["id"].(string)
		model["name"] = mmap["name"].(string)
		output_info := mmap["output_info"].(map[string]interface{})
		model["type"] = output_info["type"].(string)
		res = append(res, model)
		if model["name"] == "general-v1.3" && model["type"] == "concept"{
			model["type"] = ""
			res = append(res, model)
		}

	}
	return res
	
}

//predict the tags of images
func (p_client* Client) Predict_model(model_id string, input interface{}) interface{}{
	
	resource := "models"+"/"+model_id+"/"+"outputs"
	url := p_client.base_url+"/"+"v2"+"/"+string(resource)
	
	header := make(map[string]string)
	
	if p_client.app_key != ""{
		header["Authorization"] = p_client.app_key
	}
	data:= make(map[string]interface{})
	data["inputs"] = input
	data["model"] = model_output_info()
	jsonstr,err := json.Marshal(data)
	if err != nil{
		panic(err)
	}
	response,err:= webclient.Post(url, header, jsonstr)
	if err != nil{
		panic(err)
	}

	var json_resp interface{}
	
	err=json.Unmarshal(response, &json_resp)
	
	if err != nil{
		panic(err)
	}
	

	return json_resp
}




