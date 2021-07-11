package reloaded

func Join(elems []string, sep string) string {
	var output = ""
	if len(elems) == 0 {
		return ""
	}
	for i, v := range elems {
		output += v
		if i != len(elems)-1 {
			output += sep
		}
	}
	return output
}
