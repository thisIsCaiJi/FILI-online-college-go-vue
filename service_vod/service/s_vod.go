package service

import (
	"github.com/thisIsCaiJi/online_college/service_vod/util"
	"mime/multipart"
	"strings"
)

func UploadVideo(file multipart.File, header *multipart.FileHeader) (videoId string, err error) {
	client := util.GetVodClient()
	uploadVideo := util.UploadVideo{
		Title: header.Filename[:strings.LastIndex(header.Filename,".")] ,
		FileName: header.Filename ,
	}
	videoId, err = util.UploadVideoReader(client,uploadVideo,file)
	return
}

func GetVideo(id string) {

}

func DeleteVideo(id string) error {
	client := util.GetVodClient()
	_,err :=util.DeleteVideo(client,id)
	return err
}