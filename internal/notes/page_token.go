package notes

import (
	"encoding/base64"
	"strconv"
)

func EncodePageToken(id int) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id)))
}

func DecodePageToken(token string) int {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return -1
	}
	id, err := strconv.Atoi(string(data))
	if err != nil {
		return -1
	}
	return id
}
