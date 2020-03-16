package handlers

import (
	"net/http"
	"io/ioutil"
	"Dnsgslb/pkg/api/types"
	"encoding/json"

)


func Acceptjson(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	str := []byte(body)
	return str
}

func Unmarshalreocrd(res *types.Records, r *http.Request) (*types.Records, error){
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func Unmarshaldomain(res *types.Domains, r *http.Request) (*types.Domains, error){
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func Unmarshalcontent(res *types.ContentType, r *http.Request) (*types.ContentType, error){
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func  Unmarshalmonitor(res *types.Monitors, r *http.Request) (*types.Monitors, error) {
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func Unmarshaltype(res *types.DnsType, r *http.Request) (*types.DnsType, error) {
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil{
		return res, err
	}
	return res, nil
}

func Unmarshalname(res *types.Name, r *http.Request) (*types.Name, error) {
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil{
		return res, err
	}
	return res, nil
}

func Unmarshalview(res *types.Views, r *http.Request) (*types.Views, error) {
	str := Acceptjson(r)
	err := json.Unmarshal(str, &res)
	if err != nil{
		return res, err
	}
	return res, nil
}