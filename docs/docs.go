package docs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Comment struct {
	FilePath string
	Line     int
	Text     string
}

func SaveDocumentation(filename string, comments []Comment) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	currentFile := ""

	for _, comment := range comments {
		if comment.FilePath != currentFile {
			if currentFile != "" {
				writer.WriteString("\n")
			}
			writer.WriteString(comment.FilePath + "\n")
			currentFile = comment.FilePath
		}
		writer.WriteString(fmt.Sprintf("Line %d: %s\n", comment.Line, strings.TrimSpace(comment.Text)))
	}

	return writer.Flush()
    
}
