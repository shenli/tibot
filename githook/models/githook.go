package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Githooks map[string]*Githook
)

type Githook struct {
	GithookId  string
	Score      int64
	PlayerName string
}

func init() {
	Githooks = make(map[string]*Githook)
	Githooks["hjkhsbnmn123"] = &Githook{"hjkhsbnmn123", 100, "astaxie"}
	Githooks["mjjkxsxsaa23"] = &Githook{"mjjkxsxsaa23", 101, "someone"}
}

func GitAddOne(object Githook) (GithookId string) {
	object.GithookId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Githooks[object.GithookId] = &object
	return object.GithookId
}

func GitGetOne(GithookId string) (object *Githook, err error) {
	if v, ok := Githooks[GithookId]; ok {
		return v, nil
	}
	return nil, errors.New("GithookId Not Exist")
}

func GitGetAll() map[string]*Githook {
	return Githooks
}

func GitUpdate(GithookId string, Score int64) (err error) {
	if v, ok := Githooks[GithookId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("GithookId Not Exist")
}

func GitDelete(GithookId string) {
	delete(Githooks, GithookId)
}
