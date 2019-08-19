package models

// UploaderFSConf holds values needed for uploading objects using
// the underlying os filesystem
type UploaderFSConf struct {
	UploaderCommons `mapstructure:",squash"`
	Folder          string `validate:"required"`
}
