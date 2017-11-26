package tindergo

import (
	"encoding/json"
)

type Position struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (t *TinderGo) Travel(position Position) error {
	url := "https://api.gotinder.com/passport/user/travel?locale=en"

	bReq, err := json.Marshal(position)
	if err != nil {
		return err
	}

	_, errs := t.requester.Post(url, string(bReq))
	if errs != nil {
		return errs[0]
	}
	return nil
}
