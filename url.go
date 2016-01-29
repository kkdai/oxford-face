package face

import "fmt"

const (
	FACE_URL     string = "https://api.projectoxford.ai/face/v1.0/"
	DETECT_API   string = "detect"
	SIMILAR_API  string = "findsimilars"
	GROUP_API    string = "group"
	IDENTIFY_API string = "identify"
	VERIFY_API   string = "verify"
	FACELIST_API string = "facelists"
	PERSON_API   string = "persongroups"
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

func getIdentifyURL() string {
	return FACE_URL + IDENTIFY_API
}

func getVerifyURL() string {
	return FACE_URL + VERIFY_API
}

func getFacelistIdURL(id string) string {
	return FACE_URL + FACELIST_API + "/" + id
}

func getFacelistURL() string {
	return FACE_URL + FACELIST_API
}

func getFacelistAddURL(id, userData, targetFace string) string {
	base := getFacelistIdURL(id)

	opURL := fmt.Sprintf("%s?persistedFaces?userData=%s&targetFace=%s",
		base,
		userData,
		targetFace)

	return opURL
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/
func getPersonGidURL(gId string) string {
	return FACE_URL + PERSON_API + "/" + gId
}

func getPersonsInGidURL(gId string) string {
	return FACE_URL + PERSON_API + "/" + gId + "/persons"
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/persons/{personId}
func getPersonPidURL(gId, pId string) string {
	return getPersonGidURL(gId) + "/persons/" + pId
}

// https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/persons/{personId}/persistedFaces/{persistedFaceId}
func getPersonFidURL(gId, pId, fId string) string {
	return getPersonPidURL(gId, pId) + "/persistedFaces/" + fId
}

// https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/persons/{personId}/persistedFaces[?userData][&targetFace]
func getPersonAddURL(gId, pId, userData, targetFace string) string {
	base := getPersonPidURL(gId, pId)
	opURL := fmt.Sprintf("%s/persistedFaces?userData=%s&targetFace=%s",
		base,
		userData,
		targetFace)

	return opURL
}
