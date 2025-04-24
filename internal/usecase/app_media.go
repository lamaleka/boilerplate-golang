package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lamaleka/boilerplate-golang/common/utils"
)

type AppMediaUseCase struct {
	Webdav ApiWebdavUsecase
}

func NewMediaUseCase(webdav ApiWebdavUsecase) *AppMediaUseCase {
	return &AppMediaUseCase{
		Webdav: webdav,
	}
}

func (u *AppMediaUseCase) View(fileName string) (*MediaViewResponse, error) {
	if strings.HasPrefix(fileName, ".generated/") {
		fileBytes, err := os.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
		newFileName := filepath.Base(fileName)
		res := &MediaViewResponse{
			ContentType:        utils.GetMimeType(filepath.Ext(newFileName)),
			ContentDisposition: fmt.Sprintf(`attachment; filename*=UTF-8''%s; filename="%s"`, newFileName, newFileName),
			FileBytes:          fileBytes,
		}
		return res, nil
	}
	return u.Webdav.View(fileName)

}
