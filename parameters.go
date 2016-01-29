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
type FaceListAddParameter struct {
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

type FaceResponse []struct {
	Faceid        string `json:"faceId"`
	Facerectangle struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"faceRectangle"`
	Facelandmarks struct {
		Pupilleft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pupilLeft"`
		Pupilright struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pupilRight"`
		Nosetip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseTip"`
		Mouthleft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"mouthLeft"`
		Mouthright struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"mouthRight"`
		Eyebrowleftouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowLeftOuter"`
		Eyebrowleftinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowLeftInner"`
		Eyeleftouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftOuter"`
		Eyelefttop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftTop"`
		Eyeleftbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftBottom"`
		Eyeleftinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftInner"`
		Eyebrowrightinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowRightInner"`
		Eyebrowrightouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowRightOuter"`
		Eyerightinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightInner"`
		Eyerighttop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightTop"`
		Eyerightbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightBottom"`
		Eyerightouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightOuter"`
		Noserootleft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRootLeft"`
		Noserootright struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRootRight"`
		Noseleftalartop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseLeftAlarTop"`
		Noserightalartop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRightAlarTop"`
		Noseleftalarouttip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseLeftAlarOutTip"`
		Noserightalarouttip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRightAlarOutTip"`
		Upperliptop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"upperLipTop"`
		Upperlipbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"upperLipBottom"`
		Underliptop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"underLipTop"`
		Underlipbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"underLipBottom"`
	} `json:"faceLandmarks"`
	Faceattributes struct {
		Smile    float64 `json:"smile"`
		Headpose struct {
			Pitch float64 `json:"pitch"`
			Roll  float64 `json:"roll"`
			Yaw   float64 `json:"yaw"`
		} `json:"headPose"`
		Gender     string  `json:"gender"`
		Age        float64 `json:"age"`
		Facialhair struct {
			Moustache float64 `json:"moustache"`
			Beard     float64 `json:"beard"`
			Sideburns float64 `json:"sideburns"`
		} `json:"facialHair"`
	} `json:"faceAttributes"`
}

type SimilarResponse []struct {
	Faceid     string  `json:"faceId"`
	Confidence float64 `json:"confidence"`
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
