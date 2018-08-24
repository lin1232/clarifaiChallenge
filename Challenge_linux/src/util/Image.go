package util

import (

)

func Image(url string) map[string]map[string]map[string]string{
	data := make(map[string]map[string]map[string]string)
	data1 := make(map[string]map[string]string)
	data2 := make(map[string]string)
	data2["url"] = url
	data1["image"] = data2
	data["data"]= data1
	return data
}