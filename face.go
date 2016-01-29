package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
)

type Face struct {
	client *Client
}

func NewFace(key string) *Face {
	if len(key) == 0 {
		return nil
	}

	f := new(Face)
	f.client = NewClient(key)
	return f
}

func (f *Face) detect(option *DetectParameters, data *bytes.Buffer, useJson bool) ([]byte, error) {
	url := getDetectURL(option)
	return f.client.Connect("POST", url, data, useJson)
}

//Detect face with input URL
func (f *Face) DetectUrl(option *DetectParameters, url string) ([]byte, error) {
	data := getUrlByteBuffer(url)
	return f.detect(option, data, true)
}

//Detect face with input image file path
func (f *Face) DetectFile(option *DetectParameters, filePath string) ([]byte, error) {
	data, err := getFileByteBuffer(filePath)
	if err != nil {
		return nil, errors.New("File path err:" + err.Error())
	}
	return f.detect(option, data, false)
}

func getSimilarData(option SimilarParameter) *bytes.Buffer {
	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil
	}
	return bytes.NewBuffer(data)
}

// Find Face similarity from  a Face List, with max return result to limited return records.
func (f *Face) FindSimilarFromList(targetID string, faceIdList []string, maxResult int) ([]byte, error) {
	var option SimilarParameter
	option.FaceId = targetID
	option.FaceIds = faceIdList
	option.MaxNumOfCandidatesReturned = maxResult

	data := getSimilarData(option)
	api := getSimilarURL()
	if data == nil {
		return nil, errors.New("Invalid parameter")
	}
	return f.client.Connect("POST", api, data, true)
}

// Find Face similarity from  a Face List ID, with max return result to limited return records.
func (f *Face) FindSimilarFromListId(targetID string, listId string, maxResult int) ([]byte, error) {
	var option SimilarParameter
	option.FaceId = targetID
	option.FaceListId = listId
	option.MaxNumOfCandidatesReturned = maxResult

	data := getSimilarData(option)
	api := getSimilarURL()
	if data == nil {
		return nil, errors.New("Invalid parameter")
	}
	return f.client.Connect("POST", api, data, true)
}

// Grouping a slice of faceID to a Face Group
func (f *Face) GroupFaces(faceIDs []string) ([]byte, error) {
	var option GroupParameter
	option.FaceIds = faceIDs
	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getGroupURL()
	return f.client.Connect("POST", url, bytes.NewBuffer(data), true)
}

// Identify a list of face to check belong to which face group
func (f *Face) IdentifyFaces(faceIDs []string, faceGroup string, maxResult int) ([]byte, error) {
	var option IdentifyParameter
	option.FaceIds = faceIDs
	option.PersonGroupId = faceGroup
	option.MaxNumOfCandidatesReturned = maxResult
	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getIdentifyURL()
	return f.client.Connect("POST", url, bytes.NewBuffer(data), true)
}

// Compare input two face id to compute the similarity
func (f *Face) VerifyWithFace(face1 string, face2 string) ([]byte, error) {
	var option VerifyParameter
	option.FaceId1 = face1
	option.FaceId2 = face2

	data, err := json.Marshal(option)
	if err != nil {
		log.Println("Error happen on json marshal:", err)
		return nil, err
	}

	url := getVerifyURL()
	return f.client.Connect("POST", url, bytes.NewBuffer(data), true)
}
