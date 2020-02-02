package main

import(
	"fmt"
)
func main(){
	config := loadConfig()
	exp := generateExpression(config.Log, config.Length, config.Trials, config.Start, config.Grammar)
	fmt.Printf("expression = %s\n", exp)
}
