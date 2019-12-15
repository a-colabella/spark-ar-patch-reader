package main

import (
  arputil "patches/arputil"
  "fmt"
  "os"
)

func main() {
  patch := arputil.ReadArp(os.Args[1])
  fmt.Println(arputil.GetName(patch))
  fmt.Println("Number of Input ports:", arputil.CountInputPorts(patch))
  fmt.Println("Number of Output ports:", arputil.CountOutputPorts(patch))
  fmt.Println("Number of Patches:", arputil.CountPatches(patch))
  fmt.Println("Number of Connections:", arputil.CountConnections(patch))
  fmt.Println("Input ports:", arputil.InputPorts(patch))
  fmt.Println("Output ports:", arputil.OutputPorts(patch))
}
