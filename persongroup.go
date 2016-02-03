package face

import (
	"bytes"
	"encoding/json"
	"log"
)

type PersonGroup struct {
	Face
}

func NewPersonGroup(key string) *PersonGroup {
	if len(key) == 0 {
		return nil
	}

	f := new(PersonGroup)
	f.client = NewClient(key)
	return f
}

// A person group is one of the most important parameters for the Face - Identify API.
// The Identify searches person faces in a specified person group.
func (p *PersonGroup) Create(gId, name, desc string) ([]byte, *ErrorResponse) {
	var option InfoParameter
	option.Name = name
	option.UserData = desc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, &ErrorResponse{Err: err}
	}

	url := getPersonGroupGidURL(gId)
	return p.client.Connect("PUT", url, bytes.NewBuffer(data), true)

}

// Delete an existing person group. Persisted face images of all people in the person group will also be deleted.
func (p *PersonGroup) Delete(gId string) ([]byte, *ErrorResponse) {
	data := bytes.NewBuffer([]byte(""))
	url := getPersonGroupGidURL(gId)
	return p.client.Connect("DELETE", url, data, true)
}

// Retrieve the information of a person group, including its name and userData.
// This API returns person group information only, use Person - List Persons in a Person Group instead to retrieve person information under the person group.
func (p *PersonGroup) Get(gId string) ([]byte, *ErrorResponse) {
	url := getPersonGroupGidURL(gId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)
}

// Retrieve the training status of a person group (completed or ongoing). Training can be triggered by the Person Group - Train Person Group API.
// The training will process for a while on the server side..
func (p *PersonGroup) GetTraining(gId string) ([]byte, *ErrorResponse) {
	url := getPersonGroupTrainingGidURL(gId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)
}

// List all person groups and their information.
func (p *PersonGroup) List() ([]byte, *ErrorResponse) {
	url := getPersonGroupURL()
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("GET", url, data, true)
}

// Training is a necessary preparation process of a person group before identify.
// Each person group needs to be trained in order to call Face - Identify.
// The training will process for a while on the server side even after this API has responded.
// You can query the training status by calling Person Group - Get Person Group Training Status API.
func (p *PersonGroup) Train(gId string) ([]byte, *ErrorResponse) {
	url := getPersonGroupTrainURL(gId)
	data := bytes.NewBuffer([]byte(""))

	return p.client.Connect("POST", url, data, true)
}

// Update an existing person group's display name and userData.
// The properties which does not appear in request body will not be updated.
func (p *PersonGroup) Update(gId, pId, updateName, updateDesc string) ([]byte, *ErrorResponse) {
	var option InfoParameter
	option.Name = updateName
	option.UserData = updateDesc

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, &ErrorResponse{Err: err}
	}

	url := getPersonGroupGidURL(gId)
	return p.client.Connect("PATCH", url, bytes.NewBuffer(data), true)
}
