package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	listaAleatoria := rand.Perm(1000)

	percorrerLista(listaAleatoria)
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
