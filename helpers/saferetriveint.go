package helpers

func SafeRetrieveInt(value interface{}, path string) int {
	result, _ := SafeRetrieve(value, path).(int)
	return result
}
