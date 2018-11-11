# App Uploader

### Brief Overview

Have a basic framework that allows you to have a user upload some files and have them processed on the server-side.

The example within takes a list of files from the application and returns them all in a zipfile.

There is 2 endpoints required:

- IndexHandler - for serving the html content
- UploadHandler - for working with the static files

### Building image

packr command is required to package the text files into a binary format

```bash
CGO_ENABLED=0 GOOS=linux packr build -a -installsuffix cgo -o app .
docker build .
```
