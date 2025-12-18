//go:build linux
// +build linux

package data

import "os"

// CheckChrome 检查 Linux 是否安装了 Chrome/Chromium 浏览器
func CheckChrome() (string, bool) {
	// 检查常见的 Linux 浏览器安装路径
	locations := []string{
		// Chrome
		"/usr/bin/google-chrome",
		"/usr/bin/google-chrome-stable",
		"/opt/google/chrome/google-chrome",
		// Chromium
		"/usr/bin/chromium",
		"/usr/bin/chromium-browser",
		"/snap/bin/chromium",
		"/usr/lib/chromium/chromium",
	}
	for _, location := range locations {
		if _, err := os.Stat(location); err == nil {
			return location, true
		}
	}
	return "", false
}

// CheckBrowser 检查 Linux 是否安装了浏览器，并返回安装路径
func CheckBrowser() (string, bool) {
	if path, ok := CheckChrome(); ok {
		return path, ok
	}
	// 检查 Firefox 作为备选
	firefoxPaths := []string{
		"/usr/bin/firefox",
		"/snap/bin/firefox",
	}
	for _, path := range firefoxPaths {
		if _, err := os.Stat(path); err == nil {
			return path, true
		}
	}
	return "", false
}
