/*
Zpl2img - Creates png image from zpl code

go build zpl2img.go
zpl2img image.png "^xa^cfa,50^fo100,100^fd Hello World ^fs^xz"
*/

package main

import (
    "os"
    "io"
	"io/ioutil"
	"log"
    "bytes"
    "net/http"
)

func main() {

	filename := os.Args[1]
	zplcode  := os.Args[2]

    zpl := []byte(zplcode)
    req, err := http.NewRequest("POST", "http://api.labelary.com/v1/printers/8dpmm/labels/4x6/0/", bytes.NewBuffer(zpl))
    if err != nil {
        log.Fatalln(err)
	}
	
    /* req.Header.Set("Accept", "application/pdf" PDF türünde çıktı almak için */

    client := &http.Client{}
    response, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer response.Body.Close()

    if response.StatusCode == http.StatusOK {
        file, err := os.Create(filename)
        if err != nil {
            log.Fatalln(err)
        }
        defer file.Close()
        io.Copy(file, response.Body)
    } else {
        body, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatalln(err)
        }
        log.Fatalln(string(body))
    }
}