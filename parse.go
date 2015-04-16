package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type SLA struct {
	SPIO     string
	Uses     string
	Status   string
	Type     string
	x        string
	y        string
	dunno    string
	LandArea string
	GFA      string
	UUID     string
	URL      string
	Price    float64
	Add1     string
	Add2     string
	Add3     string
	Add4     string
	Add5     string
}

func main() {

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var data []string
	err = json.Unmarshal(content, &data)
	if err != nil {
		panic(err)
	}

	var props []SLA

	for index := 0; index < len(data); index += 18 {
		var p SLA
		p.SPIO = data[index]
		p.Uses = data[index+1]
		p.Status = data[index+2]
		p.Type = data[index+3]
		p.x = data[index+4]
		p.y = data[index+5]
		p.dunno = data[index+6]
		p.LandArea = data[index+7]
		p.GFA = data[index+8]
		p.UUID = data[index+9]
		p.URL = strings.Replace(data[index+10], "../../", "http://www.landapplications.gov.sg/SPIOWeb/", 1)
		p.Price, err = strconv.ParseFloat(strings.Replace(data[index+11], ",", "", 1), 64)
		if err != nil {
			panic(err)
		}
		// p.Price, err = strconv.Atoi(data[index+11])
		p.Add1 = data[index+12]
		p.Add2 = data[index+13]
		p.Add3 = data[index+14]
		p.Add4 = data[index+15]
		props = append(props, p)
	}

	t, err := template.New("foo").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
</head>
<body>
<ol>
{{ range . }}{{if and (lt .Price 3000.0) (eq .Status "AN") }}<li><a href="{{ .URL }}">{{ .GFA }} {{ .Price }}</li>
{{ end }}{{ end }}</ol></body>
</html>`)
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err = t.Execute(w, props)
	if err != nil {
		panic(err)
	}
	w.Flush()

	h := sha1.New()
	h.Write(b.Bytes())
	fmt.Println(b.Bytes())
	fmt.Printf("% x\n", h.Sum(nil))

	sha1_hash := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("Writing", sha1_hash)
	ioutil.WriteFile(sha1_hash+".txt", b.Bytes(), 0644)

}
