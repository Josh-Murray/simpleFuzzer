package main
import (
	"fmt"
	"sync"
	"encoding/hex"
	"os/exec"
)

func fuzzInputs(conf Config, seeds []string){
	inputs := make(chan string)
	//ouputs := make(chan bool)
	var wg sync.WaitGroup
	// fill worker chan
	go func(){
		for _, i := range seeds{
			inputs <- i
		}
		close(inputs)
	}()
	wg.Add(conf.NumWorkers)
	for i:=0; i < conf.NumWorkers; i++{
		go func(){
			defer wg.Done()
			runWorker(conf.File, inputs, conf.Log)
		}()
	}
	wg.Wait()
}
// unclear how i want to handle flags and inputs
func runWorker(file string, inputs <- chan string, log bool){
	for i := range inputs{
		if log{
			// print hex since some ascii chars will break the terminal
			fmt.Println("[Log] Running with 0x"+ hex.EncodeToString([]byte(i)))
		}
		out, err := exec.Command(file, i).Output()
		if err != nil{
			// Save this to output file
			fmt.Printf("Input 0x%s (hex) caused error %s\n", hex.EncodeToString([]byte(i)), err.Error())
			fmt.Printf("Output was %s\n", out)
		}
	}
}

