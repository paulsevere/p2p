package seeker

import (
	"encoding/gob"
	"net"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/paulsevere/go-p2p/manifest"
)

//
type Seeker struct {
	Manifest *manifest.Manifest
	Conn     *net.Conn
	OutPath  string
}

func New(hostname string, outpath string) Seeker {
	conn, err := net.Dial("tcp", hostname)
	if err != nil {
		println(err.Error())
		panic(err)
	}
	dec := gob.NewDecoder(conn)
	m := new(manifest.Manifest)
	dec.Decode(m)
	return Seeker{Manifest: m, OutPath: outpath}
}

func (s Seeker) SeekAll() {
	for i := range s.Manifest.Segments {
		println(i)
		// s.Seek(i)
	}
}

func (s Seeker) Seek(n int) {
	targ := strconv.Itoa(n)
	conn, _ := net.Dial("tcp", "localhost:8081")
	spew.Dump(conn)
	conn.Write([]byte(targ))
	buff := make([]byte, manifest.LEN)
	conn.Read(buff)
	s.WriteSegment(n, buff)

}
