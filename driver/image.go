package driver

import (
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

// SaveImage saves an image to the specified file path in PNG format.
// Parameters:
//   - img: The image to save
//   - path: Path where the image will be saved
//
// Returns:
//   - Any error encountered during saving
func (d *driver) SaveImage(img image.Image, path string) error {
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
func (d *driver) LoadImage(path string) (image.Image, error) {
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
func (d *driver) Screenshot() image.Image {
	if d.FileExists(IMAGE_PATH) {
		d.DeleteFile(IMAGE_PATH)
		DeleteAll(IMAGE_PATH)
	}

	d.Run("screencap", "-p", IMAGE_PATH)

	img, _ := d.LoadImage(IMAGE_PATH)

	return img
}
