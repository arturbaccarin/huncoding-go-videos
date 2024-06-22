package singleresponsibilityprinciple

import (
	"os"
)

func PrintNews(filename string, journal JournalClassic) {
	_ = os.WriteFile(filename, []byte(journal.String()), 0644)
}
