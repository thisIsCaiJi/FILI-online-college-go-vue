package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/thisIsCaiJi/online_college/service_vod/config"
)


func main(){
	client, err := InitVodClient(config.GetVodConfig().KeyId,config.GetVodConfig().KeySecret)
	if err != nil {
		panic(err)
	}
	response, err := MyCreateUploadVideo(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetHttpContentString())
	fmt.Printf("VideoId: %s\n UploadAddress: %s\n UploadAuth: %s",
		response.VideoId, response.UploadAddress, response.UploadAuth)
}

func MyCreateUploadVideo(client *vod.Client) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = "Sample Video Title"
	request.Description = "Sample Description"
	request.FileName = "/opt/video/sample/video_file.mp4"
	//request.CateId = "-1"
	request.CoverURL = "http://img.alicdn.com/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-700.png"
	request.Tags = "tag1,tag2"
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

func MyGetPlayInfo(client *vod.Client, videoId string) (response *vod.GetPlayInfoResponse, err error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetPlayInfo(request)
}

func MyGetPlayAuth(client *vod.Client, videoId string) (response *vod.GetVideoPlayAuthResponse, err error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.GetVideoPlayAuth(request)
}

func InitVodClient(accessKeyId  string, accessKeySecret string) (client *vod.Client, err error) {

	// 点播服务接入区域
	regionId := "cn-shanghai"

	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}

	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true      // 失败是否自动重试
	config.MaxRetryTime = 3      // 最大重试次数
	config.Timeout = 3000000000  // 连接超时，单位：纳秒；默认为3秒

	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}
