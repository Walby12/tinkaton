package main

import (
	"math/rand/v2"
	"fmt"
)

var train = [...][2]int{
    {0, 0},
    {1, 2},
    {2, 4},
    {3, 6},
    {4, 8},
}

func cost(w float64, bias float64) float64 {
  var result float64
  for i := range train {
    x := train[i][0]
    y := float64(x) * w + bias
    d := y - float64(train[i][1])
    result += d*d
  }
  result /= float64(len(train))
  return result
}

func main() {
  w := rand.Float64()*10
  bias := rand.Float64()*5
  eps := 1e-3
  learn_rate := 1e-3
  for range 500 {
    distance_cost := (cost(w + eps, bias) - cost(w, bias)) / eps
    distance_bias :=  (cost(w, eps + bias) - cost(w, bias)) / eps
    w -= learn_rate * distance_cost
    bias -= learn_rate * distance_bias
    fmt.Println("Cost:",cost(w, bias), "value:",w, "bias",)
  }
  fmt.Println("--------------")
  fmt.Println("w:",w, "b:", bias)
}
