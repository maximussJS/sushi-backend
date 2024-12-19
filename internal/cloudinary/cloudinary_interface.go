package cloudinary

import (
	"context"
	"mime/multipart"
)

type ICloudinary interface {
	Upload(ctx context.Context, file multipart.File, handler *multipart.FileHeader) (publicId, secureURL string)
}
