package importer

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"unicode/utf8"
)

type Reader interface {
	ReadPDF(io.Reader) (PDF, error)
}

func NewReader() Reader {
	return reader{}
}

type reader struct{}

func (r reader) ReadPDF(in io.Reader) (PDF, error) {

	bufReader := bufio.NewReader(in)
	header, _, _ := bufReader.ReadLine()

	nihongo := string(header)
	fmt.Println(nihongo)
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts1 at byte position %d\n", runeValue, index)
	}
	nihongo = "%PDF-1.4"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	fmt.Println(string(header))
	fmt.Printf("%q\n", header)

	b, err := ioutil.ReadAll(in)
	if err != nil {
		return PDF{}, err
	}

	return PDF{
		header: string(header),
		file:   b,
	}, nil
}
