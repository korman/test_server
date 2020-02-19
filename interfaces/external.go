package interfaces

import "comic_server/pb"

type Book interface {
	GetBookInfo() (p *pb.PbBookInfo)
	GetChapterList() (p *pb.PbChapterList)
	GetChapter(index int) (p Chapter)
}

type Chapter interface {
	GetChapterInfo() (p *pb.PbChapterInfo)
	GetPageList() (p *pb.PbPageList)
}

type Page interface {
	GetPageInfo() (p *pb.PbPageInfo)
}