package writer

type Writer struct {
	Path string
	Q    chan Wrt
}

func New(path string) {
	return Writer{Path: path, Q: make(chan Wrt)}
}
