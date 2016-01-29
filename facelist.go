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

// Add a Face to FaceList by Face URL
func (f *FaceList) AddFaceByURL(faceurl, id, name, faceCurve string) ([]byte, error) {
	data := getUrlByteBuffer(faceurl)
	url := getFacelistAddURL(id, name, faceCurve)
	return f.client.Connect("PUT", url, data, true)
}

// Add a Face to FaceList by Image file path
func (f *FaceList) AddFaceByPath(facePath, id, name, faceCurve string) ([]byte, error) {
	data, err := getFileByteBuffer(facePath)
	if err != nil {
		return nil, err
	}

	url := getFacelistAddURL(id, name, faceCurve)
	return f.client.Connect("PUT", url, data, true)
}

// Create a Face List ID
func (f *FaceList) Create(id, name, desc string) ([]byte, error) {
	var option InfoParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistIdURL(id)
	return f.client.Connect("PUT", url, bytes.NewBuffer(data), true)
}

// Update Face List by ID
func (f *FaceList) Update(id, name, desc string) ([]byte, error) {
	var option InfoParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistIdURL(id)
	return f.client.Connect("PATCH", url, bytes.NewBuffer(data), true)
}

// Delete a Face List by ID
func (f *FaceList) Delete(id string) ([]byte, error) {
	data := bytes.NewBuffer([]byte(""))
	url := getFacelistIdURL(id)
	return f.client.Connect("DELETE", url, data, true)
}

// Delete a Face from Face List
func (f *FaceList) DeleteFace(faceid, listid string) ([]byte, error) {
	var option FaceListDeleteFaceParameter
	option.FaceListId = listid
	option.PersistedFaceId = faceid

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getFacelistIdURL(listid)
	return f.client.Connect("DELETE", url, bytes.NewBuffer(data), true)
}

// Get specific Face list bu Face List ID
func (f *FaceList) Get(id string) ([]byte, error) {
	url := getFacelistIdURL(id)
	data := bytes.NewBuffer([]byte(""))

	return f.client.Connect("GET", url, data, true)
}

// Get all Face list
func (f *FaceList) List() ([]byte, error) {
	url := getFacelistURL()
	data := bytes.NewBuffer([]byte(""))
	return f.client.Connect("GET", url, data, true)
}
