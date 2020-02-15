package internal

import (
	"comic_server/pb"
)

type ChapterImp struct {
	info *pb.PbChapterInfo
	pageList *pb.PbPageList
	FilePath string
	FileName string
}

func (ptr *ChapterImp) Init() {
	ptr.info = &pb.PbChapterInfo{}
	ptr.info.Name = ptr.FileName
}

func (ptr *ChapterImp) GetChapterInfo() (p *pb.PbChapterInfo) {
	return ptr.info
}

func (ptr *ChapterImp) GetPageList() (p *pb.PbPageList) {
	return ptr.pageList
}