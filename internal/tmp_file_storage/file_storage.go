package tmp_file_storage

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sushi-backend/pkg/logger"
	"sushi-backend/utils"
)

type TmpFileStorage struct {
	logger logger.ILogger
}

func NewTmpFileStorage(deps FileStorageDependencies) *TmpFileStorage {
	return &TmpFileStorage{
		logger: deps.Logger,
	}
}

func (fs *TmpFileStorage) CreateTmpDir() string {
	tmpDir := os.TempDir()

	fs.logger.Debug(fmt.Sprintf("Creating temporary directory in %s", tmpDir))

	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		utils.PanicIfError(os.Mkdir(tmpDir, os.ModePerm))
	}

	return tmpDir
}

func (fs *TmpFileStorage) WriteFile(tmpDir, filename string, file multipart.File) string {
	defer file.Close()

	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		utils.PanicIfError(os.Mkdir(tmpDir, os.ModePerm))
	}

	f, err := os.CreateTemp(tmpDir, filename)
	utils.PanicIfError(err)

	bytes, err := io.ReadAll(file)
	utils.PanicIfError(err)

	utils.PanicIfErrorWithResult(f.Write(bytes))

	return f.Name()
}

func (fs *TmpFileStorage) RemoveTmpDir(tmpDir string) {
	fs.logger.Debug("Removing temporary directory")
	utils.PanicIfError(os.RemoveAll(tmpDir))
}
