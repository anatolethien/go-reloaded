package reloaded

func Join(elems []string, sep string) string {
	var output = ""
	if len(elems) == 0 {
		return ""
	}
	if len(elems) == 1 {
		return elems[0]
	}

	for i, s := range elems {
		output += s
		if i != len(elems)-1 {
			output += sep
		}
	}
	return output
}
