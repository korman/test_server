package internal

import (
	"comic_server/pb"
	"io/ioutil"
	"log"
)

type ChapterImp struct {
	info *pb.PbChapterInfo
	pageList *pb.PbPageList
	pages []*PageImp
	FilePath string
	FileName string
}

func (ptr *ChapterImp) Init() {
	ptr.info = &pb.PbChapterInfo{}
	ptr.info.Name = ptr.FileName

	readerInfo, err := ioutil.ReadDir(ptr.FilePath)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if nil == ptr.pages {
		ptr.pages = make([]*PageImp,0)
	}

	if nil == ptr.pageList {
		ptr.pageList = &pb.PbPageList{}
	}

	var pages []*pb.PbPageInfo = make([]*pb.PbPageInfo, 0)

	for _, info := range readerInfo {
		if info == nil {
			continue
		}

		if info.IsDir() {
			continue
		}

		page := &PageImp{
			FileName: info.Name(),
			FilePath: ptr.FilePath + "/" + info.Name(),
		}

		page.Init()

		ptr.pages = append(ptr.pages,page)
		pages = append(pages, page.GetPageInfo())
	}

	ptr.pageList.Pages = pages
}

func (ptr *ChapterImp) InitForList (paths []string,names []string) {
	ptr.info = &pb.PbChapterInfo{}
	ptr.info.Name = ptr.FileName

	if len(paths) != len(names) {
		log.Fatal("paths count != names count")
		return
	}

	if nil == ptr.pages {
		ptr.pages = make([]*PageImp,0)
	}

	if nil == ptr.pageList {
		ptr.pageList = &pb.PbPageList{}
	}

	var pages []*pb.PbPageInfo = make([]*pb.PbPageInfo, 0)

	for index,info := range paths {
		page := &PageImp{
			FileName: names[index],
			FilePath: info,
		}

		page.Init()

		ptr.pages = append(ptr.pages,page)
		pages = append(pages, page.GetPageInfo())
	}

	ptr.pageList.Pages = pages
}

func (ptr *ChapterImp) GetChapterInfo() (p *pb.PbChapterInfo) {
	return ptr.info
}

func (ptr *ChapterImp) GetPageList() (p *pb.PbPageList) {
	return ptr.pageList
}