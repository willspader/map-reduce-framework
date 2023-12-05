package main

import (
	"log"
	"os"
)

const (
	mapWorkers    int = 3
	reduceWorkers int = 3
)

func main() {
	file, err := os.Open("C:/Users/william_spader/go-projects/map-reduce-framework/urls.txt")
	if err != nil {
		log.Fatalln("Could not open specified file. ", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln("Could not obtain file's information. ", err)
	}

	// bytes for each mapWorkers read
	maxBytesRead := fileInfo.Size() / int64(mapWorkers)
	var offset int64 = 0
	for i := 0; i < mapWorkers; i++ {
		buffer := make([]byte, maxBytesRead)
		file.ReadAt(buffer, offset)
		log.Println(buffer)
		offset = offset + maxBytesRead
	}
}

func mapWorker() {

}

func reduceWorker() {

}
