package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var ls, _ = exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(ls))
	var pwd, _ = exec.Command("pwd").Output()
	fmt.Printf("-> pwd\n%s\n", string(pwd))
	var ifconfig, _ = exec.Command("ifconfig").Output()
	fmt.Printf("-> ifconfig\n%s\n ", string(ifconfig))
}
