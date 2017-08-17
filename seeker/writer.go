package seeker

import (
	"fmt"
	"os"

	"github.com/paulsevere/p2p/manifest"
	"github.com/paulsevere/p2p/util"
)

func (s Seeker) WriteSegment(seg int, data []byte) {
	fmt.Printf("%v", data)
	util.CreateFile(s.OutPath)
	file, _ := os.OpenFile(s.OutPath, 0666, os.ModePerm)
	_, err := file.WriteAt(data, int64(seg*manifest.LEN))
	if err != nil {
		println("Error Writing file", err.Error())
	}
}
