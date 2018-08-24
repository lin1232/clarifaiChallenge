package webclient

import (
	//"fmt"
    "io/ioutil"
    "net/http"
    "bytes"
    //"reflect"
)

//get data from url
func Get(url string, header map[string]string)([]uint8, error){
	client := &http.Client{}
	
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
        panic(err)
    }
	
	if header != nil{
		for key := range header{
			reqest.Header.Set(key, header[key])
		}
	}
	
    response, _ := client.Do(reqest)
    
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    
    if err != nil {
        panic(err)
        return nil, err
    }    
    
    //fmt.Println(reflect.TypeOf(body))
    return body, err
}

//post data to url
func Post(url string, header map[string]string, param []byte)([]uint8, error){
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	
	if err != nil {
        panic(err)
    }
	
	for key := range header{
		request.Header.Set(key, header[key])
	}
	
	client := &http.Client{}
    response, _ := client.Do(request)
    
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()
    
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
        return nil, err
    }   
    return body, err
    
}
