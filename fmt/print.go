package fmt

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v any) (err error) {
	s, err := jsonMarshal(v)
	if err != nil {
		return err
	}
	_, err = fmt.Println(s)
	return err
}

func jsonMarshal(v any) (s string, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
