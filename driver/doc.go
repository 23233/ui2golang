package driver

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
)

// documenter represents a UI element in the Android UI hierarchy
type documenter struct {
	d       *driver
	RawXML  string
	root    *xmlquery.Node
	element *xmlquery.Node
	x       int
	y       int
}

// Bounds represents the coordinates of a UI element's bounding box
type Bounds struct {
	LTX int // Left-Top X coordinate
	LTY int // Left-Top Y coordinate
	RBX int // Right-Bottom X coordinate
	RBY int // Right-Bottom Y coordinate
}

// Document returns the UI structure document of current screen
// Returns nil if unable to get UI dump or parse XML
func (d *driver) Document() *documenter {
	xml, err := d.Dump()
	if err != nil {
		return nil
	}

	doc, err := xmlquery.Parse(strings.NewReader(xml))
	if err != nil {
		return nil
	}

	return &documenter{
		d:      d,
		RawXML: xml,
		root:   doc,
	}
}

// FindElement finds first UI element matching the XPath expression
// Parameters:
//   - xpth: XPath expression to locate the element
//
// Returns:
//   - *documenter: matching element or nil if not found
func (d *documenter) FindElement(xpth string) *documenter {
	el := d.root
	if d.element != nil {
		el = d.element
	}

	element := xmlquery.FindOne(el, xpth)
	if element == nil {
		return nil
	}

	doc := &documenter{
		d:       d.d,
		RawXML:  d.RawXML,
		root:    el,
		element: element,
	}

	bounds := doc.GetBounds()
	// get center point
	x := (bounds.LTX + bounds.RBX) / 2
	y := (bounds.LTY + bounds.RBY) / 2

	doc.x = x
	doc.y = y

	return doc
}

// FindElements finds all UI elements matching the XPath expression
// Parameters:
//   - xpth: XPath expression to locate elements
//
// Returns:
//   - []*documenter: slice of matching elements or nil if none found
func (d *documenter) FindElements(xpth string) []*documenter {
	el := d.root
	if d.element != nil {
		el = d.element
	}

	elements := xmlquery.Find(el, xpth)
	if len(elements) == 0 {
		return nil
	}

	docs := make([]*documenter, 0, len(elements))
	for _, element := range elements {
		doc := &documenter{
			d:       d.d,
			RawXML:  d.RawXML,
			root:    el,
			element: element,
		}

		bounds := doc.GetBounds()

		x := (bounds.LTX + bounds.RBX) / 2
		y := (bounds.LTY + bounds.RBY) / 2

		doc.x = x
		doc.y = y

		docs = append(docs, doc)
	}

	return docs
}

// Text returns the text attribute value of the element
func (d *documenter) Text() string {
	if d.element == nil {
		return ""
	}

	return d.element.SelectAttr("text")
}

// ContentDesc returns the content-desc attribute value of the element
func (d *documenter) ContentDesc() string {
	if d.element == nil {
		return ""
	}

	return d.element.SelectAttr("content-desc")
}

// ClassName returns the class attribute value of the element
func (d *documenter) ClassName() string {
	if d.element == nil {
		return ""
	}

	return d.element.SelectAttr("class")
}

// ResourceID returns the resource-id attribute value of the element
func (d *documenter) ResourceID() string {
	if d.element == nil {
		return ""
	}

	return d.element.SelectAttr("resource-id")
}

// Checked returns whether the element is checked
func (d *documenter) Checked() bool {
	if d.element == nil {
		return false
	}

	return d.element.SelectAttr("checked") == "true"
}

// Selected returns whether the element is selected
func (d *documenter) Selected() bool {
	if d.element == nil {
		return false
	}

	return d.element.SelectAttr("enabled") == "true"
}

// Index returns the index attribute value of the element as integer
func (d *documenter) Index() int {
	if d.element == nil {
		return -1
	}

	index := d.element.SelectAttr("index")
	i, _ := strconv.Atoi(index)

	return i
}

// GetBounds returns the element's bounding box coordinates
func (d *documenter) GetBounds() *Bounds {
	if d.element == nil {
		return &Bounds{}
	}

	bounds := d.element.SelectAttr("bounds")

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
func (d *documenter) GetAttribute(name string) string {
	if d.element == nil {
		return ""
	}

	return d.element.SelectAttr(name)
}

// Tap performs a tap action at element's center point
func (d *documenter) Tap() {
	d.d.Tap(d.x, d.y)
}

// LongTap performs a long tap action at element's center point
func (d *documenter) LongTap() {
	d.d.LongTap(d.x, d.y)
}

// Swipe performs a swipe gesture within element's bounds
// Parameters:
//   - direction: swipe direction (SWIPE_UP/DOWN/LEFT/RIGHT)
func (d *documenter) Swipe(direction Direction) {
	bounds := d.GetBounds()

	d.d.SwipeInRange(bounds, direction, 40, 0.8)
}
