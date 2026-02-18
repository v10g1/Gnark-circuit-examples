package privateTransfers
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
	senderOld frontend.Variable
	
	r1Old frontend.Variable
	r2Old frontend.Variable
	r1new frontend.Variable `gnark:",public"`
	r2new frontend.Variable `gnark:",public"`
	senderNew frontend.Variable `gnark:",public"`
	amount1 frontend.Variable
	amount2 frontend.Variable
}
func (c* Circuit) Define(api frontend.API) error{
	
	diff1:= api.Sub(c.senderOld,c.senderNew)
	subdiff1:=api.Sub(c.r1new,c.r1Old)
	api.AssertIsEqual(c.amount1,subdiff1)
	
	subdiff2:=api.Sub(c.r2new,c.r2Old)
	api.AssertIsEqual(c.amount2,subdiff2)

	totalSent:= api.Add(subdiff1,subdiff2)
	api.AssertIsEqual(diff1,totalSent)
	api.AssertIsEqual(c.amount1,subdiff1)
	api.AssertIsEqual(c.amount2,subdiff2)
	limit:=uint64(1<<32)
	api.AssertIsLessOrEqual(c.senderOld,limit)
	api.AssertIsLessOrEqual(c.senderNew,limit)
	api.AssertIsLessOrEqual(c.r1Old,limit)
	api.AssertIsLessOrEqual(c.r1new,limit)
	api.AssertIsLessOrEqual(c.r2Old,limit)
	api.AssertIsLessOrEqual(c.r2new,limit)
	api.AssertIsLessOrEqual(c.amount1,limit)
	api.AssertIsLessOrEqual(c.amount2,limit)

	totalsentAmount:=api.Add(c.amount1,c.amount2)
	api.AssertIsEqual(diff1,totalsentAmount)
	api.AssertIsLessOrEqual(1,c.amount1)
	api.AssertIsLessOrEqual(1,c.amount2)


	return nil
}