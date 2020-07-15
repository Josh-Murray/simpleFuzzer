# simpleFuzzer
Implementation of a simple fuzzer that uses grammars to produce more syntacticly correct inputs.
Written in golang

## Motivations
- Get a high level understanding of how fuzzers work
- Learn a new language
- Get more comfortable with multi-threading programming

## TODO
- Implement a logging system
- Add more mutation functions
- Add ability to re-run cases with consistent outputs
	- Probably involves saving the rng seeds
	- Maybe define a file structure to pass as an arg? Ideally
	the fuzzer would re-run interesting cases automatically

- Implment an algorithm that finds the minimum substring that causes a crash
- replace string mutations with derviation tree mutations

