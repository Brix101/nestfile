package files

import "os"

type FileReader struct {
	basePath string
}

func NewFileReader(basePath string) *FileReader {
	return &FileReader{
		basePath: basePath,
	}
}

func (fSer *FileReader) ListFiles() ([]string, error) {
	files, err := os.ReadDir(fSer.basePath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}
