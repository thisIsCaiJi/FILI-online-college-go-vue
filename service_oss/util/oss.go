package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
	"github.com/thisIsCaiJi/online_college/service_oss/config"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func UploadFile(file multipart.File,header *multipart.FileHeader) (string,error) {
	ossConf := config.GetOssConfig()
	fmt.Printf("ossConf:%v\n",ossConf)
	client, err := oss.New(ossConf.EndPoint,ossConf.KeyId,ossConf.KeySecret)
	if err != nil {
		handleError(err)
	}
	bucket, err := client.Bucket(ossConf.BucketName)
	if err != nil {
		handleError(err)
	}

	fileName := strings.Replace(uuid.NewV4().String(),"-","",-1)  + "file" + path.Ext(header.Filename)
	dateStr := time.Now().Format("2006/01/02")
	objectName := dateStr + "/" + fileName
	// 上传文件流。
	err = bucket.PutObject(objectName, file)
	if err != nil {
		handleError(err)
	}
	url := fmt.Sprintf("https://%s.%s/%s",ossConf.BucketName,ossConf.EndPoint,objectName);
	return url,nil
}