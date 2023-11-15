package main

import (
	"net/http"
	"strconv"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"

	"github.com/fermyon/spin/sdk/go/v2/kv"
)

var (
	countStr string
	count    int32
	countKey string
)

func main() {
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, req *http.Request) {
		countStr = "0"
		count = 0
		countKey = "count"

		store, err := kv.OpenStore("default")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer store.Close()
		countStr = func() string {
			v, err := store.Get(countKey)
			if err != nil {
				return "0"
			}
			return string(v)
		}()

		count, _ = func() (int32, error) {
			i, err := strconv.Atoi(countStr)
			return int32(i), err
		}()
		countStr = strconv.Itoa(int(count + 1))
		store.Set(countKey, []byte(countStr))

		w.Write([]byte(countStr))
	})
}

func empty() {
}