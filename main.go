package main

import (
	"fmt"
		
	"os/exec"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2{
		fmt.Println("Usage: scgen <bin_file>")
		return
	}

	bin := args[1]

	// Reference: https://gist.github.com/EdoardoVignati/64f2c864e70d8120e55d1b8253467c2b
	// Thanks :-)
	command := `objdump -d ` + bin + ` | grep '[0-9a-f]:' | grep -v 'file' | cut -f2 -d: | cut -f1-6 -d ' '| tr -s ' '| tr '\t' ' '| sed 's/ $//g' | sed 's/ /\\x/g' | paste -d '' -s | sed 's/^/"/' | sed 's/$/"/g'`	

	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

