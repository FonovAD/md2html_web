package processing

import (
	"errors"
	"os"
)

func ReadFile(FileName string) ([]byte, error) {
	MarkdownFile, err := os.Open(FileName)
	if err != nil {
		return []byte{0}, errors.New("error opening file")
	}
	defer MarkdownFile.Close()
	FileData := make([]byte, 8192)
	for {
		MarkdownFile.Read(FileData)
		if FileData[len(FileData)-1] == byte(0) {
			break
		}
	}
	return FileData, nil
}

func Split(file []byte) ([]string, error) {
	buf := []string{}
	start := 0
	for i := 0; i < len(file); i++ {
		if file[i] == byte(10) || i == len(file)-1 {
			buf = append(buf, string(file[start:i]))
			start = i + 1
		}
	}
	return buf, nil
}
