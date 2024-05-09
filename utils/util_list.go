package utils

// InStringList 检查ket是否在String类型的数组中
func InStringList(key string, list []string) bool {
	for _, value := range list {
		if key == value {
			return true
		}
	}
	return false
}

// InList 检查是否在list中
func InList(key string, list []string) bool {
	for _, value := range list {
		if key == value {
			return true
		}
	}
	return false
}
