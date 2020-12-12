package utils

import (
  "log"
  "os"
)

func ReadEnv(name string) string {
  value, present := os.LookupEnv(name)
  if !present {
    log.Fatal(name + " not present")
  }

  return value
}

