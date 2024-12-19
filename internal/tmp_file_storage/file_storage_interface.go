package tmp_file_storage

import "mime/multipart"

type ITmpFileStorage interface {
	CreateTmpDir() string
	RemoveTmpDir(tmpDir string)
	WriteFile(tmpDir, filename string, file multipart.File) string
}
