package main

import (
    "log"
    "os/exec"
    "strconv"
    "strings"

    "github.com/alexflint/go-arg"
    "github.com/fatih/color"
)

func main() {
    var vals []string

    var args struct {
        Port    int   `arg:"positional,required"`
        Simple  bool  `arg:"-s,--simple" help:"Show simplified output (just the protocol, local address, and PID/process"`
    }
    arg.MustParse(&args)

    portInfo, err := exec.Command("/usr/bin/netstat", "-tulpn").Output()
    if err != nil { log.Fatalln(err) }
    lines := strings.Split(string(portInfo), "\n")

    for _,v := range lines {
        if strings.Contains(v, ":" + strconv.Itoa(args.Port)) {
            split := strings.Split(v, " ")
            for _,v2 := range split {
                if strings.TrimSpace(v2) != "" {
                    vals = append(vals, v2)
                }
            }
        }
    }

    printVals(vals, args.Simple)
}


func printVals(vals []string, simple bool) {
    key := color.New(color.Bold, color.FgWhite).PrintfFunc()
    value := color.New(color.FgCyan).PrintlnFunc()
    keys := []string{"Protocol", "Recv-Q", "Send-Q", "Local Addr", "Remote Addr", "State", "PID/Process"}

    if simple {
        simpleVals := []int{0,3,6}
        for _,v := range simpleVals {
            key("%-14v", keys[v] + ": ")
            value(vals[v])
        }

    } else {
        for k,v := range keys {
            key("%-14v", v + ": ")
            value(vals[k])
        }
    }
}