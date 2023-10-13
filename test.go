package main

import (
  "fmt"
  "net/http"
  "os"
  "strconv"
  "time"
)

func hello() float64 {
  t1 := time.Now().UnixNano()
  http.Get("http://127.0.0.1:8000/")
  diff := time.Now().UnixNano() - t1
  s_diff := float64(diff) / float64(time.Second)
  return s_diff
}

func main() {
  nums_str := os.Args[1]
  nums, err := strconv.Atoi(nums_str)
  if err != nil {
    panic(err)
  }

  diffs := []float64{}
  total := 0.0
  for i := 0; i < nums; i++ {
    diff := hello()
    diffs = append(diffs, diff)
    total += diff
  }

  for _, diff := range diffs {
    fmt.Printf("%.08f\n", diff)
  }

  fmt.Printf("total: %.08f, avg: %.08f", total, total/float64(nums))
}