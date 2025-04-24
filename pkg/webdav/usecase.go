package webdav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lamaleka/boilerplate-golang/internal/entity"
)

type WebdavUsecase struct {
	Config *entity.ConfApiWebdav
}

func NewWebdavUseCase(config *entity.ConfApiWebdav) *WebdavUsecase {
	return &WebdavUsecase{
		Config: config,
	}
}

func (u *WebdavUsecase) View(fileName string) (*MediaViewResponse, error) {
	fmt.Println("Downloading : ", fileName)
	url := fmt.Sprintf("%s/%s/%s", u.Config.Url, u.Config.Path, fileName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(u.Config.User, u.Config.Secret)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	contentType := res.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/xml") {
		var soapErr SoapError
		body, _ := io.ReadAll(res.Body)
		err := xml.Unmarshal(body, &soapErr)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(soapErr.Message)
	}
	fileBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	webdavViewResponse := &MediaViewResponse{
		ContentType:        contentType,
		ContentDisposition: res.Header.Get("Content-Disposition"),
		FileBytes:          fileBytes,
	}
	fmt.Println("Downloaded : ", fileName)
	return webdavViewResponse, nil
}

func (u *WebdavUsecase) Upload(payload *MediaUploadRequest) (*MediaUploadResponse, error) {
	fmt.Println("Uploading : ", payload.FileName)
	url := fmt.Sprintf("%s/%s/%s", u.Config.Url, u.Config.Path, payload.FileName)
	req, err := http.NewRequest("PUT", url, bytes.NewReader(payload.FileBytes))
	if err != nil {
		return nil, errors.New("Error creating request: " + err.Error())
	}
	req.SetBasicAuth(u.Config.User, u.Config.Secret)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error uploading file: " + err.Error())
	}
	defer res.Body.Close()
	resContentType := res.Header.Get("Content-Type")
	if strings.Contains(resContentType, "xml") {
		var soapErr SoapError
		body, _ := io.ReadAll(res.Body)
		err := xml.Unmarshal(body, &soapErr)
		if err != nil {
			return nil, errors.New("Error parsing SOAP error: " + err.Error())
		}
		return nil, errors.New("Error uploading file: " + soapErr.Message)
	}
	fmt.Println("Uploaded : ", payload.FileName)

	uploadResponse := &MediaUploadResponse{
		FileName: payload.FileName,
	}

	return uploadResponse, nil
}
