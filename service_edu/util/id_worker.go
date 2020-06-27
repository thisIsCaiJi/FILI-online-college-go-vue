package util

import (
	idworker "github.com/gitstliu/go-id-worker"
	"strconv"
)

func GetId() (string,error) {
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	id,err := currWoker.NextId()
	if err!=nil {
		return "",err
	}
	return strconv.FormatInt(id,10),nil
}
