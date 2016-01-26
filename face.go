package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type Face struct {
	SecretKey string
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
	f.SecretKey = key
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
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("File not exist!")
	}
	fw, err := w.CreateFormFile("image", path)
	if err != nil {
		return nil, errors.New("Could not cteate temp files")

	}
	if _, err = io.Copy(fw, f); err != nil {
		return nil, errors.New("Could not copy file content")

	}
	// Add the other fields
	if fw, err = w.CreateFormField("key"); err != nil {
		return nil, errors.New("Could not create form")

	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		return nil, errors.New("Could not prepare key")

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()
	return &b, nil
}

func (f *Face) DetectUrl(option *DetectParameters, fileUrl string) (FaceResponse, error) {
	byteData := []byte(fmt.Sprintf(`{"url":"%s"}`, fileUrl))
	data := bytes.NewBuffer(byteData)

	return f.detect(option, data, false)
}

func (f *Face) DetectFile(option *DetectParameters, filePath string) (FaceResponse, error) {
	data, err := getFileByteBuffer(filePath)
	if err != nil {
		return FaceResponse{}, errors.New("File path err:" + err.Error())
	}
	return f.detect(option, data, true)
}

func (f *Face) detect(option *DetectParameters, fileBuffer *bytes.Buffer, useFile bool) (FaceResponse, error) {

	client := &http.Client{}
	r, _ := http.NewRequest("POST", getDetectURL(option), fileBuffer)
	if useFile {
		r.Header.Add("Content-Type", "application/octet-stream")
	} else {
		r.Header.Add("Content-Type", "application/json")
	}
	//r.Header.Add("Content-Length", strconv.Itoa(len(fileffer)))
	r.Header.Add("Ocp-Apim-Subscription-Key", f.SecretKey)

	resp, _ := client.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println("Error happen! body:", string(body))
		return FaceResponse{}, errors.New("Error on:" + string(body))
	}

	//json unmarshall
	fmt.Println(string(body))
	ret := FaceResponse{}
	err := json.Unmarshal(body, &ret)
	if err != nil {
		return FaceResponse{}, err
	}
	return ret, nil
}
