/*
  This go file opens and returns .arp json data
*/
package arputil

import (
  "fmt"
  "io"
  "io/ioutil"
  "os"
  "archive/zip"
  "encoding/json"
)

// Error handler function
func check(e error) {
  if e != nil {
    panic(e)
  }
}

// Create ext name of .arp file name
// Must take in a .arp file name
func newName(ext string, name string) string {
  rune := []rune(name)
  return string(rune[:len(name)-3]) + ext
}

// Reads json data from zipped arp
func getJson(src string) map[string]interface{} {
  // Map of json
  var result map[string]interface{}
  // Open
  arch, err := zip.OpenReader(src)
  check(err)
  defer arch.Close()

  for _, f := range arch.File {
    json_file, err := f.Open()
    check(err)
    defer json_file.Close()

    // Read json object
    byte_value, _ := ioutil.ReadAll(json_file)

    // Make map of json
    json.Unmarshal([]byte(byte_value), &result)
  }

  return result
}

// Reads a .arp file
func ReadArp(filename string) map[string]interface{} {
  // Create name of zip of .arp
  zip_filename := newName("zip", filename)

  // Open .arp file
  arp_file, err := os.Open(filename)
  check(err)
  defer arp_file.Close()

  // Create new .zip file
  zip_file, err := os.Create(zip_filename)
  check(err)
  defer zip_file.Close()

  // Copy .arp file to .zip file
  copied, err := io.Copy(zip_file, arp_file)
  check(err)

  fmt.Println("Copied %d bytes.", copied)

  // Read json data
  output := getJson(zip_filename)

  os.Remove(zip_filename)

  return output
}
