package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSPictureBed struct {
	EndPoint   string
	BucketName string
	AccessID   string
	AccessKey  string

	Client *oss.Client
	Bucket *oss.Bucket
}

func (o *OSSPictureBed) Connect() error {
	client, err := oss.New(o.EndPoint, o.AccessID, o.AccessKey)
	if err != nil {
		return err
	}
	o.Client = client

	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return err
	}
	o.Bucket = bucket

	return nil
}

func (o *OSSPictureBed) UploadPic(obj, file string) error {
	err := o.Bucket.PutObjectFromFile(obj, file)
	if err != nil {
		return err
	}

	return nil
}
