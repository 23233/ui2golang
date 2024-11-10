package driver

import (
	"fmt"
	"image"
	"regexp"
	"strconv"
	"time"

	"github.com/beevik/etree"
)

// document represents the document structure in Android UI hierarchy
type document struct {
	d       *driver        // driver instance
	RawXML  string         // raw XML string
	root    *etree.Element // root XML node
	element *etree.Element // currently selected XML node
}

// Bounds represents the coordinates of a UI element's bounding box
type Bounds struct {
	LTX int // Left-Top X coordinate
	LTY int // Left-Top Y coordinate
	RBX int // Right-Bottom X coordinate
	RBY int // Right-Bottom Y coordinate
}

// Document retrieves and parses the UI hierarchy of the current screen.
// It executes a UI dump command and parses the resulting XML.
//
// Returns:
//   - *document: The parsed UI document structure
//   - nil: If unable to get UI dump or parse the XML
func (d *driver) Document() *document {
	xml, err := d.dump()
	if err != nil {
		return nil
	}

	doc := etree.NewDocument()
	err = doc.ReadFromString(xml)
	if err != nil {
		return nil
	}

	return &document{
		d:      d,
		RawXML: xml,
		root:   &doc.Element,
	}
}

// Text returns the text attribute value of the element
func (d *element) Text() string {
	if d.element == nil {
		return ""
	}

	return d.GetAttribute("text")
}

// ContentDesc returns the content-desc attribute value of the element
func (d *element) ContentDesc() string {
	if d.element == nil {
		return ""
	}

	return d.GetAttribute("content-desc")
}

// ClassName returns the class attribute value of the element
func (d *element) ClassName() string {
	if d.element == nil {
		return ""
	}

	return d.GetAttribute("class")
}

// ResourceID returns the resource-id attribute value of the element
func (d *element) ResourceID() string {
	if d.element == nil {
		return ""
	}

	return d.GetAttribute("resource-id")
}

// Checked returns whether the element is checked
func (d *element) Checked() bool {
	if d.element == nil {
		return false
	}

	return d.GetAttribute("checked") == "true"
}

// Selected returns whether the element is selected
func (d *element) Selected() bool {
	if d.element == nil {
		return false
	}

	return d.GetAttribute("enabled") == "true"
}

// Index returns the index attribute value of the element as integer
func (d *element) Index() int {
	if d.element == nil {
		return -1
	}

	index := d.GetAttribute("index")
	i, _ := strconv.Atoi(index)

	return i
}

// GetBounds returns the element's bounding box coordinates
func (d *element) GetBounds() *Bounds {
	if d.element == nil {
		return nil
	}

	bounds := d.GetAttribute("bounds")

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(bounds, -1)

	rect := make([]int, 4)
	for i := range matches {
		rect[i], _ = strconv.Atoi(matches[i])
	}

	return &Bounds{
		LTX: rect[0],
		LTY: rect[1],
		RBX: rect[2],
		RBY: rect[3],
	}
}

// GetAttribute returns the value of specified attribute
// Parameters:
//   - name: attribute name
func (d *element) GetAttribute(name string) string {
	if d.element == nil {
		return ""
	}

	return d.element.SelectAttr(name).Value
}

// Tap performs a tap action at element's center point
func (d *element) Tap() {
	d.d.Tap(d.x, d.y)
}

// LongTap performs a long tap action at element's center point
func (d *element) LongTap() {
	d.d.LongTap(d.x, d.y)
}

// Swipe performs a swipe gesture within element's bounds
// Parameters:
//   - direction: swipe direction (SWIPE_UP/DOWN/LEFT/RIGHT)
func (d *element) Swipe(direction Direction) {
	bounds := d.GetBounds()

	d.d.swipeInRange(bounds, direction, 40, 0.8)
}

// Input enters text into the element by simulating keyboard input
// Parameters:
//   - text: the text string to input
func (d *element) Input(text string) {
	d.d.Input(d.x, d.y, text)
}

// Clear clears the text content of the current element.
// It simulates clearing text at the element's coordinates.
func (d *element) Clear() {
	d.d.Clear(d.x, d.y)
}

// Search simulates pressing the search key on the element.
// It broadcasts an intent to trigger the search action.
// This is equivalent to pressing the search button on the keyboard.
func (d *element) Search() {
	d.d.Run("am", "broadcast", "-a", "STAR_EDITOR_CODE", "--ei", "code", fmt.Sprintf("%d", IME_ACTION_SEARCH))
}

// Enter simulates pressing the enter key on the element.
// It broadcasts an intent to trigger the done action.
// This is equivalent to pressing the enter/done button on the keyboard.
func (d *element) Enter() {
	d.d.Run("am", "broadcast", "-a", "STAR_EDITOR_CODE", "--ei", "code", fmt.Sprintf("%d", IME_ACTION_DONE))
}

// Next simulates pressing the next key on the element
// to move focus to the next input field
func (d *element) Next() {
	d.d.Run("am", "broadcast", "-a", "STAR_EDITOR_CODE", "--ei", "code", fmt.Sprintf("%d", IME_ACTION_NEXT))
}

// Send simulates pressing the send key on the element
// to submit the current input
func (d *element) Send() {
	d.d.Run("am", "broadcast", "-a", "STAR_EDITOR_CODE", "--ei", "code", fmt.Sprintf("%d", IME_ACTION_SEND))
}

// Previous simulates pressing the previous key on the element
// to move focus to the previous input field
func (d *element) Previous() {
	d.d.Run("am", "broadcast", "-a", "STAR_EDITOR_CODE", "--ei", "code", fmt.Sprintf("%d", IME_ACTION_PREVIOUS))
}

// Go simulates pressing the go key on the element
// to trigger the default action
func (d *element) Go() {
	d.d.Run("am", "broadcast", "-a", "STAR_EDITOR_CODE", "--ei", "code", fmt.Sprintf("%d", IME_ACTION_GO))
}

// Screenshot captures and crops a screenshot of the current element.
// It takes a full device screenshot, crops it to the element's bounds,
// saves the cropped image to a temporary file, and returns the cropped image.
//
// Returns:
//   - image.Image: The cropped screenshot of the element
func (d *element) Screenshot() image.Image {
	bounds := d.GetBounds()

	img := d.d.Screenshot()

	cropImage := CropImage(img, image.Rect(bounds.LTX, bounds.LTY, bounds.RBX, bounds.RBY))

	tempfile := fmt.Sprintf("%s/%v.png", ROOT_PATH, time.Now().UnixMilli())
	d.d.SaveImage(cropImage, tempfile)

	return cropImage
}

// ScreenshotBase64 captures a screenshot of the current element and converts it to base64 string.
// It first takes a screenshot using Screenshot() method, then encodes the image to base64 format.
//
// Returns:
//   - string: Base64 encoded string of the screenshot
//   - error: Error if base64 encoding fails
func (d *element) ScreenshotBase64() (string, error) {
	img := d.Screenshot()
	return Image2Base64(img)
}
