package internal

import (
	"comic_server/interfaces"
	"comic_server/pb"
	"io/ioutil"
	"strconv"
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

	if nil == ptr.chapters {
		ptr.chapters = make([]*ChapterImp,0)
	}

	if nil == ptr.chapterList {
		ptr.chapterList = &pb.PbChapterList{}
	}

	var chapters []*pb.PbChapterInfo = make([]*pb.PbChapterInfo, 0)
	var defaultOneChapterPagesCount = 30
	var pathList = make([]string,0)
	var nameList = make([]string,0)
	var indexForJpgs = 0

	for _, info := range readerInfo {
		if info == nil {
			continue
		}

		if !info.IsDir() {
			pathList = append(pathList,ptr.FilePath + "/" + info.Name())
			nameList = append(nameList,info.Name())

			if defaultOneChapterPagesCount <= len(pathList) {
				chapter := &ChapterImp{
					FileName: "自动" + strconv.Itoa(indexForJpgs),
					FilePath: ptr.FilePath + "/" + "自动" + strconv.Itoa(indexForJpgs),
				}

				chapter.InitForList(pathList,nameList)

				pathList = make([]string,0)
				nameList = make([]string,0)

				log.Print("Added File ",info.Name()," to chapters\n")
				log.Print("PathList size:",len(pathList),"\n")

				ptr.chapters = append(ptr.chapters,chapter)

				if nil == chapter.GetChapterInfo() {
					log.Fatal("Has Nil Elment: name is ",ptr.FilePath)
				}
		
				chapters = append(chapters, chapter.GetChapterInfo())
			}

			indexForJpgs++

			continue
		}

		chapter := &ChapterImp{
			FileName: info.Name(),
			FilePath: ptr.FilePath + "/" + info.Name(),
		}

		print("Walk ", ptr.FilePath + "/" + info.Name()," chapter\n")

		chapter.Init()

		ptr.chapters = append(ptr.chapters,chapter)

		chapters = append(chapters, chapter.GetChapterInfo())
	}

	if len(pathList) > 0 {
		chapter := &ChapterImp{
			FileName: "自动" + strconv.Itoa(indexForJpgs),
			FilePath: ptr.FilePath + "/" + "自动" + strconv.Itoa(indexForJpgs),
		}

		chapter.InitForList(pathList,nameList)

		ptr.chapters = append(ptr.chapters,chapter)

		if nil == chapter.GetChapterInfo() {
			log.Fatal("Has Nil Elment: name is ",ptr.FilePath)
		}

		chapters = append(chapters, chapter.GetChapterInfo())
	}

	// for _,info := range chapters {
	// 	log.Print("Current chapter: ",info)
	// }

	ptr.chapterList.Chapters = chapters

	return true
}

func (ptr *BookImp) GetBookInfo() (p *pb.PbBookInfo) {
	return ptr.info
}

func (ptr *BookImp) GetChapterList() (p *pb.PbChapterList) {
	return ptr.chapterList
}

func (ptr *BookImp) GetChapter(index int) (p interfaces.Chapter) {
	if index >= len(ptr.chapters) || 0 > index {
		log.Fatal("Out range,size is ",len(ptr.chapters),"get index is ",index)
		return nil
	}

	return ptr.chapters[index]
}