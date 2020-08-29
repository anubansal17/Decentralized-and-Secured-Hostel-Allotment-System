// For making invoke in the blockchain:-
package blockchain
    import (
        "fmt"
        "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
        "time"
)
// InvokeHello
func (setup *FabricSetup) InvokeHello(args []string) (string, error) {
    eventID := "eventInvoke"
    // Add data that will be visible in the proposal, like a description of
    the invoke request
    transientDataMap := make(map[string][]byte)
    transientDataMap["result"] = []byte("Transient data in hello invoke")
    reg, notifier, err :=
    setup.event.RegisterChaincodeEvent(setup.ChainCodeID, eventID)
    if err != nil {
        return "", err
    }
    defer setup.event.Unregister(reg)
    // Create a request (proposal) and send it
    response, err :=
    setup.client.Execute(channel.Request{ChaincodeID:
    setup.ChainCodeID, Fcn: "invoke", Args: [][]byte{[]byte("set"),
    []byte(args[0]), []byte(args[1]), []byte(args[2]),[]byte(args[3])},
    TransientMap: transientDataMap})
    if err != nil {
        return "", fmt.Errorf("failed to do traction : %v", err)
    }
    // Wait for the result of the submission
    select {
        case ccEvent := <-notifier:
        fmt.Printf("Received CC event: %s\n", ccEvent)
        case <-time.After(time.Second * 20):
        return "", fmt.Errorf("did NOT receive CC event for
        eventId(%s)", eventID)
    }
    return string(response.TransactionID), nil
}
