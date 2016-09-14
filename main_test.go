package test

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_main(t *testing.T) {
	data := `{"jsonrpc": "2.0", "id": 12345, "method": "Pinyin.Query" , "params":{"hans": ["测试", "汉字", "贝美互动", "123"]}}`
	reader := bytes.NewReader([]byte(data))
	resp, err := http.Post("http://localhost:4000/", "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}
