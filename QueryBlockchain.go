// Code used to for transactions in the project -
// For making query in the blockchain:-
package blockchain
import (
"fmt"
"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)
// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello(value []string) (string, error) {
if len(value) != 2 {
fmt.Printf("Incorrect no of arguments")}
var args []string
args = append(args, "invoke")
args = append(args, value[0])
args = append(args, value[1])
response, err := setup.client.Query(channel.Request{ChaincodeID:
setup.ChainCodeID, Fcn: args[0], Args:
[][]byte{[]byte(args[1]),[]byte(args[2])}})
if err != nil {
return "", fmt.Errorf("failed to query: %v", err)
}
return string(response.Payload), nil
}
