package models

// UploaderS3Conf holds values needed for uploading a new
// object to S3 Buckets
// Its values are used by the Upload method for configuring
// the s3.PutObjectInput
type UploaderS3Conf struct {
	UploaderCommons      `mapstructure:",squash"`
	Bucket               string `validate:"required" env:"S3_BUCKET"`                                      // ex : "mybucket"
	ACL                  string `validate:"required" env:"S3_ACL" envDefault:"private"`                    // ex : "public-read" -> https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
	ContentDisposition   string `validate:"required" env:"S3_CONTENT_DISPOSITION" envDefault:"attachment"` // ex : "attachment"
	ServerSideEncryption string `validate:"required" env:"S3_ENCRYPTION" envDefault:"AES256"`              // ex : "AES256"
	StorageClass         string `validate:"required" env:"S3_STORAGE_CLASS" envDefault:"STANDARD"`         // ex : "INTELLIGENT_TIERING"
	Region               string `validate:"required" env:"S3_REGION"`                                      // ex : "us-east-1"
	SecretID             string `validate:"required" env:"S3_ACCESS_KEY_ID"`
	SecretKey            string `validate:"required" env:"S3_SECRET_ACCESS_KEY"`
	Token                string `env:"S3_TOKEN"`
}
