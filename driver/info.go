package driver

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// Retrieves device information and returns it as a JSON string
func (d *driver) Info() string {
	if d.deviceInfo != "" {
		return d.deviceInfo
	}

	deviceInfo := make(map[string]string, 0)

	deviceInfo["model"], _ = d.Run("getprop", "ro.product.model")
	deviceInfo["brand"], _ = d.Run("getprop", "ro.product.brand")
	deviceInfo["market_name"], _ = d.Run("getprop", "ro.product.marketname")
	deviceInfo["android_version"], _ = d.Run("getprop", "ro.build.version.release")
	deviceInfo["sdk_version"], _ = d.Run("getprop", "ro.build.version.sdk")
	deviceInfo["device_id"], _ = d.Run("getprop", "ro.serialno")
	deviceInfo["cpu_platform"], _ = d.Run("getprop", "ro.board.platform")
	deviceInfo["system_device_name"], _ = d.Run("getprop", "persist.sys.device_name")
	deviceInfo["operator_name"], _ = d.Run("getprop", "gsm.sim.operator.alpha")
	deviceInfo["phone_number"], _ = d.Run("getprop", "gsm.sim.operator.numeric")
	deviceInfo["meid"], _ = d.Run("getprop", "ro.ril.oem.meid")
	deviceInfo["system_version"], _ = d.Run("getprop", "ro.build.version.incremental")
	deviceInfo["system_arch"], _ = d.Run("getprop", "ro.product.cpu.abi")
	screen_size, _ := d.Run("wm", "size")
	deviceInfo["screen_size"] = strings.Split(screen_size, ":")[1]
	screen_density, _ := d.Run("wm", "density")
	deviceInfo["screen_density"] = strings.Split(screen_density, ":")[1]
	deviceInfo["imei"] = d.GetIMEI()
	deviceInfo["ip"] = d.GetIP()

	for k, v := range deviceInfo {
		deviceInfo[k] = strings.TrimSpace(v)
	}

	jsonBytes, _ := json.MarshalIndent(deviceInfo, "", " ")

	d.deviceInfo = string(jsonBytes)

	return d.deviceInfo
}

// Retrieves memory information of the device
func (d *driver) MemoryInfo() string {
	output, _ := d.Run("dumpsys", "meminfo")
	return output
}

// Retrieves storage information of the device
func (d *driver) StorageInfo() string {
	output, _ := d.Run("df", "/sdcard", "|", "grep", "'/dev'", "|", "awk", "'{print $5}'")
	return output
}

// Retrieves the IP address of the device
func (d *driver) GetIP() string {
	output, err := d.Run("ip", "-4", "addr", "show", "wlan0")
	if err != nil {
		return "unknown"
	}

	reg := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	ip := reg.FindString(output)

	if ip != "" {
		return ip
	}

	return "localhost"
}

// Retrieves the IMEI of the device
func (d *driver) GetIMEI() string {
	version, _ := d.Run("getprop", "ro.build.version.release")
	v, _ := strconv.Atoi(strings.TrimSpace(version))

	var imei string
	if v >= 12 {
		imei, _ = d.Run("getprop", "ro.ril.oem.imei")
	} else {
		timei, _ := d.Run("service call iphonesubinfo 4 i32 2")
		imei = func(i string) string {
			re := regexp.MustCompile(`'([^']*)'`)
			matches := re.FindAllStringSubmatch(i, -1)

			if len(matches) == 0 {
				return ""
			}

			quoted := ""
			for _, match := range matches {
				quoted += match[1]
			}

			quoted = strings.ReplaceAll(quoted, ".", "")
			quoted = strings.ReplaceAll(quoted, " ", "")

			return quoted
		}(timei)
	}

	return strings.TrimSpace(imei)
}

// Retrieves the screen resolution of the device
func (d *driver) GetResolution() (int, int) {
	screen_size, _ := d.Run("wm", "size")
	screen_size = strings.TrimSpace(strings.Split(screen_size, ":")[1])
	temp := strings.Split(screen_size, "x")

	w, _ := strconv.Atoi(temp[0])
	h, _ := strconv.Atoi(temp[1])

	return w, h
}
