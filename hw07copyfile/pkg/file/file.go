package file

import (
	"fmt"
	"io"
	"os"
)

// CopyWithProgress copies bytes from a specified path to a specified destination with a specified
// offset and shows a progress bar.
// You can limit count of bytes to be copied with the limit param.
// Limit equals zero means that we will skip it.
// Also you can specify offset in bytes started from the beginning of a file.
// Offset equals zero means that we start copying from the beginning of a file.
// It returns nil if bytes were copied successfully.
// It returns error if It can't define a file size to copy from.
// It returns error if offset greater than the source file size and there is nothing to copy.
func CopyWithProgress(from string, to string, limit int, offset int) error {
	srcFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	info, err := srcFile.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("can't copy directory, only simple file")
	}
	if int64(offset) > info.Size() {
		return fmt.Errorf("offset %d is bigger than file size %d, nothing to copy",
			offset, info.Size())
	}
	var bufSize int64
	if limit == 0 {
		bufSize = info.Size()
	} else {
		bufSize = int64(limit)
	}
	buf := make([]byte, bufSize)
	_, err = srcFile.ReadAt(buf, int64(offset))
	if err != nil && err != io.EOF {
		return err
	}
	dstFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
