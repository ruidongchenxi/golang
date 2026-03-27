package main
import(
	"testing"
	"fmt"
)

func TestGetsub(t *testing.T) {
	res:=getSub(50,23)
	fmt.Println("差",res)
	
}