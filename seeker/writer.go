package seeker

import (
	"os"

	"github.com/paulsevere/p2p/manifest"
	"github.com/paulsevere/p2p/util"
)

func (s Seeker) WriteSegment(seg int, data []byte) {
	// fmt.Printf("%v\n", data)
	util.CreateFile(s.OutPath)
	file, err1 := os.OpenFile(s.OutPath, 0666, os.ModePerm)
	if err1 != nil {
		println(err1.Error())
	}
	_, err := file.WriteAt(data, int64(seg*manifest.LEN))
	if err != nil {
		println("Error Writing file", err.Error())
	}
	file.Close()
}
