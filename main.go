package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/sheran/shernum/modules"
	"os/exec"

)

type Scmd struct {
	Cmd string
	Args []string
}

func (c *Scmd) Run() (string, error){
	cmd := exec.Command(c.Cmd, c.Args...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(),err
}

func NewCommand(cmd string, args []string) Scmd{
	return Scmd{Cmd: cmd, Args: args}
}


func main(){
	cmd := NewCommand("nmap",[]string{"-Pn","-n","-v0","-oX","/dev/stdout","-sC","-sV","-T4","--top-ports","10","192.168.69.119,123"})
	out, err := cmd.Run()
	if err != nil{
		panic(err)
	}
	//fmt.Println(out)
	var nmaprun modules.NmapRun
	err = xml.Unmarshal([]byte(out),&nmaprun)
	if err != nil{
		panic(err)
	}
	fmt.Println(nmaprun.GetHosts())

	//cmd := NewCommand("feroxbuster",[]string{"--stdin","-w","/Users/sheran/Documents/code/shernum/SecLists/" +
	//	"Discovery/Web-Content/raft-small-directories.txt","-t","5"})
	//cmd.Run("http://192.168.69.119\nhttp://192.168.69.123")

	//cmd := NewCommand("smbclient",[]string{"-L",fmt.Sprintf("\\\\%s\\","192.168.68.125"),"-U","\"\""})
	//cmd.Run()



}
