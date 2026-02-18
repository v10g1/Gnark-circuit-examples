package merkleLeafUpdate

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
const Depth=20
type Circuit struct{
	oldBalance frontend.Variable
	withdrawBalance frontend.Variable
	pathElements [Depth]frontend.Variable
	pathBits [Depth]frontend.Variable
	OldRoot frontend.Variable `gnark:",public"`
	NewRoot frontend.Variable `gnark:",public"`
	NewBalance frontend.Variable `gnark:",public"`
}
func (c *Circuit) Define(api frontend.API) error{
	limit := uint64(1<<32)
	api.AssertIsLessOrEqual(c.oldBalance,limit)
	api.AssertIsLessOrEqual(c.oldBalance,limit)
	api.AssertIsLessOrEqual(c.withdrawBalance,limit)

	api.AssertIsLessOrEqual(c.withdrawBalance,c.oldBalance)

	expected :=api.Sub(c.oldBalance,c.withdrawBalance)
	api.AssertIsEqual(c.NewBalance,expected)

	cur := c.oldBalance

	for i:=0;i<Depth;i++{
		left := api.Add(cur,c.pathElements[i])
		right := api.Add(c.pathElements[i],cur)

		cur = api.Select(c.pathBits[i],right,left)
	}
	api.AssertIsEqual(cur,c.OldRoot)
	for i := 0; i < Depth; i++ {

		left := api.Add(cur, c.pathElements[i])
		right := api.Add(c.pathElements[i], cur)

		cur = api.Select(c.pathBits[i], right, left)
	}

	api.AssertIsEqual(cur, c.NewRoot)
	return nil
}