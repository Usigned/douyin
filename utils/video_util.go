package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func ReadFrameAsJpeg(videoFilePath string, frameNum int, imgFilePath string) int {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoFilePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatalln("ffmpeg convert video failed!")
		return -1
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatalln("reader decode failed!")
		return -1
	}
	err = imaging.Save(img, imgFilePath)
	if err != nil {
		log.Fatalln("save img failed!")
		return -1
	}
	return 1
}
