package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
)

var repoPortInfo = make(map[string]models.Port)

// InsertPortData will read the data from the json file and insert the data in repository
func InsertPortData(w http.ResponseWriter, r *http.Request) {
	//read the port data from the post.json file
	pData, err := ioutil.ReadFile("ports.json")
	if err != nil {
		log.Fatal(err)
	}
	var portsInfo = make(map[string]models.Port)
	err2 := json.Unmarshal(pData, &portsInfo)
	if err2 != nil {
		log.Fatal(err2)
	}
	for k, v := range portsInfo {
		repoPortInfo[k] = v
	}
	fmt.Fprintf(w, "ok")

}

// GetPortData gets the port data from repository
func GetPortData(w http.ResponseWriter, r *http.Request) {
	pData := new(bytes.Buffer)
	for k, v := range repoPortInfo {
		fmt.Fprintf(pData, "%s  :  \"%s\"\n", k, v)
	}
	fmt.Fprintf(w, pData.String())
}
