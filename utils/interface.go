package utils

import "comic_server/pb"

type Book interface {
	GetBookInfo() (p *pb.PbBookInfo)
}
