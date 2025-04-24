package webdav

import (
	"encoding/xml"

	"github.com/lamaleka/boilerplate-golang/internal/model"
)

type SoapError struct {
	XMLName   xml.Name `xml:"error"`
	Exception string   `xml:"http://sabredav.org/ns exception"`
	Message   string   `xml:"http://sabredav.org/ns message"`
}

type MediaUploadRequest = model.MediaUploadRequest
type MediaUploadResponse = model.MediaUploadResponse
type MediaViewResponse = model.MediaViewResponse
