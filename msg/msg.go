package msg

type Message struct {
	Segs []int
}

type Wrt struct {
	Content []byte
	Seg     int
}

func Msg(segs []int) Message {
	return Message{
		Segs: segs,
	}
}

func Content(n int, content []byte) Wrt {
	return Wrt{
		Content: content,
		Seg:     n,
	}
}

func Done() Wrt {
	return Wrt{Seg: -1}
}
