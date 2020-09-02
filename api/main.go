package main

import (
	// "bytes"
	//  "compress/gzip"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	"errors"
	_ "fmt"
	_ "net/http"
	"strconv"
	"strings"

	types "code.vegaprotocol.io/chain-explorer-api/proto"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type UnsignedTx struct {
	Type    string
	Command string
}

type SignedTx struct {
	Type    string
	Command string
	Sig     []byte
	PubKey  string
	Nonce   uint64
}

func unpackSignedTx(rawtx []byte) (interface{}, error) {
	txBundle := types.SignedBundle{}
	err := proto.Unmarshal(rawtx, &txBundle)
	if err != nil {
		return nil, err
	}

	tx := types.Transaction{}
	err = proto.Unmarshal(txBundle.Tx, &tx)
	if err != nil {
		return nil, err
	}

	protoCmd, cmd, err := txDecode(tx.InputData)
	if err != nil {
		return nil, err
	}

	cmdTx, err := unmarshalCommand(protoCmd, cmd)
	if err != nil {
		return nil, err
	}

	hexPubKey := hex.EncodeToString(tx.GetPubKey())

	m := jsonpb.Marshaler{}
	marshalTx, err := m.MarshalToString(cmdTx)
	if err != nil {
		return nil, err
	}

	return &SignedTx{
		Type:    cmd.String(),
		Command: marshalTx,
		Sig:     txBundle.Sig.Sig,
		PubKey:  "0x" + hexPubKey,
		Nonce:   tx.Nonce,
	}, nil
}

func txDecode(input []byte) ([]byte, Command, error) {
	if len(input) > 37 {
		return input[37:], Command(input[36]), nil
	}
	return nil, 0, errors.New("payload size is incorrect, should be > 38 bytes")
}

func unpack(tx []byte) (interface{}, error) {
	return unpackSignedTx(tx)
}

type request struct {
	NodeURL     string  `json:"node_url"`
	BlockHeight *uint64 `json:"block_height"`
	TxHash      string  `json:"tx_hash"`
}

func handler(ev events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST",
		"Access-Control-Allow-Headers": "*",
	}
	multiValHeaders := map[string][]string{
		"Access-Control-Allow-Methods": []string{"POST"},
		"Access-Control-Allow-Headers": []string{"*"},
	}

	if strings.EqualFold(ev.HTTPMethod, "OPTIONS") {
		// cors
		return &events.APIGatewayProxyResponse{
			StatusCode:        200,
			Headers:           headers,
			MultiValueHeaders: multiValHeaders,
		}, nil
	}

	if !strings.EqualFold(ev.HTTPMethod, "POST") {
		return nil, errors.New("method POST supported only")
	}

	req := request{}
	err := json.Unmarshal([]byte(ev.Body), &req)
	if err != nil {
		return nil, err
	}

	if len(req.NodeURL) <= 0 {
		return nil, errors.New("node_url is required")
	}

	if req.BlockHeight == nil && len(req.TxHash) <= 0 {
		return nil, errors.New("block_height or tx_hash is required")
	}

	var out interface{}
	if req.BlockHeight != nil {
		out, err = getTxsAtBlockHeight(req.NodeURL, *req.BlockHeight)
		if err != nil {
			return nil, err
		}
	}

	buf, err := json.Marshal(out)
	if err != nil {
		return nil, err
	}

	var ret string = string(buf)

	headers["Content-Type"] = "application/json"
	// if hasGzipEncoding(ev.Headers) {
	// 	buf := new(bytes.Buffer)
	// 	gw := gzip.NewWriter(buf)
	// 	gw.Write([]byte(ret))
	// 	gw.Close()
	// 	ret = buf.String()

	// 	// from here all will be in gzip
	// 	headers["Content-Encoding"] = "gzip"
	// 	// write into gzip

	// 	// var buffer bytes.Buffer
	// 	// zw := gzip.NewWriter(&buffer)
	// 	// _, err = zw.Write(buf)
	// 	// if err != nil {
	// 	// 	return nil, err
	// 	// }

	// 	// if err := zw.Close(); err != nil {
	// 	// 	return nil, err
	// 	// }
	// 	// buf = buffer.Bytes()
	// }

	headers["Content-Length"] = strconv.Itoa(len(ret))

	return &events.APIGatewayProxyResponse{
		StatusCode:        200,
		Body:              ret,
		Headers:           headers,
		MultiValueHeaders: multiValHeaders,
	}, nil
}

func hasGzipEncoding(m map[string]string) bool {
	for k, v := range m {
		if strings.EqualFold(k, "Accept-Encoding") && strings.EqualFold(v, "gzip") {
			return true
		}
	}
	return false
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
