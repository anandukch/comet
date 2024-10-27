package scanner

import (
	"bufio"
	"comet/docs"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var commentDelimiters = map[string]string{
	".go":   "//",
	".py":   "#",
	".js":   "//",
	".java": "//",
	".c":    "//",
	".cpp":  "//",
	".rb":   "#",
	".php":  "//",
	".rs":   "//",
	".cs":   "//",
	".ts":   "//",
}

type Comment struct {
	FilePath string
	Line     int
	Text     string
}

var comments []Comment

func ScanProject(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if !info.IsDir() && isSupportedFile(ext) { // Scans only Go files for simplicity
			detectComments(path, commentDelimiters[ext])
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error scanning project:", err)
	}
}
func isSupportedFile(ext string) bool {
	_, exists := commentDelimiters[ext]
	return exists
}

func detectComments(filePath, delimiter string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	linesToRemove := []int{}

	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), delimiter) {
			comment := Comment{
				FilePath: filePath,
				Line:     i + 1,
				Text:     line,
			}
			comments = append(comments, comment)

			linesToRemove = append(linesToRemove, i)

		}
	}

	if len(linesToRemove) > 0 && promptForRemoval(filePath) {
		removeComments(filePath, lines, linesToRemove)
	}

	err = saveUpdatedFile(filePath, lines)
	if err != nil {
		fmt.Println("Error saving updated file:", err)
	}

	docsComments := make([]docs.Comment, len(comments))
	for i, c := range comments {
		docsComments[i] = docs.Comment{
			FilePath: c.FilePath,
			Line:     c.Line,
			Text:     c.Text,
		}
	}

	err = docs.SaveDocumentation("comments.txt", docsComments)
	if err != nil {
		fmt.Printf("Error saving documentation: %v\n", err)
	}
}

func promptForRemoval(filePath string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Found comment at %s. Do you want to remove it? (y/n): ", filePath)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text) == "y"
}

func removeComments(filePath string, lines []string, linesToRemove []int) {
	for _, i := range linesToRemove {
		lines[i] = "" // Remove the comment line

	}
	fmt.Printf("Comment removed from %s\n", filePath)

}

func saveUpdatedFile(filePath string, lines []string) error {
	return os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
}
