package form

import (
	"io/ioutil"
	"mime/multipart"
	"os"
)

// PartFile represents form-data file part model.
type PartFile struct {
	filePath *string
}

// Marshal encodes part file to writer.
func (p PartFile) Marshal(w *multipart.Writer, partName string, omitempty bool) error {
	var fPath string
	switch {
	case p.filePath == nil:
		if omitempty {
			return nil
		}
	default:
		fPath = *p.filePath
	}

	file, err := os.Open(fPath)
	if err != nil {
		return err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}

	part, err := w.CreateFormFile(partName, fi.Name())
	if err != nil {
		return err
	}
	if _, err := part.Write(fileContents); err != nil {
		return err
	}

	return nil
}

// NewPartFile returns new PartFile instance.
func NewPartFile(filePath string) PartFile {
	return PartFile{
		filePath: &filePath,
	}
}
