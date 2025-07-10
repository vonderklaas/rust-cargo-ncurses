package p2p

// Peer is an interface that represents the remote node connection
type Peer interface{}

// Transport is anything that handles the communication between
// the nodes in the network. This can be of the form TCP, UDP, and WS.
type Transport interface {
	ListenAndAccept() error
}
