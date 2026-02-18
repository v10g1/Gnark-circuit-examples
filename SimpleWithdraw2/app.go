package simplewithdraw
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

type Circuit struct{
	oldBalance frontend.Variable
	newBalance frontend.Variable `gnark:",public"`
	withdrawAmount frontend.Variable
}

// I am using chatgpt to give me some question for Gnark circuits, so this is fun the below commented was my version

// func(circuit *Circuit) Define(api frontend.API) error{
// 	api.AssertIsLessOrEqual(circuit.withdrawAmount,circuit.oldBalance)
// 	api.AssertIsLessOrEqual(0,circuit.withdrawAmount)
// 	tmp :=api.Sub(circuit.oldBalance,circuit.withdrawAmount)
// 	api.AssertIsEqual(circuit.newBalance,tmp)
// 	return nil
// }

// This was my version i gave it to chatgpt
// it gave me some security consideration

// which were
//check the upper threshold for the balances for not make them absurd values
// so it will become like

func (c *Circuit) Define(api frontend.API) error {

    limit := uint64(1 << 32)

    // bounds (CRITICAL)
    api.AssertIsLessOrEqual(c.oldBalance, limit)
    api.AssertIsLessOrEqual(c.withdrawAmount, limit)
    api.AssertIsLessOrEqual(c.newBalance, limit)

    // overdraft protection
    api.AssertIsLessOrEqual(c.withdrawAmount, c.oldBalance)

    // arithmetic invariant
    expected := api.Sub(c.oldBalance, c.withdrawAmount)
    api.AssertIsEqual(c.newBalance, expected)

    return nil
}
// this is interesting