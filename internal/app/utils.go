package app

import (
	"sort"
	"strings"

	api "github.com/freonservice/freon/pkg/freonApi"
)

func uniqueStringSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	var list = make([]string, 0, len(intSlice))
	for _, entry := range intSlice {
		if entry == "" {
			continue
		}
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func createConcatenatedString(data []string) string {
	var f string
	if len(data) > 0 {
		buf := strings.Builder{}
		data = uniqueStringSlice(data)
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		for _, v := range data {
			buf.WriteString(v)
			buf.WriteString(",")
		}
		f = buf.String()
		f = f[:len(f)-1]
	}
	return f
}

func getPlatformByString(platform string) int64 {
	switch platform {
	case "ios":
		return int64(api.PlatformType_PLATFORM_TYPE_IOS)
	case "android":
		return int64(api.PlatformType_PLATFORM_TYPE_ANDROID)
	default:
		return int64(api.PlatformType_PLATFORM_TYPE_WEB)
	}
}

func getStorageTypeByString(storageType string) int64 {
	switch storageType { //nolint:gocritic
	default:
		return int64(api.StorageType_STORAGE_TYPE_LOCAL)
	}
}
