package broadcaster

import (
	"encoding/gob"
	"net"
	"strconv"

	"github.com/paulsevere/go-p2p/manifest"
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

		go func() {
			conn, _ := b.FileServer.Accept()
			println("Request Handled")
			buff := make([]byte, 10)
			conn.Read(buff)
			num := string(buff)
			println(num)
			i, _ := strconv.Atoi(num)
			segment := b.ReadSegment(i)
			conn.Write(segment)
		}()
	}
}
