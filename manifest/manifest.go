package manifest

import (
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"os"
)

const LEN int = 64

type bytes []byte

type Manifest struct {
	Name     string
	Segments [][16]byte
}

func New(path string) Manifest {
	m := Manifest{Name: path, Segments: make([][16]byte, 0)}
	file, _ := os.Open(path)
	buff := make(bytes, LEN)
	var n int
	var err error
	for i := 0; err == nil; i++ {
		n, err = file.ReadAt(buff, int64(i*LEN))
		if err != nil {
			println(err.Error())

		}
		println(n)

		m.Segments = append(m.Segments, md5.Sum(buff))
	}
	return m
}

func (m Manifest) WriteToFile(path string) {
	createFile(path)
	file, err := os.OpenFile(path, 0666, os.ModePerm)
	if err != nil {
		println(err.Error())
		return
	}
	enc := gob.NewEncoder(file)
	enc.Encode(m)

}

func ReadFromFile(path string) Manifest {
	file, _ := os.Open(path)
	dec := gob.NewDecoder(file)
	m := new(Manifest)
	dec.Decode(&m)
	return *m
}

func (m Manifest) Print() {
	fmt.Printf(`
	Name: %v
	Segment Length: %v
	First Segment: %v
	`, m.Name, len(m.Segments), m.Segments[0])
}
