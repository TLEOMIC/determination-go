package tool

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

var env map[string]string

func init(){
	env = make(map[string]string)
	makeEnv()
}

func Env(key string,value string) string{
	if env[key]	== ""{
		return value
	}
	return env[key]	
}

func makeEnv(){
	f, err := os.Open("./.env")
	if err != nil {
	  // return nil, err
	  fmt.Println(err)
	  return
	}
	
	br := bufio.NewReader(f)
	var arr []string
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
        arr = strings.Split(string(a), "=")
        //避免异常
        if len(arr) >= 2 {
        	env[arr[0]] = arr[1]
        }
    }
}