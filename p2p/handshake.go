package p2p

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
