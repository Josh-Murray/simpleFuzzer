package main

import (
	"fmt"
	"strings"
)
// Mutations
func insertRandomChar(s string) string{
	// something happens here that causes a panic
	if s == ""{
		return s
	}
	index := rng.Intn(len(s))
	char := rune(33 + rng.Intn(94))
	s = s[:index] + string(char) + s[index:]
	return s
}

func deleteRandom(s string) string{
	if s == ""{
		return s
	}
	letters := strings.Split(s, "")
	index := rng.Intn(len(letters))
	copy(letters[index:], letters[index+1:])
	return strings.Join(letters[:len(letters)-1],"")
}

func flipRandom(s string) string{
	if s == ""{
		return s
	}
	index:= rng.Intn(len(s))
	mask := 1 << uint(rng.Intn(7))
	char := rune(int(s[index]) ^ mask)
	return s[:index] + string(char) + s[index + 1:]
}

// Another useful mutation is to insert 'magic' values that naturally show up in bugs eg 0x0000
func mutateSeeds(conf Config,seeds []string) []string{
	if conf.Log{
		fmt.Println("[Log] --- Begin seed mutation ---")
	}
	inputs := make(chan string)
	outputs := make(chan string)
	go func(){
		for _, i := range seeds{
			inputs <- i
		}
		close(inputs)
	}()
	for i:=0; i< conf.NumWorkers; i++{
		go func(){
			mutateWorker(conf, inputs, outputs)
		}()
	}
	for i := range seeds{
		temp := <- outputs
		seeds[i] = temp
	}
	close(outputs)

	if conf.Log{
		fmt.Println("[Log] --- Finished seed mutation ---")
		fmt.Println("[Log] Mutated seeds: ", seeds)
	}
	return seeds

}

func mutateWorker(conf Config, inputs <- chan string, outputs chan <- string){
	mutations := []func(string)string{flipRandom, deleteRandom, insertRandomChar}
	var n string
	for i := range inputs{
		str := i
		for j:=0; j < conf.NumMutations; j++{
			function := mutations[rng.Intn(len(mutations))]
			n = function(str)
			if conf.Log{
				fmt.Printf("[Log] %s --> %s\n", i, n)
			}
		}

		outputs <- str
	}
}
