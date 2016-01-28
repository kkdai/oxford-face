package face_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/kkdai/oxford-face"
)

var API_KEY string

func init() {
	API_KEY = os.Getenv("MSFT_KEY")
	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
	}
}

func TestFaceDetect(t *testing.T) {
	if API_KEY == "" {
		return
	}

	imageURL := "https://oxfordportal.blob.core.windows.net/face/demov1/verification1-1.jpg"

	client := NewFace(API_KEY)

	res, err := client.DetectUrl(nil, "")
	if err == nil {
		t.Error("Error happen on face detect URL: imput empty")
	}

	res, err = client.DetectUrl(nil, imageURL)
	if err != nil {
		t.Error("Error happen on face detect URL")
	}

	param := DetectParameters{RceturnFaceIdcdd: true, ReturnFaceLandmarks: true, ReturnFaceAttributes: "age,gender"}
	res, err = client.DetectUrl(&param, imageURL)
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	res2, err := client.DetectFile(&param, "test_data/verification1-1.jpg")
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	if res[0].Faceid == "" || res2[0].Faceid == "" {
		t.Error("Error: Not found any face")
		log.Println("Face detect:", res, res2)
	}
}

func TestFaceSimilar(t *testing.T) {
	if API_KEY == "" {
		return
	}

	client := NewFace(API_KEY)
	param := DetectParameters{RceturnFaceIdcdd: true, ReturnFaceLandmarks: true, ReturnFaceAttributes: "age,gender"}

	var faceList []string
	res1, err := client.DetectFile(&param, "test_data/verification1-1.jpg")
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}
	faceList = append(faceList, res1[0].Faceid)
	res2, err := client.DetectFile(&param, "test_data/verification1-2.jpg")
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	faceList = append(faceList, res2[0].Faceid)
	result, err := client.FindSimilarFromList(faceList[0], faceList, 20)
	if err != nil {
		t.Error("Error happen on similar:" + err.Error())
		log.Println("Ret:", result)
	}
}

func TestFaceGroup(t *testing.T) {
	if API_KEY == "" {
		return
	}

}
