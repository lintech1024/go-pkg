package main

import (
	"crypto/elliptic"
	"encoding/json"
	"fmt"
)

func main() {
	curve := elliptic.P256()
	fmt.Println(Json(curve.Params()))
}

func Json(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
