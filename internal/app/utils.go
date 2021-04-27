package app

import (
	"fmt"
	"sort"

	"github.com/freonservice/freon/pkg/api"
)

func uniqueStringSlice(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
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
	var result string
	if len(data) > 0 {
		data = uniqueStringSlice(data)
		sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
		for i, v := range data {
			if i > 0 {
				result = fmt.Sprintf("%s,%s", result, v)
			} else {
				result = v
			}
		}
	}
	return result
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
