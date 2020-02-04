package main

import(
	"fmt"
	"math/rand"
	"time"
)
var rng *rand.Rand
func main(){
	config := loadConfig()
	rng = rand.New(rand.NewSource(time.Now().Unix()))
	seeds := generateSeeds(config)
	seeds = mutateSeeds(config, seeds)
	fmt.Println(seeds)
}
