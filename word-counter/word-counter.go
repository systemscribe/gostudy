package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

func main() {
// Count number of words
  fmt.Println(count(os.Stdin))
}

func count (r io.Reader) int {
  //Scanner for reading text
  scanner := bufio.NewScanner(r)
  //How the scanner splits
  scanner.Split(bufio.ScanWords)
  
  //Define the counter
  wc := 0

  for scanner.Scan() {

    wc++
  }

  //return the counter
  return wc
}
