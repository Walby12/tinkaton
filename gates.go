package main

import(
  "fmt"
  "math/rand/v2"
  "math"
)

var or_train = [...][3]int{
    {0, 0, 0},
    {1, 0, 1},
    {0, 1, 1},
    {1, 1, 1},
}

var and_train = [...][3]int{
    {0, 0, 0},
    {1, 0, 0},
    {0, 1, 0},
    {1, 1, 1},
}

var nand_train = [...][3]int{
    {0, 0, 1},
    {1, 0, 1},
    {0, 1, 1},
    {1, 1, 0},
}

var train = nand_train

func sigmoidf(x float64) float64 {
  return 1 / (1 + math.Exp(-x))
}

func cost(w1 float64, w2 float64, bias float64) float64 {
  var result float64
  for i := range train {
    x1 := train[i][0]
    x2 := train[i][1]
    y := sigmoidf(float64(x1) * w1 + float64(x2) * w2 + bias)
    d := y - float64(train[i][2])
    result += d*d
  }
  result /= float64(len(train))
  return result
}

func main() {
  eps := 1e-3
  learn_rate := 1e-1
  w1 := rand.Float64()
  w2 := rand.Float64()
  bias := rand.Float64()
  
  for range 10000*1000 {
    c := cost(w1, w2, bias)
    dw1 := (cost(w1 + eps, w2, bias) -c) / eps
    dw2 := (cost(w1 , eps + w2, bias) -c) / eps
    db := (cost(w1, w2, bias + eps) - c) / eps
    w1 -= learn_rate * dw1
    w2 -= learn_rate * dw2
    bias -= learn_rate * db 
  }
  fmt.Println("w1:", w1, "w2:", w2, "bias:", bias, "c:", cost(w1, w2, bias))

  for i := range 2 {
    for j := range 2 {
      fmt.Println(i,"|", j, sigmoidf(float64(i) * w1 + float64(j) * w2 + bias))

    }
  }
}
