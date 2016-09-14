package pinyin

import (
	"log"
	"net/http"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

type Pinyin struct {
}

type Query_params struct {
	Hans []string `json:"hans"`
}

type Query_response struct {
	Pinyin []string `json:"pinyin"`
}

func (h *Pinyin) Query(r *http.Request, params *Query_params, reply *Query_response) error {
	log.Println("params:", params)

	for i, item := range params.Hans {
		log.Println(i, item)

		pinyin_list := pinyin.LazyPinyin(item, pinyin.NewArgs())
		pinyin_str := strings.Join(pinyin_list, " ")
		reply.Pinyin = append(reply.Pinyin, pinyin_str)
	}

	return nil
}
