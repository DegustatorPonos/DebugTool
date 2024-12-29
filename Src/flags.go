package main

import (
	"flag"
	"strconv"
	"strings"
)

var Port *int 

func ParseFlags() {
	Port = flag.Int("p", 6969, "A port that will be used. The default value will be used if passed incorrectly")
	flag.Parse()
	if *Port > 65535 || *Port <= 0{
		*Port = 6969
	}
}

func GetPort() string {
	var sb strings.Builder = strings.Builder{}
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(*Port))
	return sb.String()
} 
