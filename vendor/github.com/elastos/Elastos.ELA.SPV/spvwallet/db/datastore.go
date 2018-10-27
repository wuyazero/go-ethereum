package db

import (
	"github.com/elastos/Elastos.ELA/core"
	"github.com/elastos/Elastos.ELA.Utility/common"
)

type DataStore interface {
	Chain() Chain
	Addrs() Addrs
	Txs() Txs
	UTXOs() UTXOs
	STXOs() STXOs

	Rollback(height uint32) error
	RollbackTx(txId *common.Uint256) error
	// Reset database, clear all data
	Reset() error

	Close()
}

type Chain interface {
	// save chain height
	PutHeight(height uint32)

	// get chain height
	GetHeight() uint32
}

type Addrs interface {
	// put a address to database
	Put(hash *common.Uint168, script []byte, addrType int) error

	// get a address from database
	Get(hash *common.Uint168) (*Addr, error)

	// get all addresss from database
	GetAll() ([]*Addr, error)

	// delete a address from database
	Delete(hash *common.Uint168) error
}

type Txs interface {
	// Put a new transaction to database
	Put(txn *Tx) error

	// Fetch a raw tx and it's metadata given a hash
	Get(txId *common.Uint256) (*Tx, error)

	// Fetch all transactions from database
	GetAll() ([]*Tx, error)

	// Fetch all transactions from the given height
	GetAllFrom(height uint32) ([]*Tx, error)

	// Delete a transaction from the db
	Delete(txId *common.Uint256) error
}

type UTXOs interface {
	// put a utxo to database
	Put(hash *common.Uint168, utxo *UTXO) error

	// get a utxo from database
	Get(outPoint *core.OutPoint) (*UTXO, error)

	// get utxos of the given address hash from database
	GetAddrAll(hash *common.Uint168) ([]*UTXO, error)

	// Get all UTXOs in database
	GetAll() ([]*UTXO, error)

	// delete a utxo from database
	Delete(outPoint *core.OutPoint) error
}

type STXOs interface {
	// Move a UTXO to STXO
	FromUTXO(outPoint *core.OutPoint, spendTxId *common.Uint256, spendHeight uint32) error

	// get a stxo from database
	Get(outPoint *core.OutPoint) (*STXO, error)

	// get stxos of the given address hash from database
	GetAddrAll(hash *common.Uint168) ([]*STXO, error)

	// Get all STXOs in database
	GetAll() ([]*STXO, error)

	// delete a stxo from database
	Delete(outPoint *core.OutPoint) error
}
