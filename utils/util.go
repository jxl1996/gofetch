package utils

import "net/url"

func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	// 限制协议
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}
