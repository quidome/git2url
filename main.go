package main

import (
	"fmt"
	giturls "github.com/whilp/git-urls"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Provide one argument, the helpers url.")
		usage()
		os.Exit(1)
	}
	url := args[0]

	fmt.Println(extractPathFromGitUrl(url))
}

func usage() {
	fmt.Printf("Usage:\n\t%s <helpers url>\n",
		filepath.Base(os.Args[0]))
}

func extractPathFromGitUrl(input string) string {
	url, err := giturls.Parse(input)
	if err != nil {
		fmt.Println("Could not parse url")
		os.Exit(2)
	}

	host := url.Host
	path := strings.TrimSuffix(url.Path, filepath.Ext(url.Path))

	return filepath.Join(host, path)
}
