package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Code struct {
	Code int
	MsgCH string
	MsgEN string
}
var codeMap = make(map[int]Code)

func InitCodeMap(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)

	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if strings.Index(line, "begincode") > 0 {
			continue
		}
		if strings.Index(line, "endcode") > 0 {
			break
		}
		//handler(line)
		linestrs := strings.Split(line,"|")
		//log.Println(linestrs)
		code := Code{}
		codeint,err := strconv.Atoi(linestrs[0])
		//log.Println(err)
		//log.Println(linestrs[0])
		//log.Println(codeint)


		code.Code = codeint
		code.MsgCH = linestrs[1]
		code.MsgEN = linestrs[2]
		codeMap[code.Code] = code

	}
	return nil
}

func Print(line string) {
	fmt.Println(line)
}

func main() {
	//dir,  _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//log.Println(dir)
	//
	//s, _ := exec.LookPath(os.Args[0])
	//log.Println(s)
	////InitCodeMap("C:/Users/soft/Desktop/code", Print)
	//
	//execpath, _ := os.Executable() // 获得程序路径
	//configfile := filepath.Join(filepath.Dir(execpath), "./config.yml")
	//log.Println(configfile)



	InitCodeMap("C:/Users/soft/Desktop/code.properties", Print)
	log.Println(codeMap)
	log.Println(len(codeMap))
}



