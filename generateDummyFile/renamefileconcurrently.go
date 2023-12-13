package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var exePath, err = os.Getwd()
var tempPath = filepath.Join(exePath,"chapter-A.59-pipeline-temp")

type FileInfo struct {
    FilePath  string // file location
    Content   []byte // file content
    Sum       string // md5 sum of content
    IsRenamed bool   // indicator whether the particular file is renamed already or not
}

func main() {
    log.Println("start")
    start := time.Now()

    // pipeline 1: loop all files and read it
    chanFileContent := readFiles()

    fmt.Println("pipeline 1 selesai")




	   // pipeline 3: rename files
	chanRename1 := rename(chanFileContent)
	chanRename2 := rename(chanFileContent)
	chanRename3 := rename(chanFileContent)
	chanRename4 := rename(chanFileContent)
	chanRename5 := rename(chanFileContent)
	chanRename6 := rename(chanFileContent)
	chanRename := mergeChanFileInfo(chanRename1, chanRename2, chanRename3, chanRename4,chanRename5,chanRename6)

	// print output
	counterRenamed := 0
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func readFiles() <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
            // if there is an error, return immediatelly
            if err != nil {
                return err
            }

            // if it is a sub directory, return immediatelly
            if info.IsDir() {
                return nil
            }

            // buf, err := os.ReadFile(path)
            // if err != nil {
            //     return err
            // }

            chanOut <- FileInfo{
                FilePath: path,
                
            }

            return nil
        })
        if err != nil {
            log.Println("ERROR:", err.Error())
        }

        close(chanOut)
    }()

    return chanOut
}

func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        for fileInfo := range chanIn {
            fileInfo.Sum = fmt.Sprintf("%x", md5.Sum([]byte(fileInfo.FilePath)))
            chanOut <- fileInfo
        }
        close(chanOut)
    }()

    return chanOut
}


func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
    wg := new(sync.WaitGroup)
    chanOut := make(chan FileInfo)

    wg.Add(len(chanInMany))
    for _, eachChan := range chanInMany {
        go func(eachChan <-chan FileInfo) {
            for eachChanData := range eachChan {
                chanOut <- eachChanData
            }
            wg.Done()
        }(eachChan)
    }

    go func() {
        wg.Wait()
        close(chanOut)
    }()

    return chanOut
}


func rename(chanIn <-chan FileInfo) <-chan FileInfo {
    chanOut := make(chan FileInfo)

    go func() {
        for fileInfo := range chanIn {
            newPath := filepath.Join(tempPath, fmt.Sprintf("file-%x.txt", md5.Sum([]byte(fileInfo.FilePath))))
            err := os.Rename(fileInfo.FilePath, newPath)
            fileInfo.IsRenamed = err == nil
            chanOut <- fileInfo
        }

        close(chanOut)
    }()

    return chanOut
}