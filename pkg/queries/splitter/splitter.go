package splitter

import "strings"

const Separator = "#"

type splitter struct {
	chunks []string
}

type Splitter interface {
	Chunks() []string
}

func (s *splitter) Chunks() []string {
	return s.chunks
}

func NewSplitter(sql string) Splitter {
	sql = removeWhitespace(sql)
	return &splitter{chunks: strings.Split(sql, Separator)}
}

func removeWhitespace(s string) string {
	var result strings.Builder
	inQuote := false

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '\'' {
			inQuote = !inQuote
			result.WriteByte(char)
			continue
		}

		if char == ' ' {
			if !inQuote && (i == 0 || s[i-1] != ' ') {
				result.WriteByte(Separator[0])
			}
		} else {
			result.WriteByte(char)
		}
	}

	return result.String()
}
