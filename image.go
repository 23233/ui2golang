package driver

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	_ "image/jpeg"
	"image/png"
	"os"
	"strings"
)

// CropImage crops a portion of the source image based on the specified rectangle.
// Parameters:
//   - img: The source image to crop from
//   - cropRect: Rectangle defining the crop boundaries
//
// Returns:
//   - The cropped image
func CropImage(img image.Image, cropRect image.Rectangle) image.Image {
	croppedImg := image.NewRGBA(image.Rect(0, 0, cropRect.Dx(), cropRect.Dy()))

	draw.Draw(croppedImg, croppedImg.Bounds(), img, cropRect.Min, draw.Src)

	return croppedImg
}

// LoadImage loads an image from the specified file path.
// Parameters:
//   - path: Path to the image file
//
// Returns:
//   - The loaded image and any error encountered
func LoadImage(path string) (image.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

// Image2Base64 converts an image to a base64 encoded string.
// Parameters:
//   - img: The image to convert
//
// Returns:
//   - The base64 encoded string and any error encountered
func Image2Base64(img image.Image) (string, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// SaveImage saves an image to the specified file path in PNG format.
// Parameters:
//   - img: The image to save
//   - path: Path where the image will be saved
//
// Returns:
//   - Any error encountered during saving
func (d *Driver) SaveImage(img image.Image, path string) error {
	if d.os != "android" {
		if path[0] != '/' {
			path = "/" + path
		}

		filepathParts := strings.Split(path, "/")
		filepath := strings.Join(filepathParts[:len(filepathParts)-1], "/")
		filename := filepathParts[len(filepathParts)-1]
		if !FileExists(TEMP_PATH + filepath) {
			CreateDir(TEMP_PATH + filepath)
		}

		path = TEMP_PATH + filepath + "/" + filename
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

// LoadImage loads an image from the device or local filesystem.
// For non-Android systems, pulls the image from device first.
// Parameters:
//   - path: Path to the image file
//
// Returns:
//   - The loaded image and any error encountered
func (d *Driver) LoadImage(path string) (image.Image, error) {
	if d.os != "android" {
		if path[0] != '/' {
			path = "/" + path
		}

		filepathParts := strings.Split(path, "/")
		filepath := strings.Join(filepathParts[:len(filepathParts)-1], "/")
		filename := filepathParts[len(filepathParts)-1]
		if !FileExists(TEMP_PATH + filepath) {
			CreateDir(TEMP_PATH + filepath)
		}

		d.Run("pull", path, TEMP_PATH+filepath)
		path = TEMP_PATH + filepath + "/" + filename
	}

	return LoadImage(path)
}

// Screenshot captures the current screen of the device.
// Deletes any existing screenshot file before capturing.
//
// Returns:
//   - The captured screenshot as an image
func (d *Driver) Screenshot() image.Image {
	if d.FileExists(IMAGE_PATH) {
		d.DeleteFile(IMAGE_PATH)
		DeleteAll(IMAGE_PATH)
	}

	d.Run("screencap", "-p", IMAGE_PATH)

	img, _ := d.LoadImage(IMAGE_PATH)

	return img
}

// ScreenshotBase64 captures the current screen and returns it as a base64 encoded string.
// It first takes a screenshot using Screenshot() and then converts it to base64 format.
//
// Returns:
//   - string: The base64 encoded screenshot image
//   - error: Any error that occurred during the process
func (d *Driver) ScreenshotBase64() (string, error) {
	img := d.Screenshot()
	return Image2Base64(img)
}

// TODO: Implement FindImage
func FindImage(sourceImage image.Image, targetImage image.Image) (*Bounds, error) {
	return &Bounds{}, nil
}
