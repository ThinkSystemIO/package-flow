package node

func DeepCopyMap(m map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	for k, v := range m {
		// Handle maps
		mapValue, isMap := v.(map[string]interface{})
		if isMap {
			result[k] = DeepCopyMap(mapValue)
			continue
		}

		// Handle slices
		sliceValue, isSlice := v.([]interface{})
		if isSlice {
			result[k] = DeepCopySlice(sliceValue)
			continue
		}

		result[k] = v
	}

	return result
}

func DeepCopySlice(s []interface{}) []interface{} {
	result := []interface{}{}

	for _, v := range s {
		// Handle maps
		mapValue, isMap := v.(map[string]interface{})
		if isMap {
			result = append(result, DeepCopyMap(mapValue))
			continue
		}

		// Handle slices
		sliceValue, isSlice := v.([]interface{})
		if isSlice {
			result = append(result, DeepCopySlice(sliceValue))
			continue
		}

		result = append(result, v)
	}

	return result
}
