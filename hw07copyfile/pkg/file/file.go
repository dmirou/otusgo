package file

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
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
func CopyWithProgress(from string, to string, limit int64, offset int64) error {
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
	if offset > info.Size() {
		return fmt.Errorf("offset %d is bigger than file size %d, nothing to copy",
			offset, info.Size())
	}
	var actualLimit int64
	if limit == 0 {
		actualLimit = info.Size()
	} else {
		actualLimit = limit
	}
	name := path.Base(from)
	container := mpb.New(mpb.WithWidth(64))
	bar := container.AddBar(actualLimit,
		mpb.BarStyle("[=>-|"),
		mpb.PrependDecorators(
			decor.CountersKibiByte("% .2f / % .2f "),
			decor.OnComplete(decor.Name(name, decor.WC{W: len(name), C: decor.DextraSpace}), "done!"),
		),
		mpb.AppendDecorators(decor.Percentage()),
	)
	bufSize := actualLimit / 100
	buf := make([]byte, bufSize)
	dstFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	for i := 1; i < 100; i++ {
		if err := copyPortion(srcFile, dstFile, &buf, offset); err != nil {
			return err
		}
		bar.IncrBy(int(bufSize))
		offset += bufSize
	}
	lastBufSize := bufSize + actualLimit%100
	lastBuf := make([]byte, lastBufSize)
	if err := copyPortion(srcFile, dstFile, &lastBuf, offset); err != nil {
		return err
	}
	bar.IncrBy(int(lastBufSize))

	container.Wait()

	return nil
}

func copyPortion(srcFile io.ReaderAt, dstFile io.Writer, buf *[]byte, offset int64) error {
	_, err := srcFile.ReadAt(*buf, offset)
	if err != nil && err != io.EOF {
		return err
	}
	_, err = dstFile.Write(*buf)
	if err != nil {
		return err
	}
	return nil
}
