package main

import (
	"flag"
	"fmt"
	"image/gif"
	"os"
)

var name = "gifr"
var version = "0.1.0"
var commitHash = "unknown"

var defaultDestGifPath = "./out.gif"

func main() {
	var sourceFilePath, destFilePath string
	var showVersion, showHelp bool

	flag.StringVar(&sourceFilePath, "f", "", "specify source gif file")
	flag.StringVar(&destFilePath, "o", "", "specify dest gif file")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")

	flag.Parse()

	if showVersion {
		fmt.Println("gifr", version, commitHash)
		os.Exit(0)
	}

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if sourceFilePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if destFilePath == "" {
		destFilePath = defaultDestGifPath
	}

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	sourceGif, err := gif.DecodeAll(sourceFile)
	if err != nil {
		panic(err)
	}

	var l int

	l = len(sourceGif.Image)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		sourceGif.Image[i], sourceGif.Image[j] = sourceGif.Image[j], sourceGif.Image[i]
	}

	l = len(sourceGif.Delay)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		sourceGif.Delay[i], sourceGif.Delay[j] = sourceGif.Delay[j], sourceGif.Delay[i]
	}

	l = len(sourceGif.Disposal)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		sourceGif.Disposal[i], sourceGif.Disposal[j] = sourceGif.Disposal[j], sourceGif.Disposal[i]
	}

	gif.EncodeAll(destFile, sourceGif)
}
