package face_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/kkdai/oxford-face"
)

var API_KEY string

const (
	// The same man
	//id: 4541fc12-55f2-4eae-b548-2310188fdb8f
	imageURL1_1 string = "https://oxfordportal.blob.core.windows.net/face/demov1/verification1-1.jpg"
	//id: e71470cc-64f3-4cda-91e3-a8e7d9f99d48
	imageURL1_2 string = "https://oxfordportal.blob.core.windows.net/face/demov1/verification1-2.jpg"

	// The same woman
	//id: 075afb27-00e5-465b-b8ee-726369bf2396
	imageURL2_1 string = "https://oxfordportal.blob.core.windows.net/face/demov1/verification3-1.jpg"
	//id: 06b04805-0a97-425d-a12b-36301f91797e
	imageURL2_2 string = "https://oxfordportal.blob.core.windows.net/face/demov1/verification3-2.jpg"
)

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

	client := NewFace(API_KEY)

	_, err := client.DetectUrl(nil, "")
	if err == nil {
		t.Error("Error happen on face detect URL: imput empty")
	}

	_, err = client.DetectUrl(nil, imageURL1_1)
	if err != nil {
		t.Error("Error happen on face detect URL")
	}

	param := DetectParameters{RceturnFaceIdcdd: true, ReturnFaceLandmarks: true, ReturnFaceAttributes: "age,gender"}
	res1, err := client.DetectUrl(&param, imageURL1_1)
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	ret := NewFaceResponse(res1)
	log.Println("Url1-1 detect:", ret)

	res1, err = client.DetectUrl(&param, imageURL1_2)
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	ret = NewFaceResponse(res1)
	log.Println("Url1-2 detect:", ret)

	//	woman
	res1, err = client.DetectUrl(&param, imageURL2_1)
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	ret = NewFaceResponse(res1)
	log.Println("Url2-1 detect:", ret)

	res1, err = client.DetectUrl(&param, imageURL2_2)
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	ret = NewFaceResponse(res1)
	log.Println("Url2-2 detect:", ret)

	res2, err := client.DetectFile(&param, "test_data/verification1-1.jpg")
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}

	ret2 := NewFaceResponse(res2)
	log.Println("File detect:", ret2)

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

	face1 := NewFaceResponse(res1)
	if face1 == nil {
		t.Error("json result failed.")
	}

	faceList = append(faceList, face1[0].Faceid)
	res2, err := client.DetectFile(&param, "test_data/verification1-2.jpg")
	if err != nil {
		t.Error("Error happen on face detect URL with option")
	}
	face2 := NewFaceResponse(res2)
	if face2 == nil {
		t.Error("json result failed.")
	}

	faceList = append(faceList, face2[0].Faceid)
	result, err := client.FindSimilarFromList(faceList[0], faceList, 20)
	if err != nil {
		t.Error("Error happen on similar:" + err.Err.Error())
		log.Println("Ret:", result)
	}
}

func TestFaceGroup(t *testing.T) {
	if API_KEY == "" {
		return
	}

	var faceList []string
	faceList = append(faceList, "4541fc12-55f2-4eae-b548-2310188fdb8f")
	faceList = append(faceList, "e71470cc-64f3-4cda-91e3-a8e7d9f99d48")
	faceList = append(faceList, "075afb27-00e5-465b-b8ee-726369bf2396")
	faceList = append(faceList, "06b04805-0a97-425d-a12b-36301f91797e")

	client := NewFace(API_KEY)
	res, err := client.GroupFaces(faceList)
	if err != nil {
		t.Error("Grouping error:", err)
	}
	log.Println("Grouping:", string(res))
}

func TestFaceVerify(t *testing.T) {
	if API_KEY == "" {
		return
	}

	source := "e71470cc-64f3-4cda-91e3-a8e7d9f99d48"
	target := "4541fc12-55f2-4eae-b548-2310188fdb8f"

	client := NewFace(API_KEY)
	res, err := client.VerifyWithFace(source, target)
	if err != nil {
		t.Error("Verify error:", err)
	}
	log.Println("Verify:", string(res))
}
