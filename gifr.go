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

func usage() {
	fmt.Println("usage: gifr -f path/to/source.gif [-o path/to/dest.gif]")
	fmt.Println("\t-h show help")
	fmt.Println("\t-v show vesion")
}

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
		usage()
		os.Exit(0)
	}

	if sourceFilePath == "" {
		usage()
		os.Exit(1)
	}

	if destFilePath == "" {
		destFilePath = defaultDestGifPath
	}

	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		panic(err)
	}

	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	destFile.Seek(0, 0)
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

	// delays := make([]int, len(sourceGif.Delay))
	// for idx, delay := range sourceGif.Delay {
	// 	delays[len(delays)-1-idx] = delay
	// }
	// sourceGif.Delay = delays

	// disposals := make([]byte, len(sourceGif.Disposal))
	// for idx, disposal := range sourceGif.Disposal {
	// 	disposals[len(disposals)-1-idx] = disposal
	// }
	// sourceGif.Disposal = disposals

	gif.EncodeAll(destFile, sourceGif)

	sourceFile.Close()
	destFile.Close()
}
