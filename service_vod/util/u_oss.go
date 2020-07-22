package util

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
)



func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile string) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(uploadAddressDTO.FileName, localFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func UploadReader(client *oss.Client, uploadAddressDTO UploadAddressDTO, reader io.Reader) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传文件流。
	err = bucket.PutObject(uploadAddressDTO.FileName, reader)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}