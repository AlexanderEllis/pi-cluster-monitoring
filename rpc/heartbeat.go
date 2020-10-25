package heartbeat

// Args is the empty input to the heartbeat RPC.
type Args struct{}

// Reply is a simple boolean response.
type Reply struct {
	OK bool
}
