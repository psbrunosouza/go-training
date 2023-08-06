package mapas

import (
	"math"
	"strings"
)

// OBJETIVO
// criar um map
// adicionar nova chave ao map
// modificar uma chave pre existente
// testar existencia de um map
func QuantityOfWords(phrase string) (result map[string]int) {
	result = make(map[string]int)
	source := strings.Split(phrase, " ")

	for _, v := range source {
		if _, ok := result[v]; ok {
			result[v] += 1
		}else {
			result[v] = 1
		}
	}

	return
}

// Mergear maps
func JoinMaps(mapa map[string]int, mapa2 map[string]int) (result map[string]int) {
	result = make(map[string]int)

	for k, v := range mapa {
		result[k] = v
	}

	for k, v := range mapa2 {
		result[k] = v
	}

	return result
}

// removendo uma chave de um map
func DeleteKeyFromMap(mapa map[string]int, key string) (result map[string]int) {
	result = mapa;
	delete(mapa, key) 
	return
}

// fazendo uma media dos valores de um map
func Avg(mapa map[string]float64) (result float64) {
	amount := 0.0
	quantity := 0.0
	
	for _, value := range mapa {
		amount += value
		quantity += 1
	}

	result = amount / quantity

	return
}

// pegando o valor maximo de um map
func FindTheWinner(mapa map[string]float64) (result float64) {
	for _, value := range mapa {
		result = math.Max(value, result)
	}

	return result
}

// Testa se um map esta vazio
func IsMapFilled(x map[string]int) bool {
	return len(x) > 0
}