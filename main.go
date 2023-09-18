package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	if len(os.Args) < 2 {
		panic("Usage: videoEditing <integer>")
	}

	arg := os.Args[1]

	d, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	in := "./testdata/in1.mp4"
	out := "./outdir/out2.mp4"

	info, err := getVideoInfo(in)

	if err != nil {
		panic(err)
	}

	err = ffmpeg.
		Input(in, ffmpeg.KwArgs{"ss": info.Duration - float64(d)}).
		Output(out, ffmpeg.KwArgs{"t": d}).
		OverWriteOutput().
		Run()

	if err != nil {
		panic(err.Error())
	}
}

type StreamInfo struct {
	CodecType string  `json:"codec_type"`
	Duration  float64 `json:"duration"`
}

func (s *StreamInfo) UnmarshalJSON(data []byte) error {
	var temp struct {
		CodecType string `json:"codec_type"`
		Duration  string `json:"duration"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.CodecType = temp.CodecType

	duration, err := strconv.ParseFloat(temp.Duration, 64)
	if err != nil {
		return err
	}
	s.Duration = duration

	return nil
}

type MediaInfo struct {
	Streams []StreamInfo `json:"streams"`
}

func getVideoInfo(videoPath string) (StreamInfo, error) {
	i, err := ffmpeg.Probe(videoPath)
	if err != nil {
		return StreamInfo{}, err
	}

	var m MediaInfo

	b := []byte(i)

	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Println("Error: ", err)
	}

	return m.Streams[0], err
}
