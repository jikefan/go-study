package os

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	file, err := os.Open("./test.txt")

	defer func() {
		err = file.Close()
		if err != nil {
			t.Error("error in opening")
		}
	}()

	if err != nil {
		t.Error("open file err:", err)
	}

	t.Errorf("file = %v\n", file)

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		t.Errorf(str)
	}

	t.Error("文件读取结束")

}

func TestReadFile(t *testing.T) {
	file := "./test.txt"
	// not open and close
	content, err := os.ReadFile(file)

	if err != nil {
		t.Error(err)
	}

	t.Errorf("%s\n", content)
}

func TestFileExists(t *testing.T) {
	_, err := os.Stat("./test.txt")

	if err == nil {
		t.Error("文件存在")
	} else {
		if os.IsNotExist(err) {
			t.Error("文件或文件夹不存在")
		} else {
			t.Error("异常情况")
		}
	}

}
