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

// Update Face List by ID
func (f *FaceList) UpdateFaceListByID(id, name, desc string) ([]byte, error) {
	var option FaceListAddParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistID(id)
	return f.client.Connect("PATCH", url, bytes.NewBuffer(data), true)
}

// Delete a Face List by ID
func (f *FaceList) DeleteFaceListByID(id string) ([]byte, error) {
	data := bytes.NewBuffer([]byte(""))
	url := getFacelistID(id)
	return f.client.Connect("DELETE", url, data, true)
}

// Delete a Face from Face List
func (f *FaceList) DeleteFaceFromListByID(faceid, listid string) ([]byte, error) {
	var option FaceListDeleteFaceParameter
	option.FaceListId = listid
	option.PersistedFaceId = faceid

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistID(listid)
	return f.client.Connect("DELETE", url, bytes.NewBuffer(data), true)
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
