package image

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestImage_FromJson(t *testing.T) {
	requestBody := fmt.Sprintf(`{"image_name": "imageobject0_1.jpg", "image_bytes": "%s"}`, TestImg)
	req := events.APIGatewayProxyRequest{
		Body: requestBody,
	}

	expected := &Image{
		ImageName:  "imageobject0_1.jpg",
		ImageBytes: TestImg,
	}

	// Image object to test
	img := &Image{}

	// Call FromJson
	err := img.FromJson(&req)
	log.Println("Image : ", img)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, expected, img)
}

func TestProcessExistingImage(t *testing.T) {
	imagePath := "/Users/ManasSaha/Desktop/My/Programming/Projects/Puc-Detection/ocr-service/pkg/test-images/imageobject0_0.jpg"

	file, err := os.Open(imagePath)
	if err != nil {
		t.Fatalf("Failed to open test image: %v", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		t.Fatalf("Failed to read test image: %v", err)
	}

	imgBase64Str := TestImg

	testImage := Image{
		ImageName:  "new-image.jpg",
		ImageBytes: imgBase64Str,
	}

	err = testImage.DecodeAndSaveImage()
	if err != nil {
		t.Errorf("DecodeAndSaveImage() returned an error: %v", err)
	}

	filePath := "/Users/ManasSaha/Desktop/My/Programming/Projects/Puc-Detection/ocr-service/pkg/test-images/imageobject0_0.jpg"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist", filePath)
	}
}
