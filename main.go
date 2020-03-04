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
	seeds := generateSeeds(config)
	seeds = mutateSeeds(config, seeds)
	fmt.Println(seeds)
	fuzzInputs(config, seeds)
	fmt.Println("done")
}
