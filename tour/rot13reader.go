package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	ctr := 0
	x, err := rot.r.Read(b)
	for i := range x {
		ctr++
		inted := int(b[i])
		if inted < 65 || (inted > 90 && inted < 97) || inted > 122 {
			b[i] = byte(inted)
			continue
		}
		converted := inted + 13
		if converted > 90 && converted < 97 {
			converted = 96 + converted%90
			b[i] = byte(converted)
			continue
		}
		if converted > 122 {
			converted = 64 + converted%122
			b[i] = byte(converted)
			continue
		}
		b[i] = byte(converted)
	}
	if err != nil {
		return ctr, io.EOF
	}
	return ctr, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
