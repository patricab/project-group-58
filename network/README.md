# Network module for p2p elevators

network.go contains functions to communicate with elevator network

<br/>

## Install
Add these lines to your go.mod file:

```go
require network v0.0.0
replace network => <relative path to module>
```

## Usage
The user interface consists of a handler function that is called as a goroutine:

```go
go network.Handler(id, tx, rx)
``` 
The handler takes in 3 parameters:
- id: Custom ID for the elevator node
- tx: Tx (transmission) channel: `tx := make(chan network.Msg)`
- rx: Rx (transmission) channel: `rx := make(chan network.Msg)`

A message can then be sent through the _Tx_ channel, and recieved through
the _Rx_ channel by regular channel I/O operations.

<br/>

## Formatting
### ID
There are 4 options for IDs, or more, depending on how many elevators that are connected
to the network.

```
(*0)  1   2   3
```

Per default the individual addressable ID ragne is from _1_ to _3_.

The address _0_ is a special case address, where the transmitted message
goes out to every node in the network. This address should be used with commands
like **_cmdReqCost_**

<br/>

### Message
The standard format for a message is described below:

```go
type Msg struct {
	Id      int
	Dest    int
	Command Cmd
	Data    int // both cost and floor
}
```

The _Id_ field denotes the same custom ID that the user passes to the Handler function. The _Dest_ field
is the ID that the user wishes to send a message to. 

(**NB**: When sending the _cmdReqCost_ command, the user should set this field to 0. See the [ID section](#id))

The _Data_ field is the data that the user wishes to send to the relevant node (e.g **cost** value, **floor**, or **ACK**)

The commands avilable are:
```go
const (
	CmdReqCost  Cmd = 0
	CmdCost         = 1
	CmdDelegate     = 2
)
```
- CmdReqCost - Request cost from all nodes on the network. (**NB**: Remeber to set the _Dest_ ID to 0 in order to address all nodes on the network). The _Data_ field should contain the floor value that the nodes should base their cost calculations on. 
- CmdCost - Reply to CmdReqCost command. The _Dest_ should be the ID of the node that requested the cost value. The _Data_ field should hold the calculated cost value.
- CmdDelegate - Delegation order to specific node in network to service a given floor. The _Data_ field should hold the floor level that the node should service.