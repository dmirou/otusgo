package file

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path"
	"testing"
)

const (
	srcDir   = "./src"
	destDir  = "./dest"
	Byte     = 1
	KiloByte = 1024 * Byte
	MegaByte = 1024 * Byte
)

type srcFiles int

const (
	oneByte srcFiles = iota
	fiveBytes
	oneKb
	oneMb
	tenMb
)

func (d srcFiles) String() string {
	return [...]string{"oneByte", "fiveBytes", "oneKb", "oneMb", "tenMb"}[d]
}

var sourceFiles = map[string]int{
	oneByte.String():   1 * Byte,
	fiveBytes.String(): 5 * Byte,
	oneKb.String():     1 * KiloByte,
	oneMb.String():     1 * MegaByte,
	tenMb.String():     10 * MegaByte,
}

func TestMain(m *testing.M) {
	if err := os.RemoveAll(srcDir); err != nil {
		exitWithError(err)
	}
	if err := os.RemoveAll(destDir); err != nil {
		exitWithError(err)
	}
	if err := os.Mkdir(srcDir, os.ModeDir|os.ModePerm); err != nil {
		exitWithError(err)
	}
	if err := os.Mkdir(destDir, os.ModeDir|os.ModePerm); err != nil {
		exitWithError(err)
	}

	for name, size := range sourceFiles {
		path := buildFilePath(srcDir, name)
		if err := generateRandomFile(path, size); err != nil {
			exitWithError(err)
		}
	}
	code := m.Run()
	os.Exit(code)
}

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func buildFilePath(baseDir string, filename string) string {
	return path.Join(baseDir, filename)
}

func generateRandomFile(path string, size int) error {
	buf := make([]byte, size)
	reader := io.LimitReader(rand.Reader, int64(size))
	if _, err := reader.Read(buf); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	n, err := file.Write(buf)
	if err != nil {
		return err
	}
	if n != size {
		return fmt.Errorf("can't write %d bytes, written only %d", size, n)
	}
	return nil
}

func TestCopy(t *testing.T) {
	testCases := []struct {
		srcName string
		limit   int
		offset  int
		err     error
	}{
		{
			srcName: oneByte.String(),
			limit:   0,
			offset:  0,
			err:     nil,
		},
	}
	for _, testCase := range testCases {
		srcPath := buildFilePath(srcDir, testCase.srcName)
		destPath := buildFilePath(destDir, testCase.srcName)
		err := CopyWithProgress(srcPath, destPath, testCase.limit, testCase.offset)
		if err != testCase.err {
			t.Errorf("result not matches for testCase %v, expected: %s, actual %s",
				testCase, testCase.err, err)
		}
	}
}
