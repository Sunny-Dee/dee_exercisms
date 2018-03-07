package flatten

// Flatten makes a nested list into a nice array of ints
func Flatten(input interface{}) []interface{} {
	result := []interface{}{}
	flatten(input, &result)
	return result
}

func flatten(input interface{}, result *[]interface{}) {
	if items, ok := input.([]interface{}); ok {
		for i := 0; i < len(items); i++ {
			flatten(items[i], result)
		}
	}

	item, _ := input.(interface{})
	if num, ok := item.(int); ok {
		*result = append(*result, num)
	}
}
