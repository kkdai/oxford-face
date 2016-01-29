package face

import "fmt"

const (
	FACE_URL        string = "https://api.projectoxford.ai/face/v1.0/"
	DETECT_API      string = "detect"
	SIMILAR_API     string = "findsimilars"
	GROUP_API       string = "group"
	IDENTIFY_API    string = "identify"
	VERIFY_API      string = "verify"
	FACELIST_API    string = "facelists"
	PERSONGROUP_API string = "persongroups"
	PERSONS_API     string = "persons"
	TRAINING_API    string = "training"
	TRAIN_API       string = "train"
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

//https://api.projectoxford.ai/face/v1.0/persongroups
func getPersonGroupURL() string {
	return FACE_URL + PERSONGROUP_API
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/
func getPersonGroupGidURL(gId string) string {
	return FACE_URL + PERSONGROUP_API + "/" + gId
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/training
func getPersonGroupTrainingGidURL(gId string) string {
	return FACE_URL + PERSONGROUP_API + "/" + gId + "/" + TRAINING_API
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/train
func getPersonGroupTrainURL(gId string) string {
	return FACE_URL + PERSONGROUP_API + "/" + gId + "/" + TRAIN_API
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/persons
func getPersonGidURL(gId string) string {
	return getPersonGroupGidURL(gId) + "/" + PERSONS_API
}

//https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/persons/{personId}
func getPersonPidURL(gId, pId string) string {
	return getPersonGidURL(gId) + "/persons/" + pId
}

// https://api.projectoxford.ai/face/v1.0/persongroups/{personGroupId}/persons/{personId}/persistedFaces/{persistedFaceId}
func getPersonFidURL(gId, pId, fId string) string {
	return getPersonPidURL(gId, pId) + "/" + PERSONS_API + "/" + fId
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
