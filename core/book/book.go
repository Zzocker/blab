package book

import "github.com/Zzocker/blab/core/ports"

type bookCore struct {
	bStore ports.BookStorePort
}
