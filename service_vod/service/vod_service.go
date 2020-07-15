package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"time"
)

type UploadAddressDTO struct {
	Endpoint string `json:"Endpoint"`
	Bucket   string `json:"Bucket"`
	FileName string `json:"FileName"`
}

type UploadAuthDTO struct {
	SecurityToken   string    `json:"SecurityToken"`
	AccessKeyId     string    `json:"AccessKeyId"`
	AccessKeySecret string    `json:"AccessKeySecret"`
	Region          string    `json:"Region"`
	ExpireUTCTime   time.Time `json:"ExpireUTCTime"`
	Expiration      int       `json:"Expiration"`
}

type UploadVideo struct {
	Title string
	Description string
	FileName string
	CoverURL string
	Tags string
}

// 创建视频上传请求并获取videoId、上传地址和凭证
func CreateUploadVideo(client *vod.Client, video UploadVideo) (videoId, UploadAddress, uploadAuth string, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = video.Title
	request.Description = video.Description
	request.FileName = "/opt/video/sample/" + video.FileName
	//request.CateId = "-1"
	request.CoverURL = video.CoverURL
	request.Tags = video.Tags
	request.AcceptFormat = "JSON"
	response, err := client.CreateUploadVideo(request)
	return response.VideoId, response.UploadAddress, response.UploadAuth, err
}

// 根据videoId获取视频信息
func GetPlayInfo(client *vod.Client, videoId string) (playInfos []vod.PlayInfo, err error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	response, err := client.GetPlayInfo(request)
	return response.PlayInfoList.PlayInfo, err
}

//根据videoId获取播放凭证
func GetPlayAuth(client *vod.Client, videoId string) (playAuth string, err error) {
	request := vod.CreateGetVideoPlayAuthRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	response, err := client.GetVideoPlayAuth(request)
	return response.PlayAuth, err
}

// 刷新视频上传凭证
func RefreshUploadVideo(client *vod.Client, videoId string) (response *vod.RefreshUploadVideoResponse, err error) {
	request := vod.CreateRefreshUploadVideoRequest()
	request.VideoId = videoId
	request.AcceptFormat = "JSON"
	return client.RefreshUploadVideo(request)
}

// 初始化vod客户端
func InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {

	// 点播服务接入区域
	regionId := "cn-shanghai"

	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}

	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒

	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

func UploadVideoLocalFile(client *vod.Client, uploadVideo UploadVideo, localFile string) {
	videoId, uploadAddress, uploadAuth, err := CreateUploadVideo(client, uploadVideo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("VideoId: %s\n UploadAddress: %s\n UploadAuth: %s",
		videoId, uploadAddress, uploadAuth)
	uploadAddressDecode, err := base64.StdEncoding.DecodeString(uploadAddress)
	if err != nil {
		panic(err)
	}
	var uploadAddressDTO UploadAddressDTO
	if err := json.Unmarshal(uploadAddressDecode, &uploadAddressDTO); err != nil {
		panic(err)
	}
	uploadAuthDecode, err := base64.StdEncoding.DecodeString(uploadAuth)
	if err != nil {
		panic(err)
	}
	var uploadAuthDTO UploadAuthDTO
	json.Unmarshal(uploadAuthDecode,&uploadAuthDTO)
	ossClient, err := InitOssClient(uploadAuthDTO, uploadAddressDTO)
	if err != nil {
		panic(err)
	}
	UploadLocalFile(ossClient,uploadAddressDTO,localFile)
	fmt.Println("Succeed, VideoId:", videoId)
}