package importer

import "fmt"

type PDF struct {
	header string
	file   []byte
}

func (p *PDF) GetHeader() string {
	return fmt.Sprintf("%q", p.header)
}
