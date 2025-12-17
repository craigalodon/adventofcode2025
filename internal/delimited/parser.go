package delimited

func ParseSpaceDelimited(s string) []string {
	vals := make([]string, 0)

	h := 0
	t := 0

	for _, r := range s {
		if r == ' ' {
			if h != t && s[h] != ' ' {
				vals = append(vals, s[h:t])
			}
			h = t + 1
		}
		t++
	}
	if h != t && s[h] != ' ' {
		vals = append(vals, s[h:t])
	}

	return vals
}
