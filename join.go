package reloaded

func Join(elems []string, sep string) string {
	var output = ""
	if len(elems) == 0 {
		return ""
	}
	for i, s := range elems {
		output += s
		if i != len(elems)-1 {
			output += sep
		}
	}
	return output
}
