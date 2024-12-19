package tmp_file_storage

import (
	"go.uber.org/dig"
	"sushi-backend/pkg/logger"
)

type FileStorageDependencies struct {
	dig.In

	Logger logger.ILogger `name:"Logger"`
}
