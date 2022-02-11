package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	pathTx = "/tx"
)

type TxResponse struct {
	Hash []byte `json:"hash"`
	Height string `json:"height"`
	Index int `json:"index"`
	Tx []byte `json:"tx"`
}

func getTx(nodeAddress string, hash string) (interface{}, error) {
	// prepare the request
	req, err := http.NewRequest("GET", nodeAddress, nil)
	if err != nil {
		return nil, err
	}
	values := req.URL.Query()
	values.Add("hash", fmt.Sprintf("%v", hash))
	req.URL.RawQuery = values.Encode()
	req.URL.Path = path.Join(req.URL.Path, pathBlock)

	// build client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// extract the response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	txResp := TxResponse{}
	err = json.Unmarshal(body, &txResp)
	if err != nil {
		return nil, err
	}

	var out []interface{}
	tx, err := unpack(txResp.Tx)
	out = append(out, tx)

	return out, nil
}
