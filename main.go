package main

import (
	"fmt"
	"os"

	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vals"
	"src.elv.sh/pkg/eval/vars"
	"src.elv.sh/pkg/parse"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <script> [args...]\n", os.Args[0])
		os.Exit(1)
	}

	scriptPath := os.Args[1]
	scriptArgs := os.Args[2:]

	content, err := os.ReadFile(scriptPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading script: %v\n", err)
		os.Exit(1)
	}

	ev := eval.NewEvaler()

	// Set $args variable in the global namespace
	anyArgs := make([]any, len(scriptArgs))
	for i, arg := range scriptArgs {
		anyArgs[i] = arg
	}
	argsList := vals.MakeList(anyArgs...)

	ev.ExtendGlobal(eval.BuildNs().AddVar("args", vars.FromInit(argsList)).Ns())

	src := parse.Source{Name: scriptPath, Code: string(content)}

	ports, cleanup := eval.PortsFromStdFiles("")
	defer cleanup()

	err = ev.Eval(src, eval.EvalCfg{Ports: ports})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
