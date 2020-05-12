package main

import(
	"fmt"
	"math/rand"
	"time"
)
var rng *rand.Rand
func main(){
	rng = rand.New(rand.NewSource(time.Now().Unix()))
	config := loadConfig()
	fmt.Println("Generating seeds...")
	seeds := generateSeeds(config)
	fmt.Println("done")
	fmt.Println("Mutating seeds...")
	seeds = mutateSeeds(config, seeds)
	fmt.Println("done")
	fmt.Println("Running fuzz cases...")
	fuzzInputs(config, seeds)
	fmt.Println("done")
}
