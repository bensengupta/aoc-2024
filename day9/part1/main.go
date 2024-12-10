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

	fileId := 1
	isFile := true
	// 0 = empty space
	// 1 = id of value 0
	// 2 = id of value 1
	// etc.
	var disk []int

	for _, c := range line {
		num, _ := strconv.ParseInt(string(c), 10, 64)
		if isFile {
			for i := 0; i < int(num); i++ {
				disk = append(disk, fileId)
			}
			fileId++
		} else {
			for i := 0; i < int(num); i++ {
				disk = append(disk, 0)
			}
		}
		isFile = !isFile
	}

	// left = index of free space
	// right = index of file chunk
	left := 0
	right := len(disk) - 1

	for {
		for disk[left] != 0 && left < right {
			left++
		}
		for disk[right] == 0 && left < right {
			right--
		}
		if left == right {
			break
		}
		swap(disk, left, right)
	}

	total := 0

	for i, n := range disk {
		if n == 0 {
			break
		}
		total += i * (n - 1)
	}

	log.Println("Total is", total)
}
