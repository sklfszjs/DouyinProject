package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func GetCoverImg(videoName string, picName string) {
	width := 1280
	height := 2276
	// cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", strconv.Itoa(index), "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	cmd := exec.Command("ffmpeg", "-i", videoName, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "image2", picName)

	if cmd.Run() != nil {
		panic("could not generate frame")
	}

}

func SaveImg(filename string, data *bytes.Buffer) {
	newfile, err := os.Create(filename)
	if err != nil {
		fmt.Println("image create error", err)
	}
	defer newfile.Close()
	newfile.Write(data.Bytes())

}
