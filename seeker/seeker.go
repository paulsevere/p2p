package seeker

import (
	"encoding/gob"
	"net"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/paulsevere/p2p/manifest"
	"github.com/paulsevere/p2p/msg"
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
	conn, _ := net.Dial("tcp", "localhost:8081")
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)
	missing := make([]int, 0)
	for i := range s.Manifest.Segments {
		missing = append(missing, i)
		// s.Seek(i)
	}
	enc.Encode(msg.Msg(missing))
	for {
		b := msg.Wrt{}
		err := dec.Decode(&b)
		if err != nil {
			println(err.Error())
			break
		}
		if b.Seg == -1 {
			println("Read Finished")
			break
		}
		s.WriteSegment(b.Seg, b.Content)
	}
}

func NewFromFile(manpath string, outpath string) Seeker {
	m := manifest.ReadFromFile(manpath)

	return Seeker{Manifest: &m, OutPath: outpath}
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
