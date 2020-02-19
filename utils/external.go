package utils

import (
	"comic_server/interfaces"
	"comic_server/pb"
	"comic_server/utils/internal"
	"io/ioutil"
	"sync"
)

type ComicManager struct {
	books    []*internal.BookImp
	bookList *pb.PbBookList
}

var instance *ComicManager
var once sync.Once

func Instance() *ComicManager {
	once.Do(func() {
		instance = &ComicManager{}
	})
	return instance
}

func (ptr *ComicManager) GetBookList() (bookList *pb.PbBookList) {
	return ptr.bookList
}

func (ptr *ComicManager) GetBook(id int) (book interfaces.Book) {
	return ptr.books[id]
}

func (ptr *ComicManager) LoadComicFiles(rootPath string) {
	if nil == ptr.books {
		ptr.books = make([]*internal.BookImp, 0)
	}

	readerInfo, err := ioutil.ReadDir(rootPath)

	if err != nil {
		return
	}

	if nil == ptr.bookList {
		ptr.bookList = &pb.PbBookList{}
	}

	var books []*pb.PbBookInfo = make([]*pb.PbBookInfo, 0)

	for _, info := range readerInfo {
		if info == nil {
			continue
		}

		if !info.IsDir() {
			continue
		}

		book := &internal.BookImp{
			FileName: info.Name(),
			FilePath: rootPath + "/" + info.Name(),
		}

		print("Walk path ", rootPath+info.Name()," completed\n")

		book.Init()

		ptr.books = append(ptr.books,book)
		books = append(books, book.GetBookInfo())
	}

	ptr.bookList.Books = books
}
