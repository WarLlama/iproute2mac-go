package ifconfig

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	flagsMtuRe     = regexp.MustCompile(`^(\w+):\s+flags=(\d+)<[A-Z,]*>\smtu\s(\d+)`)
	loopbackLineRe = regexp.MustCompile(`^\tinet\s127.0.0.1`)
	etherLineRe    = regexp.MustCompile(`^\tether\s+([a-f:]+)`)
)

func ParseLink(lines []string) (*Link, error) {
	l := &Link{
		Operstate: "UP",
		Linkmode:  "DEFAULT",
		Group:     "default",
		Qdisc:     "noqueue",
	}
	flagMtuMatches := flagsMtuRe.FindStringSubmatch(lines[0])
	if len(flagMtuMatches) < 3 {
		return nil, errors.New("failed to process link data, missing flags or mtu")
	}
	l.Ifname = flagMtuMatches[1]
	flags, err := strconv.ParseInt(flagMtuMatches[2], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid flag value: %w", err)
	}
	l.Flags = ParseFlags(int(flags))
	mtu, err := strconv.Atoi(flagMtuMatches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid mtu value: %w", err)
	}
	l.Mtu = mtu
	for _, line := range lines[1:] {
		if match := loopbackLineRe.MatchString(line); match {
			l.LinkType = "loop"
			l.Address = "00:00:00:00:00:00"
			l.Broadcast = "00:00:00:00:00:00"
			l.Txqlen = 1000
		}
		if match := etherLineRe.FindStringSubmatch(line); len(match) > 1 {
			l.LinkType = "ether"
			l.Address = match[1]
			l.Broadcast = "ff:ff:ff:ff:ff:ff"
			l.Txqlen = 1000
		}
	}
	return l, nil
}

type Link struct {
	Ifindex   int      `json:"ifindex"`
	Ifname    string   `json:"ifname"`
	Flags     []string `json:"flags"`
	Mtu       int      `json:"mtu"`
	Qdisc     string   `json:"qdisc"`
	Operstate string   `json:"operstate"`
	Linkmode  string   `json:"linkmode"`
	Group     string   `json:"group"`
	Txqlen    int      `json:"txqlen"`
	LinkType  string   `json:"link_type"`
	Address   string   `json:"address"`
	Broadcast string   `json:"broadcast"`
}
