package docs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// // SaveDocumentation saves detected comments in a structured format
// func SaveDocumentation(filePath string, comments []string) {
//     file, err := os.Create(filePath)
//     if err != nil {
//         fmt.Println("Error creating documentation file:", err)
//         return
//     }
//     defer file.Close()

//     for _, comment := range comments {
//         _, err := file.WriteString(comment + "\n")
//         if err != nil {
//             fmt.Println("Error writing to documentation file:", err)
//         }
//     }
//     fmt.Println("Documentation saved at", filePath)
// }

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