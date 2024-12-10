package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func swap(disk []int, i int, j int) {
	temp := disk[i]
	disk[i] = disk[j]
	disk[j] = temp
}

type Chunk struct {
	startIdx int
	size     int
	fileId   int
}

func main() {
	log.SetFlags(0)

	filename := "day9/input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("failed to read line from file")
	}
	line := scanner.Text()

	fileId := 0
	isFile := true
	// 0 = empty space
	// 1 = id of value 0
	// 2 = id of value 1
	// etc.
	diskIdx := 0
	var used []Chunk
	var free []Chunk

	for _, c := range line {
		num, _ := strconv.ParseInt(string(c), 10, 64)
		if isFile {
			used = append(used, Chunk{diskIdx, int(num), fileId})
			fileId++
		} else {
			free = append(free, Chunk{diskIdx, int(num), 0})
		}
		diskIdx += int(num)
		isFile = !isFile
	}

	for i := len(used) - 1; i >= 0; i-- {
		for j := 0; j < len(free); j++ {
			if free[j].startIdx > used[i].startIdx {
				break
			}
			if free[j].size < used[i].size {
				continue
			}
			// move used[i] into free[j]
			used[i].startIdx = free[j].startIdx
			free[j].size -= used[i].size
			free[j].startIdx += used[i].size
		}
	}

	total := 0

	for _, chunk := range used {
		for i := 0; i < chunk.size; i++ {
			total += (chunk.startIdx + i) * chunk.fileId
		}
	}

	log.Println("Total is", total)
}

