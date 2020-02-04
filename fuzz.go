package main

import (
	"fmt"
	"strings"
)
// Mutations
func insertRandomChar(s string) string{
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

// Another useful mutation is to insert 'magic' values that naturally show up in bugs
func mutateSeeds(conf Config,seeds []string) []string{
	if conf.Log{
		fmt.Println("[Log] --- Begin seed mutation ---")
	}
	mutations := []func(string)string{flipRandom, deleteRandom, insertRandomChar}
	var n string
	for i := range seeds{
		for j:=0; j < conf.NumMutations; j++{
			function := mutations[rng.Intn(len(mutations))]
			n = function(seeds[i])
			//flipRandom(seeds[i])
			if conf.Log{
				fmt.Printf("[Log] %s --> %s\n", seeds[i], n)
			}
			seeds[i] = n
		}
	}
	if conf.Log{
		fmt.Println("[Log] --- Finished seed mutation ---")
		fmt.Println("[Log] Mutated seeds: ", seeds)
	}
	return seeds

}
