package hmtype

import (
	"encoding/json"
	"testing"
)

type TJson struct {
	A Int64 `json:"a"`
}

func TestInt64_UnmarshalJSON(t *testing.T) {
	s := `{"a":"123123"}`
	var j *TJson
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(j.A)
	rj, err := json.Marshal(j)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(string(rj))
}
