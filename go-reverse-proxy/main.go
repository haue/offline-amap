package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var downloadPath = ""

func handleAll(w http.ResponseWriter, r *http.Request) {
	ruri := r.RequestURI
	var response = ""
	reFilename := regexp.MustCompile(`\w+\.[^\.]+$`)
	if strings.HasPrefix(ruri, "/maps") {
		a, _ := ioutil.ReadFile("js/1.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/v3/log/init") {
		callback, _ := r.URL.Query()["callback"]
		re := regexp.MustCompile(`^.*?\(`)
		a, _ := ioutil.ReadFile("js/2.js")
		response = re.ReplaceAllString(string(a), callback[0]+"&&"+callback[0]+"(")
	} else if strings.HasPrefix(ruri, "/ui/1.1/main.js") {
		a, _ := ioutil.ReadFile("js/3.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/ui/1.1/ui/misc/PathSimplifier.js") {
		a, _ := ioutil.ReadFile("js/4.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/ui/1.1/ui/misc/MarkerList.js") {
		a, _ := ioutil.ReadFile("js/5.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/count") {
		a, _ := ioutil.ReadFile("js/6.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/ui/1.1/plug/ext/jquery-1.12.4.min") {
		a, _ := ioutil.ReadFile("js/7.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/maps/plugin?v=2.0&cls=AMap.Scale") {
		a, _ := ioutil.ReadFile("js/8.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/ui/1.1/ui/geo/DistrictExplorer.js") {
		a, _ := ioutil.ReadFile("js/9.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/style?name=") {
		callback, _ := r.URL.Query()["callback"]
		a, _ := ioutil.ReadFile("js/10.js")
		re := regexp.MustCompile(`^.*?\(`)
		response = re.ReplaceAllString(string(a), callback[0]+"&&"+callback[0]+"(")
	} else if strings.HasPrefix(ruri, "/style/2.0") {
		a, _ := ioutil.ReadFile("js/11.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/ui/1.1/ui/geo/DistrictExplorer/assets/d_v2/an_410000.json?v=1.1.2") {
		a, _ := ioutil.ReadFile("js/12.js")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/static/commonWordsInfo.v2.1.json") {
		a, _ := ioutil.ReadFile("js/13.js")
		w.Write(a)
	} else if strings.HasPrefix(ruri, "/static/32-40863.v2.1.webp") {
		a, _ := ioutil.ReadFile("js/32-40863.v2.1.webp")
		w.Write(a)
	} else if strings.HasPrefix(ruri, "/nebula") {
		msg, _ := r.URL.Query()["msg"]
		h := sha1.New()
		h.Write([]byte(msg[0]))
		if strings.Contains(msg[0], "285fd") {
		}
		sha1 := hex.EncodeToString(h.Sum(nil))
		filepath := path.Join(downloadPath, sha1+".txt")
		if _, err := os.Stat(filepath); os.IsNotExist(err) || len(msg) < 1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		a, _ := ioutil.ReadFile(filepath)
		w.Header().Set("Content-Type", "application/x-protobuf")
		response = string(a)
	} else if strings.HasPrefix(ruri, "/sdf") {
		filename := reFilename.FindAllString(ruri, -1)[0]
		filepath := path.Join(downloadPath, filename)
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		a, _ := ioutil.ReadFile(path.Join(downloadPath, filename))
		w.Write(a)
	}
	fmt.Fprint(w, response)
}

func main() {
	port := getenvDefault("PORT", "8001")
	downloadPath = getenvDefault("DOWNLOADDIR", "download")
	http.HandleFunc("/", handleAll)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getenvDefault(k string, d string) string {
	r := ""
	a, _ := ioutil.ReadFile(".env")
	sa := string(a)
	arr := strings.Split(strings.Replace(sa, "\r\n", "\n", -1), "\n")
	b := false
	for _, s := range arr {
		kv := strings.Split(s, "=")
		if kv[0] == k {
			r = kv[1]
			b = true
			break
		}
	}
	if b {
		return r
	}
	return d
}
