package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func GetPecahan(in string) (result string, err error) {
	pecahans := map[string]int{}

	type PecahanCount struct {
		Pecahan int
		Total   int
		Label   string
	}
	in = strings.ReplaceAll(in, ".", "")
	target, err := strconv.Atoi(in)

	if err != nil {
		return
	}

	if target%100 != 0 {
		target += 1
	}

	// create pecahans list and label on rupiah format for fast labeling
	counter := []PecahanCount{
		{100000, 0, "100.000"},
		{50000, 0, "50.000"},
		{20000, 0, "20.000"},
		{10000, 0, "10.000"},
		{5000, 0, "5.000"},
		{2000, 0, "2.000"},
		{1000, 0, "1.000"},
		{500, 0, "500"},
		{200, 0, "200"},
		{100, 0, "100"},
	}

	/*
		145000 / 100000 = 1
		145000 % 100000 = 45000 <- kurang
		---
		45000 / 50000 = 0
		45000 / 50000 = 45000 <- kurang
		---
		45000 / 20000 = 2
		45000 % 20000 = 5000 <-kurang
		---
		10000 ->skip
		5000 -> ok
	*/

	for i := 0; i < len(counter); i++ {
		if target > 0 {
			// get pecahan possibility round floor or even
			counter[i].Total = target / counter[i].Pecahan
			// get remander, remainder will be iterate to get pecahan possibility again
			target = target % counter[i].Pecahan
		}
	}

	for _, v := range counter {
		if v.Total > 0 {
			pecahans[fmt.Sprintf("Rp. %s", v.Label)] = v.Total
		}
	}

	pecahansJson, err := json.Marshal(pecahans)
	if err != nil {
		return
	}

	result = string(pecahansJson)

	return

}

func TestSolvePecahan(t *testing.T) {

	testCases := []struct {
		Input    string
		Expected string
	}{
		{"145.000", `{"Rp. 100.000":1,"Rp. 20.000":2,"Rp. 5.000":1}`},
		{"11.500", `{"Rp. 1.000":1,"Rp. 10.000":1,"Rp. 500":1}`},
		{"5000", `{"Rp. 5.000":1}`},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			res, err := GetPecahan(tc.Input)
			if err != nil {
				t.Error(err)
			}

			if res != tc.Expected {
				t.Error("Expected :", tc.Expected, "Actual :", string(res))
			}
		})
	}

}
