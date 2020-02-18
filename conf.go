package main
import (
	"encoding/json"
	"os"
	"fmt"
)
type Config struct{
	Log bool
	Length int
	Trials int
	Start string
	Grammar map[string] []string
	NumMutations int
	NumSeeds int
	NumWorkers int
	Flags string

}
func loadDefaults() Config{
	 grammar := map[string][]string{
		"<start>":	{"<expr>"},
		"<expr>":	{"<term> + <expr>", "<term> - <expr>", "<term>"},
		"<term>":	{"<factor> * <term>", "<factor> / <term>", "<factor>"},
		"<factor>":	{"+<factor>","-<factor>","(<expr>)","<integer>.<integer>","<integer>"},
		"<integer>":	{"<digit><integer>", "<digit>"},
		"<digit>":	{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}
	return Config{true, 10, 100, "<start>", grammar, 3, 5, 5}
}
// TODO make this better
func checkConf(conf Config) (bool, string){
	// unclear why there isn't a trivial way to ensure missing fields throw an error but here we are
	if conf.Trials == 0 {
		return false, "Trials not found or 0"
	}
	if conf.Length == 0{
		return false, "Length not found or 0"
	}
	if conf.NumMutations == 0{
		return false, "NumMutations not found or 0"
	}
	if conf.NumSeeds == 0{
		return false, "NumMumSeeds not found or 0"
	}
	if conf.NumWorkers == 0{
		return false, "NumWorkers not found or 0"
	}
	if _, ok := conf.Grammar[conf.Start]; !ok{
		return false, "Start not in grammar"
	}
	// TODO: implement something that checks a grammar is valid
	return true, ""
}
func loadConfig() Config{
	file, err := os.Open("conf.json")
	if err != nil {
		fmt.Println("[Warn] Could not find config file, loading defaults")
		loadDefaults()
		return loadDefaults()
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	config := Config{}
	if err = decoder.Decode(&config); err != nil{
		fmt.Printf("[Warn] Error loading config.json (%s), loading defaults\n", err.Error())
		return loadDefaults()
	}
	if ok, msg := checkConf(config); !ok{
		fmt.Printf("[Warn] Error loading config.json (%s), loading defaults\n",msg)
		return loadDefaults()
	}
	if config.Log{
		fmt.Println("[Log] Successfully loaded conf.json")
	}
	return config
}
