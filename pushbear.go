package pushbear

import (
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type PushBearStruct struct {
	sendKey string
}

type ResStruct struct {
	//Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	//Created time.Time `json:"created"`
}

func (p *PushBearStruct) Send(title, content string) error {
	params := fmt.Sprintf("sendkey=%s&text=%s&desp=%s", p.sendKey, title, content)
	res, err := httpGet(params)
	if err != nil {
		return err
	}

	r := ResStruct{}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return err
	}
	if r.Message != "" {
		return errors.New(r.Message)
	}
	return nil
}

func New(key string) *PushBearStruct {
	pushbear := &PushBearStruct{}
	pushbear.sendKey = key
	return pushbear
}
