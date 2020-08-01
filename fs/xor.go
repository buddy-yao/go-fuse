package fs

func xorEnc(buf []byte) {
	for i, v := range buf {
		buf[i] = v ^ 97
	}
}
