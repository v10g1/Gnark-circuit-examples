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
	senderOld frontend.Variable `gnark:",public"`
	senderNew frontend.Variable
	recieverNew frontend.Variable `gnark:",public"`
	recieverOld frontend.Variable
	transferAmount frontend.Variable

}

func(c *Circuit) Define(api frontend.API) error{
	api.AssertIsLessOrEqual(c.transferAmount,c.senderOld)
	diff1 := api.Sub(c.senderNew,c.senderOld)
	diff2 := api.Sub(c.recieverOld,c.recieverNew)
	api.AssertIsEqual(diff1,diff2)
	api.AssertIsEqual(diff1,c.transferAmount)
	limit := uint64(1<<32)
	api.AssertIsLessOrEqual(c.senderOld,limit)
	api.AssertIsLessOrEqual(c.senderNew,limit)
	api.AssertIsLessOrEqual(c.recieverOld,limit)
	api.AssertIsLessOrEqual(c.recieverOld,limit)
	


	return nil
}