package upload

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/rainss/ossupload/pkg/LogGet"
	"os"
	"path"
)

func OssUpload(ossFilePath string, localFilePath string) {
	v1 := ossConfigValue{}
	v1.ossConfigSet()
	//fmt.Println(v1)
	// 创建OSSClient实例。
	client, err := oss.New(v1.endpoint, v1.accessKeyID, v1.accessKeySecret)
	if err != nil {
		//fmt.Println("Error:", err)
		LogGet.Error(fmt.Sprintf("%v", err))
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(v1.bucketName)
	if err != nil {
		//fmt.Println("Error:", err)
		LogGet.Error(fmt.Sprintf("%v", err))
		os.Exit(-1)
	}

	// 设置分片大小为100 KB，指定分片上传并发数为3，并开启断点续传上传。
	// 其中<yourObjectName>与objectKey是同一概念，表示断点续传上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// "LocalFile"为filePath，100*1024为partSize。isRecursive判断是否为目录

	//上传目录至oss
	filePathMap := GetFiles(localFilePath, localFilePath)
	for k, v := range *filePathMap {
		fmt.Printf("upload%v to %v\r\n", k, path.Join(ossFilePath, v))
		//fmt.Println(bucket)
		LogGet.Info(fmt.Sprintf("upload%v to %v\r\n", k, path.Join(ossFilePath, v)))
		err = bucket.UploadFile(path.Join(ossFilePath, v), k, 1000*1024,
			oss.Routines(3), oss.Checkpoint(true, ""), oss.Progress(&OssProgressListener{}))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
	}
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}
