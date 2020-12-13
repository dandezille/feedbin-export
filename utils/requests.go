package utils

import (
  "bytes"
  "io"
)

func BodyFromString(data string) io.Reader {
  if data == "" {
    return nil
  } else {
    return bytes.NewBuffer([]byte(data))
  }
}
