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

func (p *Person) AddPersonFaceByURL(faceUrl, gId, pId, userData, targeFace string) ([]byte, error) {
	data := getUrlByteBuffer(faceUrl)
	url := getPersonAddURL(gId, pId, userData, targeFace)
	return p.client.Connect("PUT", url, data, true)
}

func (p *Person) AddPersonFaceByPath(path, gId, pId, userData, targeFace string) ([]byte, error) {
	data, err := getFileByteBuffer(path)
	if err != nil {
		return nil, err
	}
	url := getPersonAddURL(gId, pId, userData, targeFace)
	return p.client.Connect("PUT", url, data, true)
}

// Create a Person Group with nam and desc
func (p *Person) CreatePerson(gId, name, desc string) ([]byte, error) {
	var option FaceListAddParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getPersonGidURL(gId)
	return p.client.Connect("PUT", url, bytes.NewBuffer(data), true)

}

// Delte A Person from A PersonGroup
func (p *Person) DeletePersonFromGroup(gId, pId string) ([]byte, error) {
	data := bytes.NewBuffer([]byte(""))
	url := getPersonPidURL(gId, pId)
	return p.client.Connect("DELETE", url, data, true)
}

// Delte A Person Face from A PersonGroup/PersonId
func (p *Person) DeleteFaceFromPerson(gId, pId, fId string) ([]byte, error) {
	data := bytes.NewBuffer([]byte(""))
	url := getPersonFidURL(gId, pId, fId)
	return p.client.Connect("DELETE", url, data, true)
}

// Retrieve a person's information, including registered faces, name and userData.
func (f *FaceList) GetPerson(gId, pId string) ([]byte, error) {
	url := getPersonPidURL(gId, pId)
	data := bytes.NewBuffer([]byte(""))

	return f.client.Connect("GET", url, data, true)
}

//Retrieve information about a face (specified by face ID, person ID and its belonging person group ID).
func (f *FaceList) GetPersonFace(gId, pId, fId string) ([]byte, error) {
	url := getPersonPidURL(gId, pId)
	data := bytes.NewBuffer([]byte(""))

	return f.client.Connect("GET", url, data, true)
}

//List all people in a person group, and retrieve person information (including person ID, name, user data and registered faces of the person).
func (p *Person) ListPersonInGroup(gId string) ([]byte, error) {
	url := getPersonsInGidURL(gId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)

}

// Update a person's name or userData field.
func (p *Person) UpdatePersonData(gId, pId, updateName, updateDesc string) ([]byte, error) {
	var option FaceListAddParameter
	option.Name = updateName
	option.UserData = updateDesc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getPersonPidURL(gId, pId)
	return p.client.Connect("PATCH", url, bytes.NewBuffer(data), true)
}

// Update a person face's userData field.
func (p *Person) UpdateFaceData(gId, pId, fId, updateDesc string) ([]byte, error) {
	data := getUserDataByteBuffer(updateDesc)
	url := getPersonFidURL(gId, pId, fId)
	return p.client.Connect("PATCH", url, data, true)
}
