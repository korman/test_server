package internal

import (
	"bytes"
	"comic_server/pb"
	"encoding/base64"
	"io"
	"log"
	"os"
)

type PageImp struct {
	info *pb.PbPageInfo
	FilePath string
	FileName string
}

func (ptr *PageImp) Init() {
	ptr.info = &pb.PbPageInfo{}

	file,err := os.OpenFile(ptr.FilePath,os.O_RDONLY, os.ModePerm)

	if nil != err {
		log.Fatal("打开文件失败:",err.Error())
		return
	}

	fileInfo,err := file.Stat()

	if nil != err {
		log.Print("Get File Stat Failed")
		return
	}

	defer file.Close()

	var buffer bytes.Buffer
	io.CopyN(&buffer, file,fileInfo.Size())

	bytes := buffer.Bytes()

	encodeString := base64.StdEncoding.EncodeToString(bytes)
	log.Print("Encoding image data size:",len(encodeString))

	ptr.info.Data = encodeString
}

func (ptr *PageImp) GetPageInfo() (p *pb.PbPageInfo) {
	return ptr.info
}