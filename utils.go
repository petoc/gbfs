package gbfs

func indexOfStr(n string, h []string) int {
	for k, v := range h {
		if n == v {
			return k
		}
	}
	return -1
}
