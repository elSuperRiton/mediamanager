package handlers

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/elSuperRiton/mediamanager/pkg/uploader"
	"github.com/elSuperRiton/mediamanager/pkg/utils"
)

// func init() {
// 	u, err := s3.NewUploader(*config.Conf.UploaderS3Conf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	uploader.InitUploader(u)
// }

// MediasFindAll returns the list of all medias uploaded through the
// media manager
func MediasFindAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented yet"))
}

// MediasUpload handler uploading single/multiple medias
func MediasUpload(w http.ResponseWriter, r *http.Request) {
	uploaderType := r.Context().Value("uploader").(string)

	// Get file from form
	if err := r.ParseMultipartForm(5 * 1024 * 1024); err != nil {
		utils.RenderErr(w, r, err.Error(), http.StatusInternalServerError)
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		utils.RenderErr(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Upload file using uploader package
	// Please note that the package initialization happends in the above
	// init function
	if err := repository.conf.InitializedUploaders[uploaderType].Upload(file, fileHeader); err != nil {

		// Cast error with awserr package
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "AccessDenied":
				utils.RenderErr(
					w,
					r,
					aerr.Message(),
					http.StatusForbidden,
				)
			default:
				utils.RenderErr(
					w,
					r,
					err.Error(),
					http.StatusInternalServerError,
				)
			}
			return
		}

		utils.RenderErr(
			w,
			r,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	// Upload is ok
	w.WriteHeader(http.StatusOK)
	utils.RenderData(
		w,
		r,
		"ok",
		http.StatusOK,
	)
	return
}

// MediasUploadURL handles the creation of a signer url for processing
// uplooad on the Front-end side
// -
// params:
//   fileName - required - string
//   fileType - required - string
// -
// Please note that it uses the GetS3PresignedPutURL which only works
// if using the s3 driver for the uploader package
func MediasUploadURL(w http.ResponseWriter, r *http.Request) {

	// TODO: Validate query
	fileName := r.URL.Query().Get("fileName")
	fileType := r.URL.Query().Get("fileType")

	uri, err := uploader.GetS3PresignedPutURL(15*time.Minute, fileName, fileType)
	if err != nil {
		utils.RenderErr(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RenderData(w, r, map[string]interface{}{
		"signedUrl": uri,
	}, http.StatusOK)
	return
}
