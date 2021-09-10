package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	pathDir := flag.String("path", ".", "Path to the directory")
	days := flag.Int("days", 30, "Delete files older than n days")
	save := flag.String("save-month", "yes", "Save files created on the 1st of each month for 1 year")

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
				switch {
				case *save == "yes":
					if fileIsOld(stats.ModTime(), *days) && isFirstDay(stats.ModTime()) {
						deleter(path, stats)
					}
				default:
					if fileIsOld(stats.ModTime(), *days) {
						deleter(path, stats)
					}
				}
			}

			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

func fileIsOld(t time.Time, days int) bool {
	return time.Since(t) > 24*time.Hour*time.Duration(days)
}

func isFirstDay(t time.Time) bool {
	if time.Since(t) > 24*time.Hour*365 {
		return true
	}
	return t.Day() != 1

}

func deleter(p string, s fs.FileInfo) {
	e := os.Remove(p)
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("File %s is remove\nDate is %s\n\n", p, s.ModTime())
}
