package cloudinary

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"path/filepath"
	"sushi-backend/config"
	"sushi-backend/internal/tmp_file_storage"
	"sushi-backend/pkg/logger"
	"sushi-backend/utils"
)

type Cloudinary struct {
	fileStorage tmp_file_storage.ITmpFileStorage
	logger      logger.ILogger
	config      config.IConfig
	cld         *cloudinary.Cloudinary
}

func NewCloudinary(deps CloudinaryDependencies) *Cloudinary {
	cld, err := cloudinary.NewFromURL(deps.Config.CloudinaryUrl())

	if err != nil {
		deps.Logger.Error(fmt.Sprintf("Failed to create cloudinary client: %s", err))
		panic(err)
	}

	cld.Config.URL.Secure = true

	return &Cloudinary{
		fileStorage: deps.FileStorage,
		logger:      deps.Logger,
		config:      deps.Config,
		cld:         cld,
	}
}

func (c *Cloudinary) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (publicId, secureURL string) {
	tmdDir := c.fileStorage.CreateTmpDir()

	defer c.fileStorage.RemoveTmpDir(tmdDir)

	fileName := fileNameWithoutExtension(header.Filename) + "-*" + filepath.Ext(header.Filename)

	tmpFileName := c.fileStorage.WriteFile(tmdDir, fileName, file)

	resp, err := c.cld.Upload.Upload(ctx, tmpFileName, uploader.UploadParams{
		PublicID: c.generateKey(fileName),
	})

	utils.PanicIfError(err)

	c.logger.Debug(fmt.Sprintf("Uploaded file to cloudinary: public id %s %s", resp.PublicID, resp.SecureURL))

	return resp.PublicID, resp.SecureURL
}

func (c *Cloudinary) Delete(ctx context.Context, publicId string) {
	utils.PanicIfErrorWithResult(c.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicId,
	}))

	c.logger.Debug(fmt.Sprintf("Deleted file from cloudinary: public id %s", publicId))
}

func (c *Cloudinary) generateKey(filename string) string {
	return c.config.CloudinaryFolder() + "/" + filename + "-" + utils.NewUUID()
}

func fileNameWithoutExtension(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
