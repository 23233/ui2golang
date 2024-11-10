package driver

import (
	"fmt"
	"strings"
)

// Dump retrieves the current UI view hierarchy from the device
// Returns:
//   - string: XML representation of the UI hierarchy
//   - error: nil if successful, otherwise error details
func (d *driver) dump() (string, error) {
	if running, _ := d.checkUiAutomator(); !running {
		d.startUiAutomator()
	}

	ip := d.GetIP()
	url := fmt.Sprintf("http://%s:9008/jsonrpc/0", ip)

	res, err := Request(&Requester{
		Url:    url,
		Method: "POST",
		Data: map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"method":  "dumpWindowHierarchy",
			"params":  []interface{}{false, 50},
		},
	})

	if err != nil {
		return "", err
	}

	return strings.TrimSpace((res["result"].(string))), nil
}
