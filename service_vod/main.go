package main

import (
	"github.com/thisIsCaiJi/online_college/service_vod/config"
	"github.com/thisIsCaiJi/online_college/service_vod/service"
)

func main() {
	localFile := "C:/Users/xp/Pictures/视频项目/6 - What If I Want to Move Faster.mp4";
	client, err := service.InitVodClient(config.GetVodConfig().KeyId,config.GetVodConfig().KeySecret)
	if err != nil {
		panic(err)
	}
	uploadVideo := service.UploadVideo{
		Title: "这是标题" ,
		Description: "这是描述" ,
		FileName: "this is file's name.mp4" ,
		CoverURL: "这是图片地址" ,
		Tags: "这是tags" ,
	}
	service.UploadVideoLocalFile(client,uploadVideo,localFile)
}


