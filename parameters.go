package face

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

// Delete a Face from Face List parameter
//
type FaceListDeleteFaceParameter struct {
	// Valid character is letter in lower case or digit or '-' or '_', maximum length is 64.
	FaceListId string `json:"faceListId"`
	// Valid character is letter in lower case or digit or '-' or '_', maximum length is 64.
	PersistedFaceId string `json:"persistedFaceId"`
}

// Face List add parameter
//
type InfoParameter struct {
	// Name of the created face list, maximum length is 128.
	Name string `json:"name"`
	// Optional user defined data for the face list. Length should not exceed 16KB.
	UserData string `json:"userData"`
}

//Face verify parameter based on V1 space
//
type VerifyParameter struct {
	//ID of one face.
	FaceId1 string `json:"faceId1"`
	//ID of two face.
	FaceId2 string `json:"faceId2"`
}

//Face Identify parameter based on V1 space
//
type IdentifyParameter struct {
	//A face ID array of candidate faces. Length of faceIds should between [1, 1000].
	// Parameter faceListId and faceIds should not be provided at the same time.
	FaceIds []string `json:"faceIds"`

	//Target person group's ID
	PersonGroupId string `json:"personGroupId"`

	//Optional parameter.
	// Only top maxNumOfCandidatesReturned most similar faces will be returned.
	// maxNumOfCandidatesReturned ranges between [1, 20], default to be 20.
	MaxNumOfCandidatesReturned int `json:"maxNumOfCandidatesReturned"`
}

//Face Group parameter based on V1 space
//
type GroupParameter struct {
	//A face ID array of candidate faces. Length of faceIds should between [1, 1000].
	// Parameter faceListId and faceIds should not be provided at the same time.
	FaceIds []string `json:"faceIds"`
}

//Face Similar parameter based on V1 space
//
type SimilarParameter struct {

	//Query face. The faceId comes from the Face - Detect.
	FaceId string `json:"faceId"`

	//A candidate face list. Face list simply represents a list of faces, reference Face List -
	// Create a Face List for more detail. faceListId and faceIds should not be provided at the same time.
	FaceListId string `json:"faceListId"`

	//A face ID array of candidate faces. Length of faceIds should between [1, 1000].
	// Parameter faceListId and faceIds should not be provided at the same time.
	FaceIds []string `json:"faceIds"`

	//Optional parameter.
	// Only top maxNumOfCandidatesReturned most similar faces will be returned.
	// maxNumOfCandidatesReturned ranges between [1, 20], default to be 20.
	MaxNumOfCandidatesReturned int `json:"maxNumOfCandidatesReturned"`
}

//Face Detect parameter based on V1 space
//
type DetectParameters struct {
	//Return face IDs of the detected faces or not. The default value is true.
	RceturnFaceIdcdd bool

	//Return face landmarks of the detected faces or not. The default value is false.
	ReturnFaceLandmarks bool

	//Analyze and return the one or more specified face attributes in the comma-separated string like "returnFaceAttributes=age,gender".
	// Supported face attributes include age, gender, headPose, smile, and facialHair.
	// Note that each face attribute analysis has additional computational and time cost.
	ReturnFaceAttributes string
}

func getBooleanString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func getFileByteBuffer(path string) (*bytes.Buffer, error) {
	fileData, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("File open err:", err)
		return nil, err
	}
	return bytes.NewBuffer(fileData), nil
}

func getUrlByteBuffer(url string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"url":"%s"}`, url))
	return bytes.NewBuffer(byteData)

}

func getUserDataByteBuffer(userData string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"userData":"%s"}`, userData))
	return bytes.NewBuffer(byteData)

}
