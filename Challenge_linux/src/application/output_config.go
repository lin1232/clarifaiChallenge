package application

import (

)

func model_output_info()interface{}{
	data := make(map[string]map[string]map[string]bool)
	data1 := make(map[string]map[string]bool)
	data2 := make(map[string]bool)
	data2["concepts_mutually_exclusive"] = false
	data2["closed_environment"] = false
	data1["output_config"] = data2
	data["output_info"]= data1
	return data
	
}