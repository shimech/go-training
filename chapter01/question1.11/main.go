package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	logs := []string{}
	for range os.Args[1:] {
		log := <-ch
		fmt.Println(log)
		logs = append(logs, fmt.Sprintf("%s\n", log))
	}
	filename := fmt.Sprintf("%s.log", time.Now().Format("01_02_15_04_05"))
	writeLog(filename, logs)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func writeLog(filename string, lines []string) error {
	filepath := fmt.Sprintf("./log/%s", filename)
	file, err := os.Create(filepath)
	if err != nil {
		defer file.Close()
		return err
	}
	for _, line := range lines {
		b := []byte(line)
		_, err := file.Write(b)
		if err != nil {
			defer file.Close()
			return err
		}
	}
	return nil
}
