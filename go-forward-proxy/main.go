package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/elazarl/goproxy"
)

var downloadDir = "download"

func main() {
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		os.Mkdir(downloadDir, os.ModePerm)
	}
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnResponse(goproxy.ReqHostMatches(regexp.MustCompile(".amap.com"))).DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		msg := ctx.Req.URL.Query()["msg"]
		ruri := ctx.Req.URL.RequestURI()
		if len(msg) > 0 {
			h := sha1.New()
			h.Write([]byte(msg[0]))
			sha1 := hex.EncodeToString(h.Sum(nil))
			outFile, _ := os.Create(path.Join(downloadDir, sha1+".txt"))
			defer outFile.Close()
			bodyBytes, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			outFile.Write(bodyBytes)
			return r
		} else if strings.HasPrefix(ctx.Req.URL.Host, "glyph") {
			if strings.Contains(ruri, "sdf") {
				re := regexp.MustCompile(`\w+\.[^\.]+$`)
				ma := re.FindAllString(ruri, -1)
				if len(ma) > 0 {
					filename := ma[0]
					filepath := path.Join(downloadDir, filename)
					if _, err := os.Stat(filepath); os.IsNotExist(err) {
						outFile, _ := os.Create(filepath)
						defer outFile.Close()
						bodyBytes, _ := ioutil.ReadAll(r.Body)
						r.Body.Close()
						r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
						outFile.Write(bodyBytes)
						return r
					}
				}
			}
		}
		return r
	})
	port := getenvDefault("PORT", "8080")
	downloadDir = getenvDefault("DOWNLOADDIR", "download")
	log.Fatal(http.ListenAndServe(":"+port, proxy))
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
