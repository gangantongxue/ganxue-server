package bf

import "github.com/russross/blackfriday/v2"

func ToHTML(input []byte) []byte {
	return blackfriday.Run(input)
}
