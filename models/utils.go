package models

import (
	"github.com/klauspost/compress/zstd"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (d HotelDump) Download(dir string) (filePath string, err error) {
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		log.Fatal(err)
		return
	}
	log.Println("Getting file information...")
	res, err := http.Head(*d.Data.URL)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Download started (%.2fM) %s\n", float32(res.ContentLength)/1048576, *d.Data.URL)
	res, err = http.Get(*d.Data.URL)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(res.Body)

	filePath = filepath.Join(dir, filepath.Base(*d.Data.URL))

	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(out *os.File) {
		err = out.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(out)

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("The downloaded file has been saved to %s\n", filePath)
	return
}

func (d HotelDump) Decompress(dir string) (newFilePath string, err error) {
	filePath, err := d.Download(dir)
	if err != nil {
		log.Fatal(err)
		return
	}

	newFilePath = filePath[0 : len(filePath)-len(filepath.Ext(filePath))]

	log.Printf("Archive unpacking started %s\n", filePath)
	zst, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	r, err := zstd.NewReader(zst)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer r.Close()

	out, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = io.Copy(out, r)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("The unpacked file has been saved to %s\n", newFilePath)

	err = os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Archive %s has been deleted\n", filePath)
	return
}
