package utils

import (
	"fmt"
	"strconv"
)

func ConvertAnyToUint(v any) (*uint, error) {
	u64, err := strconv.ParseUint(fmt.Sprintf("%v", v), 10, 32)
	if err != nil {
		return nil, err
	}
	u := uint(u64)
	return &u, nil
}
