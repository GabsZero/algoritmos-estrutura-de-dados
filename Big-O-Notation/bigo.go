package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// listaAleatoria := rand.Perm(100000)
	// lutadores := []string{"ken", "chunli", "blanka", "sagat", "bison", "vega", "balrog", "dhalsin", "akuma", "ryu"}
	// percorrerLista(listaAleatoria)
	// calculaPrimeiroNumero(listaAleatoria)
	// encontraRyu(lutadores)
	logParesDoArray([]int{1, 2, 3, 4, 5})
}

func encontraRyu(lutadores []string) {
	for _, lutador := range lutadores {
		fmt.Println("Tentando encontrar o ryu")
		if lutador == "ryu" {
			fmt.Println("Encontrei o ryu!")
			break
		}
	}
}

func logParesDoArray(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Printf("%d %d\n", array[i], array[j])
		}
	}
}

func percorrerLista(array []int) {
	inicio := time.Now()              // O(1) por ser apenas uma atribuição
	for i := 0; i < len(array); i++ { // O(n) porque vai acontecer n vezes a partir de uma chamada da função
		if array[i] > 0 { // O(n)
			fmt.Println("Número é maior que zero!") // O(n)
		}
	}
	tempo := time.Since(inicio)                     // O(1)
	log.Printf("Percorrer a lista levou %s", tempo) // O(1)
}

func calculaPrimeiroNumero(array []int) {
	inicio := time.Now()
	if array[0] > 0 {
		fmt.Println("Número é maior que 0!")
	}
	tempo := time.Since(inicio)
	log.Printf("Calcular o primeiro número levou %s", tempo)
}
