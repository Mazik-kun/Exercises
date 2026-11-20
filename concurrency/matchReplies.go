package main

import (
	"errors"
	"slices"
)

type Reply struct {
	ID    string
	Value int
}

func MatchReplies(requestIDs []string, in <-chan Reply) ([]Reply, error) {
	var err error
	err = nil
	reply := make([]Reply, len(requestIDs))
	lenIn := 0
	for v := range in {
		if len(requestIDs) == 0{
			err = errors.New("Ожидался пустой канал")
		}
		lenIn++
		if err == nil{
			idx := slices.Index(requestIDs, v.ID)
			if idx != -1{
				
				if slices.ContainsFunc(reply, func(r Reply)bool{return r.ID == v.ID}){
					err = errors.New("повтор Id")
					reply = nil
				}else{
					reply[idx]= Reply{v.ID, v.Value}
				}

			}else{
				err = errors.New("неизвестный ID")
				reply = nil
			} 

		}
	}
	if err == nil && len(requestIDs) > lenIn{
		reply = nil
		err = errors.New("Не хватает ответов")
	}

	return reply, err

	
}

