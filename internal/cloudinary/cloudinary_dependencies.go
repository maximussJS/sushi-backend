package cloudinary

import (
	"context"
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/tmp_file_storage"
	"sushi-backend/pkg/logger"
)

type CloudinaryDependencies struct {
	dig.In

	ShutdownContext context.Context                  `name:"ShutdownContext"`
	Logger          logger.ILogger                   `name:"Logger"`
	Config          config.IConfig                   `name:"Config"`
	FileStorage     tmp_file_storage.ITmpFileStorage `name:"TmpFileStorage"`
}
