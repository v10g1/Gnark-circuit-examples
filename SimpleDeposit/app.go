package app

import (
	// "bytes"
	// "encoding/hex"
	// "fmt"
	// "math/big"
	// "os"

	// "github.com/consensys/gnark-crypto/ecc"
	// "github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
)

type Circuit struct {
	newBalance frontend.Variable
	oldBalance frontend.Variable
	transferAmount frontend.Variable
}
func (circuit *Circuit) Define(api frontend.API) error{
	tmp := api.Add(circuit.oldBalance,circuit.transferAmount)
	api.AssertIsEqual(circuit.newBalance,tmp)
	return nil
}