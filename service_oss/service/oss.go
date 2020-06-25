package service

import (
	"errors"
	"github.com/thisIsCaiJi/online_college/service_oss/util"
	"mime/multipart"
)

type OssService struct {

}


func UploadFileAvtar(file multipart.File,header *multipart.FileHeader) (string, error) {
	dataLen := header.Size
	if dataLen > int64(2)*1024*1024 {
		err := errors.New("File size cannot be greater than 5MB!")
		return "",err
	}


	// 上传文件流。
	url,err := util.UploadFile(file,header)
	return url,err
}
