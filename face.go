package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Face struct {
	client *Client
}

const (
	FACE_URL   string = "https://api.projectoxford.ai/face/v1.0/"
	DETECT_API string = "detect"
)

func NewFace(key string) *Face {
	if len(key) == 0 {
		return nil
	}

	f := new(Face)
	f.client = NewClient(key)
	return f
}

func getDetectURL(option *DetectParameters) string {
	apiURL := FACE_URL + DETECT_API
	if option == nil {
		return apiURL
	}

	opURL := fmt.Sprintf("%s?returnFaceId=%s&returnFaceLandmarks=%s&returnFaceAttributes=%s",
		getBooleanString(option.RceturnFaceIdcdd),
		getBooleanString(option.PreturnFaceLandmarks),
		option.ReturnFaceAttributes)

	return opURL
}

func getFileByteBuffer(path string) (*bytes.Buffer, error) {
	fileData, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("File open err:", err)
		return nil, err
	}
	return bytes.NewBuffer(fileData), nil
}
func (f *Face) DetectUrl(option *DetectParameters, fileUrl string) (FaceResponse, error) {
	byteData := []byte(fmt.Sprintf(`{"url":"%s"}`, fileUrl))
	data := bytes.NewBuffer(byteData)

	return f.detect(option, data, true)
}

func (f *Face) DetectFile(option *DetectParameters, filePath string) (FaceResponse, error) {
	data, err := getFileByteBuffer(filePath)
	if err != nil {
		return FaceResponse{}, errors.New("File path err:" + err.Error())
	}
	return f.detect(option, data, false)
}

func (f *Face) detect(option *DetectParameters, data *bytes.Buffer, useJson bool) (FaceResponse, error) {

	api := getDetectURL(option)
	body, err := f.client.Connect(api, data, useJson)
	fmt.Println(string(body))
	if err != nil {
		return FaceResponse{}, err
	}

	ret := FaceResponse{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return FaceResponse{}, err
	}
	return ret, nil
}
