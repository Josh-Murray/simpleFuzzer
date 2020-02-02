package main

import(
	"fmt"
	"regexp"
	"math/rand"
	"strings"
	"time"
)
// re to find the nonTerminal elements
var re = regexp.MustCompile(`(<[^<> ]*>)`)

// return splice of nonTerminal elements in subSet
func nonTerminal(subSet string) []string{
	return re.FindAllString(subSet, -1)
}

func generateExpression(log bool, maxLength int, maxTrials int, start string, grammar map[string][]string) string {
	rand.Seed(time.Now().Unix())
	trial := 0
	term := start
	for len(nonTerminal(term)) > 0{
		// Replace a random expression from term with a random expansion from grammar
		choice := nonTerminal(term)[rand.Intn(len(nonTerminal(term)))]
		expansions := grammar[choice]
		// Could be improved by enforcing expansions choice to have < maxLength nonTerminals, instead of current trial and error approach
		expansionChoice := expansions[rand.Intn(len(expansions))]
		temp := strings.Replace(term, choice, expansionChoice,1)
		// Check if temp follows constraints
		if len(nonTerminal(temp)) < maxLength{
			term = temp
			trial = 0
			if log{
				fmt.Printf("[Log] expression = %s (%s --> %s) \n", term, choice, expansionChoice)
			}
		}else{
			trial ++
			if trial > maxTrials{
				fmt.Println("[ERROR] Reached max trial length")
				break
			}
		}

	}
	return term
}
