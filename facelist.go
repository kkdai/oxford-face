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

// Create a Face List ID
func (f *FaceList) Create(id, name, desc string) ([]byte, error) {
	var option FaceListAddParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistID(id)
	return f.client.Connect("PUT", url, bytes.NewBuffer(data), true)
}

// Get specific Face list bu Face List ID
func (f *FaceList) GetFaceListByID(id string) ([]byte, error) {
	url := getFacelistID(id)
	data := bytes.NewBuffer([]byte(""))

	return f.client.Connect("GET", url, data, true)
}

// Get all Face list
func (f *FaceList) GetFaceList() ([]byte, error) {
	url := getFacelist()
	data := bytes.NewBuffer([]byte(""))
	return f.client.Connect("GET", url, data, true)
}
