package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/hydra13142/chardet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

type resContent struct {
	Status  string `json:"Status"`
	Content string `json:"Content"`
}

type content struct {
	Content string `json:"Content"`
}

type file struct {
	FileName string `json:"FileName"`
	Data     string `json:"Data"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/simplifiedGarbleds", simplifiedGarbleds).Methods("Post")
	router.HandleFunc("/simplifiedGarbled", simplifiedGarbled).Methods("Post")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func simplifiedGarbled(w http.ResponseWriter, r *http.Request) {
	var newContent content
	var response resContent

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error!!")
	}

	response.Status = "false"
	response.Content = ""
	if json.Unmarshal(reqBody, &newContent) != nil {
		json.NewEncoder(w).Encode(response)
	} else {

		var bytes = []byte(newContent.Content)
		t, _ := decodeGBK(bytes)
		fmt.Println(string(t))
		if !strings.Contains(chardet.Mostlike(bytes), "utf") &&
			!strings.Contains(chardet.Mostlike(bytes), "utf16") {

			response.Status = "true"

			// GBK 2 UTF-8
			s, _ := decodeGBK(bytes)
			bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
			data := string(bomUtf8) + string(s)

			response.Content = data
		}

		json.NewEncoder(w).Encode(response)
	}
}

func simplifiedGarbleds(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

//convert GBK to UTF-8
func decodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert UTF-8 to GBK
func encodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert BIG5 to UTF-8
func decodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//convert UTF-8 to BIG5
func encodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
