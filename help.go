package main

import "fmt"

func showHelp() {
	fmt.Println()
	fmt.Println("-")
	fmt.Println("Usage: goc COMMAND [ARGS]")
	fmt.Println()
	fmt.Println("COMMAND")
	for _, build := range builds {
		fmt.Printf("     %-12s%s\n", build.cmd+":", build.help)
	}
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("     goc b")
	fmt.Println("       -go build in current directory")
	fmt.Println("     goc rl")
	fmt.Println("       -release linux package")
	fmt.Println()
	fmt.Println("   Ver: " + version + ", Author: NadOo")
	fmt.Println("   Source: https://github.com/nadoo/goc")
	fmt.Println("-")
}
