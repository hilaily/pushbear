package pushbear

import (
	"testing"
)

func TestPushBearStruct_Send(t *testing.T) {
	key := "your-key"
	push := New(key)

	res := push.Send("title11", "content1")
	t.Log(res)
}
