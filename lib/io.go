package lib

import (
	"io"
	"bufio"
	"fmt"
)

// CopyLines : copy [start, end) lines from src to dest
// @Param : index start from 0
// Warning : it does not deal well with lines longer than 65536 characters
func CopyLines(src io.Reader, dest io.Writer, start, end int) error {
	scanner := bufio.NewScanner(src)

	for i := 0; i < start; i++ {
		if !scanner.Scan() {
			return fmt.Errorf("Index Error: start_page (%v) greater than total pages (%v), no output written",
				start, i);
		}
	}

	writer := bufio.NewWriter(dest)
    for ; start < end; start++ {
		if (!scanner.Scan()) {
			if err := writer.Flush(); err != nil {
				panic(err)
			}
			return fmt.Errorf("Index Error: end_page (%v) greater than total pages (%v), less output than expected",
				end, start)
		}
		fmt.Fprintln(writer, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
	return nil
}
