Project Oxford Face API for Golang
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/oxford-face/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/oxford-face?status.svg)](https://godoc.org/github.com/kkdai/oxford-face)  [![Build Status](https://travis-ci.org/kkdai/oxford-face.svg)](https://travis-ci.org/kkdai/oxford-face)
 
![](https://www.projectoxford.ai/images/bright/face/FaceAPI-Main.png)
 
This package is [Project Oxford](https://www.projectoxford.ai/) [Face API](https://www.projectoxford.ai/face) in Golang

What is Project Oxford Face API
---------------

[Project Oxford](https://www.projectoxford.ai/) is a web services from Microsoft. It contains following services. (refer from this [page](https://www.projectoxford.ai/face).)

#### Face Detection
![](https://www.projectoxford.ai/images/bright/face/FaceDetection.png)  

To detect human faces in image with face rectangles and face attributes including face landmarks, pose, and machine learning-based predictions of gender and age.

#### Face Verification
![](https://www.projectoxford.ai/images/bright/face/FaceVerification.png)

To check two faces belonging to same person or not, with confidence score.

##### Similar Face Searching
![](https://www.projectoxford.ai/images/bright/face/SimilarFaceSearching.png)

To find similar-looking faces from many faces by a query face.

#### Face Grouping
![](https://www.projectoxford.ai/images/bright/face/FaceGrouping.png)

To organize many faces into face groups based on their visual similarity.

#### Face Identification
![](https://www.projectoxford.ai/images/bright/face/FaceIdentification.png)

To search which specific person entity a query face belongs to, from user-provided person-face data.

Installation
---------------
```
go get github.com/kkdai/oxford-face
```

How to use it
---------------

Sign-up for Microsoft Translator API (see [here](http://blogs.msdn.com/b/translation/p/gettingstarted1.aspx) for more detail) and get your developer credentials. Use the client ID and secret to instantiate a translator as shown below.

```go
package main

import (
	"fmt"
	"os"

	. "github.com/kkdai/oxford-face"
)

func main() {
	key := os.Getenv("MSFT_KEY")
	if key == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
		return
	}
	f := NewFace(key)

	//Detect
	ret, err := f.DetectFile(nil, "./1.jpg")
	fmt.Println("ret:", ret, " err=", err)
}
```

Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`

Inspired
---------------

- [Project Oxford: Face API](https://www.projectoxford.ai/face)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

