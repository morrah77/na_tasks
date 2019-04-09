package helpers

import (
	"errors"
	"encoding/json"
)

func ParseList(s string) (result [][]int, err error)  {
	err = json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil, errors.New(`helpers: incorrect input`)
	}
	return result, nil
}
