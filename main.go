package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/WarLlama/iproute2mac-go/ifconfig"
)

const (
	argLink      = "link"
	ifconfigExec = "/sbin/ifconfig"
)

func main() {
	var isJson bool
	jsonFlag := flag.Bool("j", false, "Output results in JavaScript Object Notation (JSON).")
	flag.Parse()
	args := flag.Args()
	if jsonFlag != nil {
		isJson = *jsonFlag
	}
	for _, a := range args {
		switch a {
		default:
			fmt.Printf("Object %q is unknown, try \"ip help\".\n", a)
			os.Exit(1)
		case argLink:
			if err := handleLink(isJson, args); err != nil {
				fmt.Printf("Failed to parse link data: %v", err)
				os.Exit(1)
			}
			return
		}
	}
}

func handleLink(isJson bool, args []string) error {
	cmd := exec.Command(ifconfigExec, "-a")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if !isJson {
		fmt.Print(out.String())
		return nil
	}
	return formatIfConfigJson(out.String())
}

func formatIfConfigJson(data string) error {
	lines := strings.Split(data, "\n")
	var links []*ifconfig.Link
	var linkLines []string
	for _, l := range lines {
		match, err := regexp.MatchString(`^\t`, l)
		if err != nil {
			fmt.Printf("failed to parse line from ifconfig: %s\n", l)
			return err
		}
		if match {
			linkLines = append(linkLines, l)
			continue
		}
		if len(linkLines) == 0 {
			linkLines = []string{l}
			continue
		}
		link, err := ifconfig.ParseLink(linkLines)
		if err != nil {
			return err
		}
		link.Ifindex = len(links) + 1
		links = append(links, link)
		linkLines = []string{l}
	}
	jsonLinks, err := json.Marshal(links)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", jsonLinks)
	return nil
}
