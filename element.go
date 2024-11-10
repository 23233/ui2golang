package driver

import (
	"fmt"
	"strings"
	"time"
)

// element represents a UI element in the Android UI hierarchy
type element struct {
	*document     // embedded document
	x         int // x coordinate of element
	y         int // y coordinate of element
}

// Selector represents a selector type
type Selector string

const (
	Text                  Selector = "text"
	ContentDesc           Selector = "content-desc"
	Class                 Selector = "class"
	ResourceID            Selector = "resource-id"
	StartsWithText        Selector = "starts-with-text"
	EndsWithText          Selector = "ends-with-text"
	StartsWithContentDesc Selector = "starts-with-content-desc"
	EndsWithContentDesc   Selector = "ends-with-content-desc"
	StartsWithClass       Selector = "starts-with-class"
	EndsWithClass         Selector = "ends-with-class"
	StartsWithResourceID  Selector = "starts-with-resource-id"
	EndsWithResourceID    Selector = "ends-with-resource-id"
)

// By represents a selector and its value
type By struct {
	Selector Selector
	Value    string
	Timeout  int
}

// FindElement finds the first UI element matching the XPath expression
// Parameters:
//   - xpath: XPath expression to locate the element
//
// Returns:
//   - *element: Matching element or nil if not found
func (d *document) FindElement(xpath string) *element {
	el := d.root
	if d.element != nil {
		el = d.element
	}

	ele := el.FindElement(xpath)
	if ele == nil {
		return nil
	}

	doc := &document{
		d:       d.d,
		RawXML:  d.RawXML,
		root:    el,
		element: ele,
	}

	e := &element{document: doc}

	bounds := e.GetBounds()
	// Calculate center point
	x := (bounds.LTX + bounds.RBX) / 2
	y := (bounds.LTY + bounds.RBY) / 2

	e.x = x
	e.y = y

	return e
}

// FindElements finds all UI elements matching the XPath expression
// Parameters:
//   - xpath: XPath expression to locate elements
//
// Returns:
//   - []*element: Slice of matching elements or nil if none found
func (d *document) FindElements(xpath string) []*element {
	el := d.root
	if d.element != nil {
		el = d.element
	}

	eles := el.FindElements(xpath)
	if len(eles) == 0 {
		return nil
	}

	es := make([]*element, 0, len(eles))
	for _, ele := range eles {
		doc := &document{
			d:       d.d,
			RawXML:  d.RawXML,
			root:    el,
			element: ele,
		}

		e := &element{document: doc}

		bounds := e.GetBounds()

		x := (bounds.LTX + bounds.RBX) / 2
		y := (bounds.LTY + bounds.RBY) / 2

		e.x = x
		e.y = y

		es = append(es, e)
	}

	return es
}

// ByText finds element by text attribute
func (d *document) ByText(text string) *element {
	return d.FindElement(fmt.Sprintf("//node[@text='%s']", text))
}

// ByContentDesc finds element by content-desc attribute
func (d *document) ByContentDesc(contentDesc string) *element {
	return d.FindElement(fmt.Sprintf("//node[@content-desc='%s']", contentDesc))
}

// ByClass finds element by class attribute
func (d *document) ByClass(className string) *element {
	return d.FindElement(fmt.Sprintf("//node[@class='%s']", className))
}

// ByResourceID finds element by resource-id attribute
func (d *document) ByResourceID(resourceID string) *element {
	return d.FindElement(fmt.Sprintf("//node[@resource-id='%s']", resourceID))
}

// ByStartsWithText finds element by text attribute starting with given value
func (d *document) ByStartsWithText(text string) *element {
	elements := d.FindElements("//node[@text]")
	for _, e := range elements {
		_text := e.Text()
		if _text != "" && strings.HasPrefix(_text, text) {
			return e
		}
	}
	return nil
}

// ByEndsWithText finds element by text attribute ending with given value
func (d *document) ByEndsWithText(text string) *element {
	elements := d.FindElements("//node[@text]")
	for _, e := range elements {
		_text := e.Text()
		if _text != "" && strings.HasSuffix(_text, text) {
			return e
		}
	}
	return nil
}

// ByStartsWithContentDesc finds element by content-desc attribute starting with given value
func (d *document) ByStartsWithContentDesc(contentDesc string) *element {
	elements := d.FindElements("//node[@content-desc]")
	for _, e := range elements {
		_contentDesc := e.ContentDesc()
		if _contentDesc != "" && strings.HasPrefix(_contentDesc, contentDesc) {
			return e
		}
	}
	return nil
}

// ByEndsWithContentDesc finds element by content-desc attribute ending with given value
func (d *document) ByEndsWithContentDesc(contentDesc string) *element {
	elements := d.FindElements("//node[@content-desc]")
	for _, e := range elements {
		_contentDesc := e.ContentDesc()
		if _contentDesc != "" && strings.HasSuffix(_contentDesc, contentDesc) {
			return e
		}
	}
	return nil
}

// ByStartsWithClass finds element by class attribute starting with given value
func (d *document) ByStartsWithClass(className string) *element {
	elements := d.FindElements("//node[@class]")
	for _, e := range elements {
		_className := e.GetAttribute("class")
		if _className != "" && strings.HasPrefix(_className, className) {
			return e
		}
	}
	return nil
}

// ByEndsWithClass finds element by class attribute ending with given value
func (d *document) ByEndsWithClass(className string) *element {
	elements := d.FindElements("//node[@class]")
	for _, e := range elements {
		_className := e.GetAttribute("class")
		if _className != "" && strings.HasSuffix(_className, className) {
			return e
		}
	}
	return nil
}

// ByStartsWithResourceID finds element by resource-id attribute starting with given value
func (d *document) ByStartsWithResourceID(resourceID string) *element {
	elements := d.FindElements("//node[@resource-id]")
	for _, e := range elements {
		_resourceID := e.GetAttribute("resource-id")
		if _resourceID != "" && strings.HasPrefix(_resourceID, resourceID) {
			return e
		}
	}
	return nil
}

// ByEndsWithResourceID finds element by resource-id attribute ending with given value
func (d *document) ByEndsWithResourceID(resourceID string) *element {
	elements := d.FindElements("//node[@resource-id]")
	for _, e := range elements {
		_resourceID := e.GetAttribute("resource-id")
		if _resourceID != "" && strings.HasSuffix(_resourceID, resourceID) {
			return e
		}
	}
	return nil
}

// WaitElement waits for an element to appear on the screen and returns it.
// It polls periodically until the element is found or timeout is reached.
//
// Parameters:
//   - by: Selector configuration containing the search criteria and timeout
//
// Returns:
//   - *element: The found UI element, or nil if not found within timeout
//   - error: ErrSelectorEmpty if selector is empty, ErrElementNotFound if element not found
func (d *driver) WaitElement(by By) (*element, error) {
	if by.Timeout == 0 {
		by.Timeout = WAIT_TIMEOUT
	}

	if by.Selector == "" {
		return nil, ErrSelectorEmpty
	}

	deadline := time.Now().Add(time.Duration(by.Timeout) * time.Millisecond)

	for time.Now().Before(deadline) {
		doc := d.Document()

		var el *element
		switch by.Selector {
		case Text:
			el = doc.ByText(by.Value)
		case ContentDesc:
			el = doc.ByContentDesc(by.Value)
		case Class:
			el = doc.ByClass(by.Value)
		case ResourceID:
			el = doc.ByResourceID(by.Value)
		case StartsWithText:
			el = doc.ByStartsWithText(by.Value)
		case EndsWithText:
			el = doc.ByEndsWithText(by.Value)
		case StartsWithContentDesc:
			el = doc.ByStartsWithContentDesc(by.Value)
		case EndsWithContentDesc:
			el = doc.ByEndsWithContentDesc(by.Value)
		case StartsWithClass:
			el = doc.ByStartsWithClass(by.Value)
		case EndsWithClass:
			el = doc.ByEndsWithClass(by.Value)
		case StartsWithResourceID:
			el = doc.ByStartsWithResourceID(by.Value)
		case EndsWithResourceID:
			el = doc.ByEndsWithResourceID(by.Value)
		}

		if el != nil {
			return el, nil
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil, ErrElementNotFound
}
