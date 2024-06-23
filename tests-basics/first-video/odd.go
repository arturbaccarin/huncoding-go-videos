package firstvideo

var (
	ODD_KEYWORD  = "ODD"
	EVEN_KEYWORD = "EVEN"
)

func isOdd(value int64) string {
	if value%2 == 0 {
		return EVEN_KEYWORD
	}

	return ODD_KEYWORD
}
