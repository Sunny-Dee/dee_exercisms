package sublist

// Relation is just a string representing the sublist response
type Relation string

// Sublist checks for a sublist from two input slices
func Sublist(listOne, listTwo []int) Relation {
	switch {
	case len(listOne) == 0 && len(listTwo) != 0:
		return "sublist"
	case len(listOne) != 0 && len(listTwo) == 0:
		return "superlist"
	case len(listOne) == len(listTwo):
		return equal(listOne, listTwo)
	case len(listOne) > len(listTwo):
		return eval(listOne, listTwo, "superlist")
	case len(listOne) < len(listTwo):
		return eval(listTwo, listOne, "sublist")
	default:
		return "no relation"
	}
}

func eval(long, short []int, response Relation) Relation {
	if sub(long, short) {
		return response
	}
	return "unequal"
}

func sub(long, short []int) bool {
	for i, val := range long {
		if val == short[0] && checkSub(long[i:], short) {
			return true
		}
	}
	return false
}

func equal(listA, listB []int) Relation {
	for i := range listA {
		if listA[i] != listB[i] {
			return "unequal"
		}
	}
	return "equal"
}

func checkSub(listA, listB []int) bool {
	// if len(listA) < len(listB) {
	// 	return false
	// }

	for i := range listB {
		if listA[i] != listB[i] {
			return false
		}
	}

	return true
}
