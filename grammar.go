package main

import(
	"fmt"
	"regexp"
	"strings"
)
// re to find the nonTerminal elements
var re = regexp.MustCompile(`(<[^<> ]*>)`)

// return splice of nonTerminal elements in subSet
func nonTerminal(subSet string) []string{
	return re.FindAllString(subSet, -1)
}

func generateExpression(Start string, Grammar map[string][]string, Length int, Trials int) string {
	trial := 0
	term := Start
	for len(nonTerminal(term)) > 0{
		// Replace a random expression from term with a random expansion from grammar
		choice := nonTerminal(term)[rng.Intn(len(nonTerminal(term)))]
		expansions := Grammar[choice]
		// Could be improved by enforcing expansions choice to have < maxLength nonTerminals, instead of current trial and error approach
		expansionChoice := expansions[rng.Intn(len(expansions))]
		temp := strings.Replace(term, choice, expansionChoice,1)
		// Check if temp follows constraints
		if len(nonTerminal(temp)) <Length{
			term = temp
			trial = 0
		}else{
			trial ++
			if trial > Trials{
				fmt.Println("[ERROR] Reached max trial length")
				break
			}
		}

	}
	return term
}

func generateSeeds(conf Config) []string{
	if conf.Log{
		fmt.Println("[Log] --- Generating Seeds ---")
	}
	var seeds [] string
	var newSeed string
	for i:=0; i < conf.NumSeeds; i++{
		newSeed = generateExpression(conf.Start, conf.Grammar, conf.Length, conf.Trials)
		if conf.Log{
			fmt.Println("[Log] Generated seed: " + newSeed)
		}
		seeds = append(seeds, newSeed)
	}
	if conf.Log{
		fmt.Println("[Log] --- Finished generating seeds ---")
		fmt.Println("[Log] Seeds: " , seeds)
	}
	return seeds
}
