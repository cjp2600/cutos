package config

import (
	"github.com/spf13/viper"
	"strings"
)

func skippedHeaders() []string {
	return viper.GetStringSlice("application_settings.skipped_headers")
}

func skippedHeadersMap() map[string]bool {
	var resp = make(map[string]bool)
	items := skippedHeaders()

	for i := 0; i < len(items); i++ {
		resp[strings.Trim(strings.ToLower(items[i]), " ")] = true
	}
	return resp
}

func IsSkippedHeader(item string) bool {
	item = strings.Trim(strings.ToLower(item), " ")
	hm := skippedHeadersMap()

	// skip auth method
	hm["authorization"] = true

	if _, ok := hm[item]; ok {
		return true
	}
	return false
}
