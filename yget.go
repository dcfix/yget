package main

import (
	"flag"
	"io"
	"net/http"
	"os"
        "fmt"
        "time"
)

func main() {
  uri := flag.String("Uri", , "the web page to grab")
  filename := flag.String("filename", , "The name of the file to save to")
  delay := flag.Int("Delay", 0, "Number of seconds to pause before retrieving the web page")

  flag.Parse()

  duration := time.Duration(*delay) * time.Second 
  time.Sleep(duration)

  err := downloadFile(*filename, *uri)
  if (err != nil) {
    fmt.Println("error:", err)
}

}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
