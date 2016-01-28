package face

import (
	"bytes"
	"encoding/json"
	"log"
)

type FaceList struct {
	Face
}

func NewFaceList(key string) *FaceList {
	if len(key) == 0 {
		return nil
	}

	f := new(FaceList)
	f.client = NewClient(key)
	return f
}

func (f *FaceList) Add(id, name, desc string) ([]byte, error) {
	var option FaceListAddParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistAdd(id)
	return f.client.Connect("POST", url, bytes.NewBuffer(data), true)
}
