package zcashd

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/base58"
)

// TODO: we should really break out the zcash address types into their own separate implementations
// TODO: the same way this is done in btcutil. For OpenBazaar purposes, however, we only need the string representation for now.

type ZcashAddress struct {
	address string
	params  *chaincfg.Params
}

func NewZcashAddress(addr string, net *chaincfg.Params) (*ZcashAddress, error) {
	return &ZcashAddress{addr, net}, nil
}

// EncodeAddress returns the string encoding of a zcash
// address.  Part of the Address interface.
func (a *ZcashAddress) EncodeAddress() string {
	return a.address
}

// FIXME: this function is not used in openbazaar but should still work properly
func (a *ZcashAddress) ScriptAddress() []byte {
	return base58.Decode(a.address)
}

// IsForNet returns whether or not the pay-to-pubkey-hash address is associated
// with the passed zcash network.
func (a *ZcashAddress) IsForNet(net *chaincfg.Params) bool {
	return a.params.Name == net.Name
}

// String returns a human-readable string for the pay-to-pubkey-hash address.
// This is equivalent to calling EncodeAddress, but is provided so the type can
// be used as a fmt.Stringer.
func (a *ZcashAddress) String() string {
	return a.EncodeAddress()
}
