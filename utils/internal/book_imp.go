package internal

import "comic_server/pb"

type BookImp struct {
	info     *pb.PbBookInfo
	FileName string
	FilePath string
}

func (ptr *BookImp) Init() (result bool) {
	if 0 == len(ptr.FileName) {
		return false
	}

	ptr.info = &pb.PbBookInfo{}

	ptr.info.Name = ptr.FileName

	return true
}

func (ptr *BookImp) GetBookInfo() (p *pb.PbBookInfo) {
	return ptr.info
}
