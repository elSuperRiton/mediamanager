package s3

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/elSuperRiton/mediamanager/pkg/models"
)

// Uploader is a struct implementing the Uploader interface
// from the uploader package
type Uploader struct {
	conf    *models.UploaderS3Conf
	session *session.Session
	s3      *s3.S3
}

// NewUploader returns an instance of an Uploader initialized with a ready to use
// AWS sessions for s3 upload
func NewUploader(conf models.UploaderS3Conf) (Uploader, error) {

	s, err := session.NewSession(&aws.Config{
		Region: aws.String(conf.Region),
		Credentials: credentials.NewStaticCredentials(
			conf.SecretID,
			conf.SecretKey,
			conf.Token,
		),
	})

	if err != nil {
		return Uploader{}, err
	}

	return Uploader{
		conf:    &conf,
		session: s,
		s3:      s3.New(s),
	}, nil
}

// Upload attempts to upload a file to s3 using the driver's underlying
// config
func (r Uploader) Upload(f multipart.File, fH *multipart.FileHeader) error {

	if r.conf == nil || r.session == nil {
		panic("s3 uploader must be initiliazed prior to being used")
	}

	// get the file size and read
	// the file content into a buffer
	size := fH.Size
	buffer := make([]byte, size)
	f.Read(buffer)

	// check if plugin exists and extact filename from it
	if r.conf.PluginName != "" {
		symbol, err := r.conf.Plugin.Lookup("GetFileName")
		if err != nil {
			panic(err)
		}

		getNameFunc, ok := symbol.(func(file *multipart.FileHeader) (newFileName string, err error))
		if !ok {
			panic("Plugin has no 'func(file *multipart.FileHeader) (newFileName string, err error)' function")
		}

		fmt.Println(getNameFunc(fH))
	}

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := r.s3.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(r.conf.Bucket),
		ContentDisposition:   aws.String(r.conf.ContentDisposition),
		ServerSideEncryption: aws.String(r.conf.ServerSideEncryption),
		StorageClass:         aws.String(r.conf.StorageClass),
		ACL:                  aws.String(r.conf.ACL),
		Key:                  aws.String(fH.Filename),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
	})

	return err
}

// GetS3PresignedPutURL
func (r Uploader) GetS3PresignedPutURL(expire time.Duration, fileName, fileType string) (string, error) {
	// Create S3 service client
	req, _ := r.s3.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(r.conf.Bucket),
		Key:         aws.String(fileName),
		ContentType: aws.String(fileType),
	})

	str, err := req.Presign(expire)
	if err != nil {
		return "", err
	}

	return str, nil
}
