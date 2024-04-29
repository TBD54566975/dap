package maddr_test

import (
	"testing"

	"github.com/TBD54566975/dap/maddr"
	"github.com/alecthomas/assert"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

func TestDecode(t *testing.T) {
	didpayUSDC := didcore.Service{
		Type:            maddr.MoneyAddrKind,
		ID:              "didpay",
		ServiceEndpoint: []string{"urn:usdc:eth:0x2345y7432"},
	}

	maddrs, err := maddr.FromDIDService(didpayUSDC)
	assert.NoError(t, err)
	assert.Len(t, maddrs, 1)

	m := maddrs[0]
	assert.Equal(t, m.Currency, "usdc")

	// muunBTC := didcore.Service{
	// 	Type:            maddr.MoneyAddrKind,
	// 	ID:              "muun",
	// 	ServiceEndpoint: "urn:btc:addr:m12345677axcv2345",
	// }

	// lnURL := didcore.Service{
	// 	Type:            maddr.MoneyAddrKind,
	// 	ID:              "lnurl",
	// 	ServiceEndpoint: "urn:btc:lnurl:https://someurl.com",
	// }
}
