package slices

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// dividir um range de um slice e retornar como um novo slice
// acaba sendo simples, apenas usando a sintaxe slice[1:5] e dizer as posições
// sendo a primeira o inicio e a ultima o final, caso vazio
// indica que é desde o começo ou desde o final dependendo da posição
func ExecOne() {
	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subslice := slice[2:5]

	for _, value := range subslice {
		println(value)
	}
}

// inverter uma string
// usamos a seguinte logica, trocamos as posições
// das primeiras letras com as ultimas
// dessa maneira a medida que iteramos até o meio do slice
// todas as palavras resultarão invertidas
func InverteOne(value string) string {
	valueInBytes := []rune(value)
	length := len(valueInBytes)

	// for i, j := 0, len(valueInBytes)-1; i < j; i, j = i+1, j-1 {
	// 	valueInBytes[i], valueInBytes[j] = valueInBytes[j], valueInBytes[i]
	// }

	for i := 0; i < length/2; i++ {
		valueInBytes[i], valueInBytes[length-i-1] = valueInBytes[length-i-1], valueInBytes[i]
	}

	return string(valueInBytes)
}

// esse exercício tem objetivo de ajudar a memorizar a criação de um slice
// usamos a seguinte sintaxe []int{}, isso implica que podemos ter slices
// com apenas um tipo de valor e podemos ter slices de outros tipos, como:
// string, ou um tipo criado
func CreateSlices() {
	sliceExample := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} 
	fmt.Println(sliceExample[2:5])
}

// 
func AddElementsToSlice() {
	sliceExample := []string{"primeiro", "segundo"}

	newSlice := append(sliceExample, sliceExample[1:]...)

	fmt.Println(newSlice)
}

func CopyElements() {
	// slice original
	source := []int{1, 2, 3, 4, 5}

	// criar um slice com o mesmo tamanho do anterior
	destination := make([]int, len(source))

	// copiando conteudo da fonte para o destino
	copy(destination, source)

	fmt.Println(destination)
}

/*
	Remove o elemento através do seu index
*/
func RemoveElementByIndex(slice []int, position int) []int {
	return append(slice[:position], slice[position+1:]...)
}

func ElementExists(slice []int, v int) bool {
	for _, value := range slice {
		if(value == v){
			return true
		}
	}
	return false
}

func FindElement(slice []int, v int) (elem int, error error) {
	for _, value := range slice {
		if(value == v){
			elem = v
			return elem, nil
		}
	}
	
	return 0, error
}

func FindIndexElement(slice []int, v int) (i int, err error ) {
	for index, value := range slice {
		if(value == v){
			i = index
			return i, nil
		}
	}
	return 0, err
}

/*
	valor maximo dentro de um slice
*/
func MaxValue(slice []int) (max int) {
	for _, value := range slice {
		if(max < value){
			max = value
		}
	}
	return max
}

/*
	valor minimo dentro de um slice
*/
func MinValue(slice []int) (min int) {
	min = slice[0]

	for _, value := range slice {
		if(min > value){
			min = value
		}
	}
	return min
}

/*
	valor minimo e maximo dentro de um slice
*/
func MinMaxValue(slice []int) (min int, max int) {
	min = slice[0]
	max = slice[0]

	for _, value := range slice {
		if(min > value){
			min = value
		}
		
		if(max < value){
			max = value
		}
	}
	return min, max
}

func FakeBin(x string) (result string) {
	for i := range x {
		eachValue, _ := strconv.Atoi(fmt.Sprintf("%c", x[i]))

		if(eachValue < 5) {
			result += "0"
		}else if(eachValue >= 5) {
			result += "1"
		}
  }
  
	return
}

func ReverseText(word string) string {
	wordConverted := []rune(word)

	for i := 0; i < len(word)/2; i++ {
		wordConverted[i], wordConverted[len(word)-i-1] = wordConverted[len(word)-i-1], wordConverted[i]
	}

	return string(wordConverted)
}


func Combat(health, damage float64) float64 {
  health -= damage
  
  if(health < 0){
    return 0
  }
  
  return health
}

func CombatAnotherSolution(health, damage float64) float64 {
	return math.Max(health - damage, 0.0)
}

func ReverseSeq(n int) (slice []int) {
	for i := n; i > 0; i-- {
		fmt.Println(i)
		slice = append(slice, i)
	}
  return
}
 

func ToAlternatingCase(str string) (result string) {

	for _, v := range str {
		if(unicode.IsUpper(v)){
			result += strings.ToLower(fmt.Sprintf("%c", v))
		}else {
			result += strings.ToUpper(fmt.Sprintf("%c", v))
		}
	
	}
  return
}

func Maps(x []int) (result []int) {
  result = make([]int, len(x))
  
  for i := range x {
    result = append(result, i * 2)
  }
  return result
}