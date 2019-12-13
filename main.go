package main

import (
  arputil "patches/arputil"
  "fmt"
  "os"
)

func main() {
  fmt.Println(arputil.ReadArp(os.Args[1]))
}
