package face

import "fmt"

const (
	FACE_URL    string = "https://api.projectoxford.ai/face/v1.0/"
	DETECT_API  string = "detect"
	SIMILAR_API string = "findsimilars"
	GROUP_API   string = "group"
)

func getDetectURL(option *DetectParameters) string {
	apiURL := FACE_URL + DETECT_API
	if option == nil {
		return apiURL
	}

	opURL := fmt.Sprintf("%s?returnFaceId=%s&returnFaceLandmarks=%s&returnFaceAttributes=%s",
		apiURL,
		getBooleanString(option.RceturnFaceIdcdd),
		getBooleanString(option.ReturnFaceLandmarks),
		option.ReturnFaceAttributes)

	return opURL
}

func getSimilarURL() string {
	return FACE_URL + SIMILAR_API
}

func getGroupURL() string {
	return FACE_URL + GROUP_API
}
