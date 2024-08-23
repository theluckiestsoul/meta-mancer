package file

import (
	"mime"
	"path/filepath"
)

func GetFileType(filePath string) string {
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)
	return mimeType
}
