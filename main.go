package main

import (
  arputil "patches/arputil"
  "fmt"
  "os"
)

func main() {
  patch := arputil.ReadArp(os.Args[1])
  fmt.Println(arputil.GetName(patch))
  fmt.Println("Input ports:", arputil.CountInputPorts(patch))
  fmt.Println("Output ports:", arputil.CountOutputPorts(patch))
  fmt.Println("Patches:", arputil.CountPatches(patch))
  fmt.Println("Connections:", arputil.CountConnections(patch))

}
