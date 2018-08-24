package util

import (
	"encoding/json"
	//"os"
	"fmt"
	"io/ioutil"
)

type Predict_res struct{
	url string
	concepts map[string]float64
}
func Read_predict_res(response interface{}) map[string]map[string]float64{
	res := make(map[string]map[string]float64)
	
	
	temp1 := response.(map[string]interface{})
	temp2 := temp1["outputs"].([]interface{})
	for _, m := range temp2{
		temp3 := m.(map[string]interface{})
		input:= temp3["input"].(map[string]interface{})
		input_data := input["data"].(map[string]interface{})
		input_img:= input_data["image"].(map[string]interface{})
		input_url:= input_img["url"].(string)

		data := temp3["data"].(map[string]interface{})
		concepts:= data["concepts"].([]interface{})
		temp_res := make(map[string]float64)
		for _, concept := range concepts{
			temp4 := concept.(map[string]interface{})
			temp_res[temp4["name"].(string)] = temp4["value"].(float64)
		}
		res[input_url] = temp_res
	}
	return res
}

func Save_json(data interface{}, filename string) bool{
	json_data, err := json.Marshal(data)
	if err != nil {
        panic(err)
        return false
    }
	
    err = ioutil.WriteFile(filename, json_data, 0644)
    if err != nil {
        panic(err)
        return false
    }
    
    return true
}

func Read_json(filename string) ([]byte, error){
	b, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Print(err)
        return nil, err
    }

    return b, err
}

