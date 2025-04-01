package main

import(
  "fmt"
  "math/rand/v2"
)

var train = [...][3]int{
    {0, 0, 0},
    {1, 0, 1},
    {0, 1, 1},
    {1, 1, 1},
}

func cost(w1 float64, w2 float64) float64 {
  var result float64
  for i := range train {
    x1 := train[i][0]
    x2 := train[i][1]
    y := float64(x1) * w1 + float64(x2) * w2
    d := y - float64(train[i][1])
    result += d*d
  }
  result /= float64(len(train))
  return result
}

func main() {
  eps := 1e-3
  learn_rate := 1e-2
  w1 := rand.Float64()
  w2 := rand.Float64()
  
  for range 4000 {
    c := cost(w1, w2)
    fmt.Println("w1:", w1, "w2:", w2, "c:", c)
    dw1 := (cost(w1 + eps, w2) -c) / eps
    dw2 := (cost(w1 , eps + w2) -c) / eps
    w1 -= learn_rate * dw1
    w2 -= learn_rate * dw2
  }
}
