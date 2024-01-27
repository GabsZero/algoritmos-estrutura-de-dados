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
	lutadores := []string{"ken", "chunli", "blanka", "sagat", "bison", "vega", "balrog", "dhalsin", "akuma", "ryu"}
	// percorrerLista(listaAleatoria)
	// calculaPrimeiroNumero(listaAleatoria)
	encontraRyu(lutadores)
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

func percorrerLista(array []int) {
	inicio := time.Now()
	for i := 0; i < len(array); i++ {
		if array[i] > 0 {
			fmt.Println("Número é maior que zero!")
		}
	}
	tempo := time.Since(inicio)
	log.Printf("Percorrer a lista levou %s", tempo)
}

func calculaPrimeiroNumero(array []int) {
	inicio := time.Now()
	if array[0] > 0 {
		fmt.Println("Número é maior que 0!")
	}
	tempo := time.Since(inicio)
	log.Printf("Calcular o primeiro número levou %s", tempo)
}
