package fs

import (
	"github.com/endpot/SpiderX-Backend/app/infra/config"
	"github.com/endpot/SpiderX-Backend/app/infra/oss"
	"github.com/minio/minio-go/v6"
	"os"
)

func UploadToRemote(localPath string, remotePath string) error {
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		return err
	}

	_, err := oss.Client.FPutObject(
		config.Config.GetString("OSS_BUCKET"),
		remotePath,
		localPath, minio.PutObjectOptions{},
	)

	return err
}
