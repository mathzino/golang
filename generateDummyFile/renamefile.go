package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var exePath, err = os.Getwd()
var tempPath = filepath.Join(exePath,"chapter-A.59-pipeline-temp")

func main() {
    log.Println("start")
    start := time.Now()

    proceed()

    duration := time.Since(start)
    log.Println("done in", duration.Seconds(), "seconds")
}

func proceed() {
    counterTotal := 0
    counterRenamed := 0
    err := filepath.Walk(tempPath, 
        func(path string, info os.FileInfo, err error) error {
        
        // if there is an error, return immediatelly
        if err != nil {
            return err
        }

        // if it is a sub directory, return immediatelly
        if info.IsDir() {
            return nil
        }

        counterTotal++

    
        // sum it
        sum := fmt.Sprintf("%x", md5.Sum([]byte(info.Name())))

        // rename file
        destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
        // fmt.Println(destinationPath,"\n",info.Name())
        err = os.Rename(path, destinationPath)
        if err != nil {
            return err
        }

        counterRenamed++
        return nil
    })

    if err != nil {
        log.Println("ERROR:", err.Error())
    }

	    log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
	}

