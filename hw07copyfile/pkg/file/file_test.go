package file

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const (
	srcDir   = "./src"
	destDir  = "./dest"
	Byte     = 1
	KiloByte = 1024 * Byte
	MegaByte = 1024 * KiloByte
)

type srcFiles int

const (
	oneByte srcFiles = iota
	fiveBytes
	oneKb
	oneMb
	tenMb
	hundredMb
	oneGb
)

func (d srcFiles) String() string {
	return [...]string{"oneByte", "fiveBytes", "oneKb", "oneMb", "tenMb", "hundredMb", "oneGb"}[d]
}

var sourceFiles = map[string]int{
	oneByte.String():   1 * Byte,
	fiveBytes.String(): 5 * Byte,
	oneKb.String():     1 * KiloByte,
	oneMb.String():     1 * MegaByte,
	tenMb.String():     10 * MegaByte,
	tenMb.String():     10 * MegaByte,
	hundredMb.String(): 100 * MegaByte,
	oneGb.String():     1000 * MegaByte,
}

func TestMain(m *testing.M) {
	rmDirs()
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
	rmDirs()
	os.Exit(code)
}

func rmDirs() {
	if err := os.RemoveAll(srcDir); err != nil {
		exitWithError(err)
	}
	if err := os.RemoveAll(destDir); err != nil {
		exitWithError(err)
	}
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
	testData := []struct {
		srcName  string
		limit    int64
		offset   int64
		hasError bool
	}{
		{
			srcName:  oneByte.String(),
			limit:    0,
			offset:   0,
			hasError: false,
		},
		{
			srcName:  fiveBytes.String(),
			limit:    0,
			offset:   0,
			hasError: false,
		},
		{
			srcName:  tenMb.String(),
			limit:    0,
			offset:   0,
			hasError: false,
		},
	}
	for _, td := range testData {
		name := fmt.Sprintf("file %q, limit %d, offset %d", td.srcName, td.limit, td.offset)
		srcPath := buildFilePath(srcDir, td.srcName)
		destPath := buildFilePath(destDir, td.srcName)

		err := CopyWithProgress(srcPath, destPath, td.limit, td.offset)
		if td.hasError && err == nil {
			t.Errorf("copying file completed without error, but expected: %s", name)
			continue
		}
		if !td.hasError && err != nil {
			t.Errorf("copying file completed unexpected with error %v, but expected: %s", err, name)
			continue
		}

		if td.hasError {
			continue
		}

		expected, err := ioutil.ReadFile(srcPath)
		if err != nil {
			t.Errorf("source file reading completed with error %v: %s", err, name)
			continue
		}

		actual, err := ioutil.ReadFile(destPath)
		if err != nil {
			t.Errorf("destination file reading completed with error %v: %s", err, name)
			continue
		}

		if !bytes.Equal(expected, actual) {
			t.Errorf("destination file is different from the source: %s", name)
			continue
		}
	}
}
