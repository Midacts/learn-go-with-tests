package iteration

// Repeat outputs the character repeats x amount of times
func Repeat(str string, times int) string {
	var list string
	for i := 0; i < times; i++ {
		list += str
	}
	return list
}
