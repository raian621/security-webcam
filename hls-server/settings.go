package main

import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"regexp"
)

type WebcamFormat struct {
	FormatName     string
	WebcamSettings []WebcamSettings
}

type WebcamSettings struct {
	Width      int
	Height     int
	FrameRates []float32
}

func getVideoSettingsOutput(videoDevicePath string) ([]byte, error) {
	cmd := exec.Command("v4l2-ctl", "--list-formats-ext", "--device", videoDevicePath)
	output, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}

	return output, nil
}

func parseVideoSettingsOutput(videoSettingsOutput []byte) ([]WebcamFormat, error) {
	reader := bufio.NewReader(bytes.NewReader(videoSettingsOutput))

	formatRegex := regexp.MustCompile(`^\s+\[\d\]: '(\w+)' \((.*)\)$`)
	sizeRegex := regexp.MustCompile(`^\s+Size:.+ (\d+)x(\d+)$`)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}

		log.Println(string(line))

		if submatches := formatRegex.FindSubmatch(line); len(submatches) == 3 {
			log.Println("--------------------------------------------")
			log.Println("format line detected")
			log.Println("format name:     ", string(submatches[1]))
			log.Println("format long name:", string(submatches[2]))
			log.Println("--------------------------------------------")
		}

		if submatches := sizeRegex.FindSubmatch(line); len(submatches) == 3 {
			log.Println("--------------------------------------------")
			log.Println("size line detected")
			log.Println("width: ", string(submatches[1]))
			log.Println("height:", string(submatches[2]))
			log.Println("--------------------------------------------")
		}

	}

	return []WebcamFormat{}, nil
}
