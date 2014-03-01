package main

import (
    "crypto/md5"
    "encoding/hex"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func crack(hash string) {
    url := fmt.Sprintf("http://www.google.co.uk/search?sourceid=chrome&q=%s", hash)
    client := &http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.8 (KHTML, like Gecko) Chrome/17.0.938.0 Safari/535.8")
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    wordlist := strings.Split(string(body), " ")
    for _, word := range wordlist {
        nhash := md5.New()
        nhash.Write([]byte(word))
        sum := nhash.Sum(nil)
        if hex.EncodeToString(sum) == hash {
            fmt.Printf("%s:%s\n", word, hash)
            break
        }
    }
}

func main() {
    var file = flag.String("file", "", "file of hashes to crack, separated by \\n")
    flag.Parse()

    if file == nil {
        panic("Please provide a file of hashes")
    }

    b, err := ioutil.ReadFile(*file)
    if err != nil {
        panic(err)
    }

    contents := strings.Split(string(b), "\n")
    for _, c := range contents {
        if len(c) < 32 {
            continue
        } else {
            crack(c)
        }
    }
}