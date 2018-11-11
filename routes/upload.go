package routes

import (
	"archive/zip"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

// Handles receiving the files and handing them out to get processed and merged
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	// create a buffer to write the zipfile to.
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)

	// Sort through the form-data receiving each of the files
	for _, fileheaders := range r.MultipartForm.File {
		for _, header := range fileheaders {
			// write each of the files into the zipfile
			packageFile(zw, header)
		}
	}

	// Close the zipfile to complete the buffer
	zw.Close()

	// Copy the zip buffer to the response stream
	io.Copy(w, buf)
}

func packageFile(zw *zip.Writer, header *multipart.FileHeader) {
	file, err := header.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w, err := zw.Create(header.Filename)
	if err != nil {
		panic(err)
	}
	io.Copy(w, file)
}
