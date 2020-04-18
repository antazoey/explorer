package main

import (
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	"errors"
	_ "fmt"
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
}

func unpackSignedTx(tx []byte) (interface{}, error) {
	txBundle := types.SignedBundle{}
	err := proto.Unmarshal(tx, &txBundle)
	if err != nil {
		return nil, err
	}

	protoCmd, cmd, err := txDecode(txBundle.Data)
	if err != nil {
		return nil, err
	}

	cmdTx, err := unmarshalCommand(protoCmd, cmd)
	if err != nil {
		return nil, err
	}

	hexPubKey := hex.EncodeToString(txBundle.GetPubKey())

	m := jsonpb.Marshaler{}
	marshalTx, err := m.MarshalToString(cmdTx)
	if err != nil {
		return nil, err
	}

	return &SignedTx{
		Type:    cmd.String(),
		Command: marshalTx,
		Sig:     txBundle.Sig,
		PubKey:  "0x" + hexPubKey,
	}, nil
}

func unpackUnsignedTx(tx []byte) (interface{}, error) {
	protoCmd, cmd, err := txDecode(tx)
	if err != nil {
		return nil, err
	}

	cmdTx, err := unmarshalCommand(protoCmd, cmd)
	if err != nil {
		return nil, err
	}

	m := jsonpb.Marshaler{}
	marshalTx, err := m.MarshalToString(cmdTx)
	if err != nil {
		return nil, err
	}

	return &UnsignedTx{
		Type:    cmd.String(),
		Command: marshalTx,
	}, nil
}

func txDecode(input []byte) ([]byte, Command, error) {
	if len(input) > 37 {
		return input[37:], Command(input[36]), nil
	}
	return nil, 0, errors.New("payload size is incorrect, should be > 38 bytes")
}

func unpack(tx []byte) (interface{}, error) {
	switch CommandKind(tx[0]) {
	case CommandKindSigned:
		return unpackSignedTx(tx[1:])
	case CommandKindUnsigned:
		return unpackUnsignedTx(tx[1:])
	default:
		return nil, errors.New("unsupported transaction")
	}
}

const (
	chainAddress = "https://geo.s.vega.xyz:8443"
)

type request struct {
	NodeURL     string  `json:"node_url"`
	BlockHeight *uint64 `json:"block_height"`
	TxHash      string  `json:"tx_hash"`
}

func handler(ev events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
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
			StatusCode: 200,
			Headers: headers,
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

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(buf),
		Headers: headers,
		MultiValueHeaders: multiValHeaders,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
