package chunker

import (
	"errors"
	"fmt"
	"strings"
)

func validateChunkSize(str string, csize int) bool {
	for _, s := range strings.Split(str, " ") {
		if len(s) > csize {
			fmt.Println(s)
			return false
		}
	}

	return true
}

func totalChunks(slen int, csize int) int {
	if slen%csize > 0 {
		return slen/csize + 1
	}

	return slen / csize
}

func Chunker(str string, csize int) ([]string, error) {
	if !validateChunkSize(str, csize) {
		return []string{}, errors.New("input string has a word bigger than chunk size, unexpected behaviour")
	}

	slen := len(str)
	i := 0
	j := csize - 1

	chunks := make([]string, totalChunks(slen, csize))

	for j < slen {
		savej := j

		// look backwards
		for str[j] != ' ' && i < j {
			j--
		}

		if i == j {
			// look forwards
			j = savej
			for str[j] != ' ' && j < slen {
				j++
			}
		}

		chunks = append(chunks, strings.TrimSpace(str[i:j]))
		i = j + 1
		if j+csize >= slen {
			chunks = append(chunks, strings.TrimSpace(str[i:slen]))
		}

		j = j + csize
	}

	return chunks, nil
}
