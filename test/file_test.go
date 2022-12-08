package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 分片大小
const chunkSize = 1 * 1024 * 1024  // 50 MB


// 文件分片
func TestGenerateChunkFile(t *testing.T) {
	fileIngo, err := os.Stat("./video/rust.pdf")
	if err != nil {
		t.Fatal(err)
	}

	chunkNum := math.Ceil(float64(fileIngo.Size()) / float64(chunkSize))  // 向上取整
	fd, err := os.OpenFile("./video/rust.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		// 指定读取文件的起始位置
		fd.Seek(int64(i * chunkSize), 0)

		if chunkSize > fileIngo.Size() - int64(i * chunkSize) {
			b = make([]byte, fileIngo.Size() - int64(i * chunkSize))
		}

		fd.Read(b)

		f, err := os.OpenFile("./video/" + strconv.Itoa(i) + "_chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	
		f.Write(b)
		defer f.Close()
	}
	 
}



// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	fd, err := os.OpenFile("./video/test2.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}	
	defer fd.Close()

	fileIngo, err := os.Stat("./video/test.mp4")
	if err != nil {
		t.Fatal(err)
	}

	chunkNum := math.Ceil(float64(fileIngo.Size()) / float64(chunkSize))  // 向上取整
	
	for i := 0; i < int(chunkNum); i++ {
		f, err := os.OpenFile("./video/" + strconv.Itoa(i) + "_chunk",os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)  // 一次性读取转成 byte 数组
		if err != nil {
			t.Fatal(err)
		}

		fd.Write(b)
		f.Close()
	}
}

// 文件一致性校验
func TestConsistency(t *testing.T) {
	fd1, err := os.OpenFile("./video/test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}	
	defer fd1.Close()

	b1, err := ioutil.ReadAll(fd1)
	if err != nil {
		t.Fatal(err)
	}	

	fd2, err := os.OpenFile("./video/test2.mp4", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}	
	defer fd1.Close()

	b2, err := ioutil.ReadAll(fd2)
	if err != nil {
		t.Fatal(err)
	}	
	str1 := fmt.Sprintf("%x", md5.Sum(b1))
	str2 := fmt.Sprintf("%x", md5.Sum(b2))
	t.Log(str1, str2)
	assert.Equal(t, true, str1 == str2, "应返回 true")
}