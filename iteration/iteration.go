package iterate

func Repeat(rand string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += rand
	}
	return repeated
}
