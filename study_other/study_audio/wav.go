package main

import (
	"fmt"
	"github.com/go-audio/wav"
	"os"
)

func splitAudio() {
	file, err := os.Open("/Users/weekend/Downloads/misaligned-chunk.wav")
	if err != nil {
		panic("open file err: " + err.Error())
	}

	decoder := wav.NewDecoder(file)
	decoder.ReadInfo()
	byteCount := decoder.AvgBytesPerSec * 200 / 1000 // 200ms字节长度

	var audioDataAll [][]byte
	for {
		audioDate := make([]byte, byteCount)
		n, err := file.Read(audioDate)
		if err != nil {
			break
		}
		audioDataAll = append(audioDataAll, audioDate[:n])
	}

	fmt.Println(len(audioDataAll))
}

func main() {
	splitAudio()
}
