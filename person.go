package face

import (
	"bytes"
	"encoding/json"
	"log"
)

type Person struct {
	Face
}

func NewPerson(key string) *Person {
	if len(key) == 0 {
		return nil
	}

	f := new(Person)
	f.client = NewClient(key)
	return f
}

func (p *Person) AddFaceByURL(faceUrl, gId, pId, userData, targeFace string) ([]byte, *ErrorResponse) {
	data := getUrlByteBuffer(faceUrl)
	url := getPersonAddURL(gId, pId, userData, targeFace)
	return p.client.Connect("PUT", url, data, true)
}

func (p *Person) AddFaceByPath(path, gId, pId, userData, targeFace string) ([]byte, *ErrorResponse) {
	data, err := getFileByteBuffer(path)
	if err != nil {
		return nil, &ErrorResponse{Err: err}
	}
	url := getPersonAddURL(gId, pId, userData, targeFace)
	return p.client.Connect("PUT", url, data, true)
}

// Create a Person Group with nam and desc
func (p *Person) Create(gId, name, desc string) ([]byte, *ErrorResponse) {
	var option InfoParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, &ErrorResponse{Err: err}
	}

	url := getPersonGidURL(gId)
	return p.client.Connect("PUT", url, bytes.NewBuffer(data), true)

}

// Delte A Person from A PersonGroup
func (p *Person) Delete(gId, pId string) ([]byte, *ErrorResponse) {
	data := bytes.NewBuffer([]byte(""))
	url := getPersonPidURL(gId, pId)
	return p.client.Connect("DELETE", url, data, true)
}

// Delte A Person Face from A PersonGroup/PersonId
func (p *Person) DeleteFace(gId, pId, fId string) ([]byte, *ErrorResponse) {
	data := bytes.NewBuffer([]byte(""))
	url := getPersonFidURL(gId, pId, fId)
	return p.client.Connect("DELETE", url, data, true)
}

// Retrieve a person's information, including registered faces, name and userData.
func (p *Person) Get(gId, pId string) ([]byte, *ErrorResponse) {
	url := getPersonPidURL(gId, pId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)
}

//Retrieve information about a face (specified by face ID, person ID and its belonging person group ID).
func (p *Person) GetFace(gId, pId, fId string) ([]byte, *ErrorResponse) {
	url := getPersonPidURL(gId, pId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)
}

//List all people in a person group, and retrieve person information (including person ID, name, user data and registered faces of the person).
func (p *Person) List(gId string) ([]byte, *ErrorResponse) {
	url := getPersonGidURL(gId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)

}

// Update a person's name or userData field.
func (p *Person) Update(gId, pId, updateName, updateDesc string) ([]byte, *ErrorResponse) {
	var option InfoParameter
	option.Name = updateName
	option.UserData = updateDesc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, &ErrorResponse{Err: err}
	}

	url := getPersonPidURL(gId, pId)
	return p.client.Connect("PATCH", url, bytes.NewBuffer(data), true)
}

// Update a person face's userData field.
func (p *Person) UpdateFace(gId, pId, fId, updateDesc string) ([]byte, *ErrorResponse) {
	data := getUserDataByteBuffer(updateDesc)
	url := getPersonFidURL(gId, pId, fId)
	return p.client.Connect("PATCH", url, data, true)
}
