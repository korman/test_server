package internal

import (
	"comic_server/pb"
	"io/ioutil"
	"log"
)

type BookImp struct {
	info     *pb.PbBookInfo
	chapterList *pb.PbChapterList
	chapters []*ChapterImp
	FileName string
	FilePath string
}

func (ptr *BookImp) Init() (result bool) {
	if 0 == len(ptr.FileName) {
		log.Fatal("FileName is null:",ptr.FileName)
		return false
	}

	ptr.info = &pb.PbBookInfo{}

	ptr.info.Name = ptr.FileName

	readerInfo, err := ioutil.ReadDir(ptr.FilePath)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if nil == ptr.chapterList {
		ptr.chapterList = &pb.PbChapterList{}
	}

	var chapters []*pb.PbChapterInfo = make([]*pb.PbChapterInfo, 0)

	for _, info := range readerInfo {
		if info == nil {
			continue
		}

		if !info.IsDir() {
			continue
		}

		chapter := &ChapterImp{
			FileName: info.Name(),
			FilePath: ptr.FilePath + "/" + info.Name(),
		}

		print("Walk ", ptr.FilePath + "/" + info.Name()," chapter\n")

		chapter.Init()

		chapters = append(chapters, chapter.GetChapterInfo())
	}

	ptr.chapterList.Chapters = chapters

	return true
}

func (ptr *BookImp) GetBookInfo() (p *pb.PbBookInfo) {
	return ptr.info
}

func (ptr *BookImp) GetChapterList() (p *pb.PbChapterList) {
	return ptr.chapterList
}