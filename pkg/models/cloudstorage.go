package models

// UploaderCloudStorageConf holds values needed for setting up
// media manager with cloud storage upload capacity
type UploaderCloudStorageConf struct {
	UploaderCommons `mapstructure:",squash"`
}
