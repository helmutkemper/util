package util

import "math"

// Arredonda um número de ponto flutuante para o inteiro mais próximo
// Round(valor, [em_torno], [dígitos_decimais]) float64
// float64 valor
// float64 em torno = 0.5
// float64 dígitos decimais

func Round( args ...interface{} ) float64 {
  var roundOn float64
  var places  float64

  value := args[0].(float64)

  if len( args ) == 2 {
    roundOn = args[1].(float64)
  } else {
    roundOn = 0.5
  }

  if len( args ) == 3 {
    places = args[2].(float64)
  } else {
    places = 0.0
  }

  var round float64
  pow := math.Pow(10, places)
  digit := pow * value
  _, div := math.Modf(digit)

  if div >= roundOn {
    round = math.Ceil(digit)
  } else {
    round = math.Floor(digit)
  }

  return round / pow
}
