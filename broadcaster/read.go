package broadcaster

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/paulsevere/p2p/manifest"
)

func (b Broadcaster) ReadSegment(n int) []byte {
	path := b.Manifest.Name
	file, _ := os.Open(path)
	buff := make([]byte, manifest.LEN)
	readAt := int64(manifest.LEN * n)
	println("Reading at : ", readAt)
	file.ReadAt(buff, readAt)
	spew.Dump(buff)
	return buff

}
