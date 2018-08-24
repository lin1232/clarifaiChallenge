package search

import (
	"application"
	"webclient"
	"util"
	"strings"
	"fmt"
	"math"
)

const output_num int = 10//the number of output results

//store the result of all images that is tagged, the structure is {url:{tag:probability}}
type Database map[string]map[string]float64


type Tfidf_format struct{
	imgs_concepts_tf map[string]map[string]float64
	concepts_idf map[string]float64
}


func Initialize(file_list_url string, app_id string, model_name string, 
	model_type string)map[string]map[string]float64{
	/*
	Tag all images
	file_list_url : the images that need to be tagged
	app_id : your api key
	model_name: the name of model that you want to use, in this program, only public 
	model is supported
	model_type: the type of model that you want to use
	return: tagged image, the structure is {url:{tag:probability}}
	*/	
	response,_ := webclient.Get(file_list_url, nil)
	//tag all urls
	app := application.NewApplication(app_id)
	model := app.Get(model_name, model_type)
	
	content := strings.Split(string(response), "\n")
	
	res := make(map[string]map[string]float64)
	var data_set []string
	length := len(content)
	for _, value := range content{
		if value != ""{
			data_set = append(data_set, value)
		}
		if len(data_set) == 128{
			length = length - len(data_set)
			fmt.Printf("There are %d images that need to be tagged\n", length)
			util.Merge_map(res, util.Read_predict_res(model.Tag_urls(data_set)))
			data_set = append([]string{})
		}
	}
	
	if len(data_set) != 0{
		util.Merge_map(res, util.Read_predict_res(model.Tag_urls(data_set)))
	}
	
	//util.Save_json(res, "./data.json")
	return res
}

func Convert_tfidf_format(input map[string]map[string]float64) *Tfidf_format{
	/*
	calculate the tf-idf according to tagged images
	input: tagged images
	*/
	tf_idf := new(Tfidf_format)
	tf_idf.concepts_idf = make(map[string]float64)
	tf_idf.imgs_concepts_tf = make(map[string]map[string]float64)
	var length float64 = float64(len(input))
	
	for url,concepts := range input{
		sum := 0.0
		for concept, prob:= range concepts{
			sum += prob
			tf_idf.concepts_idf[concept] += 1
		}
		
		temp := make(map[string]float64)
		tf_idf.imgs_concepts_tf[url] = temp
		for concept, prob := range concepts{
			tf_idf.imgs_concepts_tf[url][concept] = prob/sum
		}
	}
	
	for concept, idf:=range tf_idf.concepts_idf{
		tf_idf.concepts_idf[concept] = math.Log(length/(idf+1.0))
	}
	
	return tf_idf
}

func (tf_idf* Tfidf_format) Search(input []string)[]util.Pair{
	/*
	input: tags
	return: url of images and probability 
	*/
	scores := make([]util.Pair, len(tf_idf.imgs_concepts_tf))
	
	i := 0
	for url, concepts := range tf_idf.imgs_concepts_tf{
		scores[i].Key = url
		for _, concept := range input{
			_, ok := concepts[concept]

			if ok {
				scores[i].Value += 
				tf_idf.imgs_concepts_tf[url][concept]*tf_idf.concepts_idf[concept]
			}
		}
		i++
	}
	
	scores = util.Sort_by_value(scores)
	for i:=0; i<output_num; i++{
		if scores[i].Value == 0.0{
			return scores[:i]
		}
	}
	return scores[:output_num]
}

func (tf *Tfidf_format)Get_imgs_concepts_tf() map[string]map[string]float64{
	return tf.imgs_concepts_tf
}

func (tf *Tfidf_format)Get_concepts_idf() map[string]float64{
	return tf.concepts_idf
}
