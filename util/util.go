package util

func UnwrapMap(inputMap map[string]string) []string {
	res := make([]string, 0, len(inputMap)*2)
	for key, val := range inputMap {
		res = append(res, key, val)
	}
	return res
}
