package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type CJSON struct {
	Cp string `json:"cp,omitempty"`
}

func main() {
	log.Println("端口：", 30888, " 等待复制")
	http.HandleFunc("/copy", CopyFunc)
	http.ListenAndServe(":30888", nil)

}

func CopyFunc(w http.ResponseWriter, r *http.Request) {
	sysType := runtime.GOOS
	log.Println(sysType)
	mod := &CJSON{}
	err := json.NewDecoder(r.Body).Decode(mod)
	if err != nil {
		log.Println(err.Error())
		return
	}

	buf := mod.Cp
	//log.Println("cp")

	dir, _ := ioutil.TempDir("", "IPHONE")
	now := time.Now().Unix()
	time := strconv.FormatInt(now, 10)
	fileName := dir + time + ".txt"

	f, err := os.Create(fileName)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer f.Close()
	_, err = f.Write([]byte(buf))
	if err != nil {
		log.Println(err.Error())
		return
	}
	if sysType == "windows" {
		log.Println(fileName)
		// windows系统
		cmd := exec.Command("cmd", "/c", "start", fileName)
		err = cmd.Start()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	fmt.Println(buf)

}
