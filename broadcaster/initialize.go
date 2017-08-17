package broadcaster

import (
	"encoding/gob"
	"net"

	"github.com/paulsevere/p2p/manifest"
	"github.com/paulsevere/p2p/msg"
)

// Broadcaster asds
type Broadcaster struct {
	Manifest   manifest.Manifest
	Server     net.Listener
	FileServer net.Listener
}

// Init asas
func Init(manifest manifest.Manifest) Broadcaster {
	ln, _ := net.Listen("tcp", ":8080")
	ln2, _ := net.Listen("tcp", ":8081")
	println("Now Listening!")

	return Broadcaster{Manifest: manifest, Server: ln, FileServer: ln2}
}

// StartRequestLoop asdfa
func (b Broadcaster) StartRequestLoop() {
	for {
		conn, _ := b.Server.Accept()
		println("Request Handled")
		go func() {
			dec := gob.NewEncoder(conn)

			dec.Encode(b.Manifest)

		}()
	}
}

func (b Broadcaster) FileRequests() {
	for {
		conn, _ := b.FileServer.Accept()
		go b.FileRequestHandler(conn)
		// go
	}
}

func (b Broadcaster) FileRequestHandler(conn net.Conn) {

	dec := gob.NewDecoder(conn)
	enc := gob.NewEncoder(conn)
	m := msg.Message{}
	dec.Decode(&m)
	for _, n := range m.Segs {
		ret := msg.Content(n, b.ReadSegment(n))
		enc.Encode(ret)
	}

	enc.Encode(msg.Done())

}
