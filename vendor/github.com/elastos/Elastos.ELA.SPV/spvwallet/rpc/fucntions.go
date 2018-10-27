package rpc

import (
	"bytes"
	"encoding/hex"

	"github.com/elastos/Elastos.ELA/core"
)

func (server *Server) notifyNewAddress(req Req) Resp {
	data, ok := req.Params[0].(string)
	if !ok {
		return InvalidParameter
	}
	addr, err := hex.DecodeString(data)
	if err != nil {
		return FunctionError(err.Error())
	}
	server.NotifyNewAddress(addr)
	return Success("New address received")
}

func (server *Server) sendTransaction(req Req) Resp {
	data, ok := req.Params[0].(string)
	if !ok {
		return InvalidParameter
	}
	txBytes, err := hex.DecodeString(data)
	if err != nil {
		return FunctionError(err.Error())
	}
	var tx core.Transaction
	err = tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return FunctionError("Deserialize transaction failed")
	}
	txId, err := server.SendTransaction(tx)
	if err != nil {
		return FunctionError(err.Error())
	}
	return Success(txId.String())
}
