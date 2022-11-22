package main

import (
	"embed"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type mychaincode struct{}

var (
	//go:embed crypto/extcc.key
	serverKey []byte

	//go:embed crypto/extcc.pem
	serverCert []byte

	//go:embed crypto/client.pem
	clientCACerts []byte

	_ embed.FS
)

// Init is called during Instantiate transaction after the chaincode container
// has been established for the first time, allowing the chaincode to
// initialize its internal data
func (m *mychaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called to update or query the ledger in a proposal transaction.
// Updated state variables are not committed to the ledger until the
// transaction is committed.
func (m *mychaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("Chaincode invoke success"))
}

func main() {
	server := &shim.ChaincodeServer{
		CCID:    os.Getenv("CHAINCODECCID"),
		Address: "0.0.0.0:9099",
		CC:      &mychaincode{},
		TLSProps: shim.TLSProperties{
			Disabled:      false,
			Key:           serverKey,
			Cert:          serverCert,
			ClientCACerts: clientCACerts,
		},
	}
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
