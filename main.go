package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// This stupid little thing is to keep me from getting the yaml/base64 decoding k8s secrets all the time
// Instead of adding k8s things in here i opted for just adding the function
//
// decode () {
// 	 kubectl get secret $1 -o json | k8s-secret-decode
// }
//
// in my .zshrc

// Secret ...
type Secret struct {
	Data map[string]string `json:"data"`
}

func main() {
	secret := Secret{}
	err := json.NewDecoder(os.Stdin).Decode(&secret)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range secret.Data {
		result, err := base64.RawStdEncoding.DecodeString(v)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n%s\n", k, string(result))
	}
}
