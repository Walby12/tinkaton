package main

import(
  "fmt"
)

type Xor struct {
  or_w1 float64
  or_w2 float64
  or_b float64

  nand_w1 float64
  nand_w2 float64
  nand_b float64

  and_w1 float64
  and_w2 float64
  and_b float64
}

var xor_train = [...][3]int{
    {0, 0, 0},
    {1, 0, 1},
    {0, 1, 1},
    {1, 1, 0},
}

var train = xor_train

func forward(model Xor, x float64, y float64) float64 {
  a := sigmoidf(model.or_w1 * x + model.or_w2 * y + model.or_b)
  b := sigmoidf(model.nand_w1 * x + model.nand_w2 * y + model.nand_b)
  return sigmoidf(a * model.and_w1 + b * model.and_w2 + model.and_b)
}

func cost(m Xor) float64 {
  var result float64
  for i := range train {
    x1 := train[i][0]
    x2 := train[i][1]
    y := forward(m, float64(x1), float64(x2))
    d := y - float64(train[i][2])
    result += d*d
  }
  result /= float64(len(train))
  return result
}
