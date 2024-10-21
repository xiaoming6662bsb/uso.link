package utils

import (
	"github.com/google/uuid"
	"strings"
	"uso.link/internal/utils/log2"
)

func GenerateUuids(count int, letterCase int) []string {
	uuids := []string{}
	for i := 0; i < count; i++ {
		if us, err := uuid.NewRandom(); !log2.IfErrPrt(err) {
			uss := us.String()
			if letterCase == 1 {
				uuids = append(uuids, strings.ToUpper(uss))
			} else {
				uuids = append(uuids, strings.ToLower(uss))
			}
		}

	}
	return uuids
}
