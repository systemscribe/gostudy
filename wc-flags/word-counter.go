package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
)

func main() {

// Define a flag -l to count lines, instead of words
  lines:=flag.Bool("l", false, "Count Lines")
  linesb:=flag.Bool("b", false, "Count Bytes")

// Parse user flags

  flag.Parse()
// Count number of words
  fmt.Println(count(os.Stdin, *lines, *linesb))
}

func count (r io.Reader, countLines bool, countBytes bool) int {
  //Scanner for reading text
  scanner := bufio.NewScanner(r)
  //How the scanner splits
  if !countLines {
  scanner.Split(bufio.ScanWords)
  }

  if !countBytes {
    scanner.Split(bufio.ScanBytes)
  }
  
  //Define the counter
  wc := 0

  for scanner.Scan() {

    wc++
  }

  //return the counter
  return wc
}
