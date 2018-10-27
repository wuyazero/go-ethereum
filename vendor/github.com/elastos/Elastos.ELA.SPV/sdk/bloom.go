package sdk

import (
	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ELA/bloom"
	"github.com/elastos/Elastos.ELA/core"
)

// Create a new bloom filter instance
// elements are how many elements will be added to this filter.
func NewBloomFilter(elements uint32) *bloom.Filter {
	return bloom.NewFilter(elements, 0, 0.00003)
}

// Build a bloom filter by giving the interested addresses and outpoints
func BuildBloomFilter(addresses []*common.Uint168, outpoints []*core.OutPoint) *bloom.Filter {
	elements := uint32(len(addresses) + len(outpoints))

	filter := NewBloomFilter(elements)
	for _, address := range addresses {
		filter.Add(address.Bytes())
	}

	for _, outpoint := range outpoints {
		filter.AddOutPoint(outpoint)
	}

	return filter
}

// Add a address into the given bloom filter
func FilterAddress(filter *bloom.Filter, address *common.Uint168) {
	filter.Add(address.Bytes())
}

// Add a account into the given bloom filter
func FilterAccount(filter *bloom.Filter, account *Account) {
	filter.Add(account.programHash.Bytes())
}

// Add a outpoint into the given bloom filter
func FilterOutpoint(filter *bloom.Filter, op *core.OutPoint) {
	filter.AddOutPoint(op)
}
