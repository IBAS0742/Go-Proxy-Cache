package MDB

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetHash(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
