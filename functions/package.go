package functions

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"image/jpeg"
	"io/ioutil"
	"mime/multipart"

	"github.com/disintegration/imaging"
	uuid "github.com/satori/go.uuid"
)

// ResizeImage - Crop the given image and produce one
func ResizeImage(src multipart.File, size image.Point) (image.Image, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}

	defer src.Close()

	dstImage := imaging.Resize(img, size.X, size.Y, imaging.Box)
	return dstImage, nil
}

// CropCenterAnchor - Crop the original image to the size using the center anchor
func CropCenterAnchor(src multipart.File, size image.Point) (image.Image, error) {
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dstImage := imaging.CropAnchor(img, size.X, size.Y, imaging.Center)
	return dstImage, nil
}

// BlurImage - blur image based on opacity value
func BlurImage(src image.Image, opacity float64) (image.Image, error) {
	dstImage128 := imaging.Blur(src, opacity)
	return dstImage128, nil
}

// SaveImage - Save new image locally
func SaveImage(dir string, image image.Image) (string, error) {

	fileName := (dir + uuid.NewV4().String() + ".jpg")
	
	imageBuffer := new(bytes.Buffer)
	if err := jpeg.Encode(imageBuffer, image, nil); err != nil {
		return "", err
	}

	err := ioutil.WriteFile(fileName, imageBuffer.Bytes(), 0666)
	if err != nil {
		return "", err
	}
	
	return fileName, nil
}

// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
func ValidMAC(message, messageMAC, key string) bool {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
	byteMAC, _ := hex.DecodeString(messageMAC)
	return hmac.Equal(byteMAC, expectedMAC)
}

// GenerateHmac - Generates Hmac hex string
func GenerateHmac(message, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}