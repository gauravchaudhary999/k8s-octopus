package octopus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type HTTP_METHOD int

const (
	GET HTTP_METHOD = iota
	POST
	PUT
	DELETE
)

func (method HTTP_METHOD) String() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[method]
}

type OctopusHost struct {
	host  string
	token string
}

type Client struct {
	host         string
	token        string
	api          string
	method       HTTP_METHOD
	body, target interface{}
}

func newClient(host string, token string, api string, method HTTP_METHOD, body interface{}, target interface{}) *Client {
	return &Client{host: host,
		token:  token,
		api:    api,
		method: method,
		body:   body,
		target: target}
}

func prepareBody(body interface{}) *bytes.Buffer {
	if body != nil {
		b, err := json.Marshal(body)

		if err != nil {
			log.Fatalln(err)
		}
		return bytes.NewBuffer(b)
	} else {
		var emptyBody []byte
		return bytes.NewBuffer(emptyBody)
	}
}

// func (client *Client) executeRaw() {
// 	var body string
// 	if filepath, ok := client.body.(string); ok {
// 		body = "data=@" + filepath
// 	} else {
// 		panic("Invalid filepath")
// 	}

// 	url := client.host + client.api
// 	tokenAppenedWithSingleQuote := client.token + "'"
// 	cmd := strings.Join([]string{"curl -X", client.method.String(), url, "-H", "'accept: application/json'", "-H", "'x-octopus-apikey:", tokenAppenedWithSingleQuote, "-F", body}, " ")
// 	out, err := exec.Command("/bin/sh", "-c", cmd).Output()

// 	if err != nil {
// 		panic(err)
// 	}

// 	res, _ := ioutil.ReadAll(NewCurlResponseReader(out))
// 	fmt.Println(string(res))
// }

func (client *Client) executeRaw() {
	url := client.host + client.api
	var path string
	if p, ok := client.body.(string); ok {
		path = p
	} else {
		panic("Invalid filepath")
	}

	file, err := os.Open(path)
	if err != nil {
		panic("Error in processing file")
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", path)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, file); err != nil {
		return
	}

	writer.Close()
	timeout := time.Duration(60 * time.Second)

	httpClient := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(client.method.String(), url, body)
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("X-Octopus-ApiKey", client.token)

	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	log.Println("Response Status: " + response.Status)

	if err != nil {
		panic(err)
	}
	json.NewDecoder(response.Body).Decode(client.target)
}

func (client *Client) execute() error {

	fmt.Println("Calling " + client.method.String() + " " + client.host + client.api)

	timeout := time.Duration(20 * time.Second)

	httpClient := http.Client{
		Timeout: timeout,
	}

	url := client.host + client.api

	request, err := http.NewRequest(client.method.String(), url, prepareBody(client.body))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "applictaion/json")
	request.Header.Set("X-Octopus-ApiKey", client.token)

	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	log.Println("Response Status: " + response.Status)
	// res, err := ioutil.ReadAll(response.Body)
	// log.Println("------------------------------------------")
	// log.Println(string(res))
	return json.NewDecoder(response.Body).Decode(client.target)
}

type CurlResponseReader struct {
	src []byte
	pos int
}

func (r *CurlResponseReader) Read(dst []byte) (n int, err error) {
	n = copy(dst, r.src[r.pos:])
	r.pos += n
	if r.pos == len(r.src) {
		return n, io.EOF
	}
	return
}

func NewCurlResponseReader(b []byte) *CurlResponseReader { return &CurlResponseReader{b, 0} }
