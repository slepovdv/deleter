package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	pathDir := flag.String("path", ".", "Folder path")

	flag.Parse()

	err := filepath.Walk(*pathDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			stats, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}

			if stats.Mode().IsRegular() {
				if fileIsOld(stats.ModTime()) && isFirstDay(stats.ModTime()) {
					e := os.Remove(path)
					if e != nil {
						log.Fatal(e)
					}
					fmt.Printf("File %s is remove\nDate is %s\n\n", path, stats.ModTime())
				}
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

func fileIsOld(t time.Time) bool {
	return time.Since(t) > 24*time.Hour*30
}
func isFirstDay(t time.Time) bool {
	if time.Since(t) > 24*time.Hour*365 {
		return true
	}
	return t.Day() != 1

}
