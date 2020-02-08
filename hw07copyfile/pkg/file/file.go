package file

// Copy copies bytes from a specified path to a specified destination with a specified offset.
// You can limit count of bytes to be copied with the limit param.
// Limit equals zero means that we will skip it.
// Also you can specify offset in bytes started from the beginning of a file.
// Offset equals zero means that we start copying from the beginning of a file.
// Copy returns nil if bytes were copied successfully.
// It returns error if It can't define a file size to copy from.
// It returns error if offset greater than the source file size and there is nothing to copy.
func Copy(from string, to string, limit int, offset int) error {
	return nil
}
