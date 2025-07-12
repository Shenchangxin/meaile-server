package utils

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
	"mime/multipart"
)

type HuaWeiOBSClient struct {
	OBSClient *obs.ObsClient
}

func (h *HuaWeiOBSClient) UploadFile(file *multipart.FileHeader, ossId, bucketName, path string) (string, string, error) {
	// 打开文件
	open, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer open.Close()

	// 构建上传参数
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: bucketName,
				Key:    ossId,
			},
			HttpHeader: obs.HttpHeader{
				ContentType: file.Header.Get("content-type"),
			},
		},
		Body: open,
	}

	_, err = h.OBSClient.PutObject(input)
	if err != nil {
		return "", "", errors.Wrap(err, "文件上传失败!")
	}

	filepath := path + "/" + file.Filename
	return filepath, file.Filename, nil
}
func (h *HuaWeiOBSClient) DeleteFile(key, bucketName, path string) error {

	input := &obs.DeleteObjectInput{
		Bucket: bucketName,
		Key:    key,
	}

	_, err := h.OBSClient.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!", key)
	}
	return nil
}
func (h *HuaWeiOBSClient) GetUrl(key, bucketName string, expires int) (string, error) {

	putObjectInput := &obs.CreateSignedUrlInput{}
	putObjectInput.Method = obs.HttpMethodGet
	putObjectInput.Bucket = bucketName
	putObjectInput.Key = key
	putObjectInput.Expires = expires
	// 生成上传对象的带授权信息的URL
	url, err := h.OBSClient.CreateSignedUrl(putObjectInput)

	if err != nil {
		return "", errors.Wrapf(err, "获取对象(%s)URL失败!", key)
	}
	return url.SignedUrl, nil
}
