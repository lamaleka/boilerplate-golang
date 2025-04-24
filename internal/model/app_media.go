package model

type MediaUploadRequest struct {
	FileName  string
	FileBytes []byte
}
type MediaUploadResponse struct {
	FileName string `json:"file_name"`
}
type MediaViewResponse struct {
	ContentType        string
	ContentDisposition string
	FileBytes          []byte
}
