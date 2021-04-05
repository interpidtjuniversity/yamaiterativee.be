package invoke

import (
	"google.golang.org/grpc"
	"sync"
)

var gcp GrpcConnPool

type GrpcConnPool struct {
	connections []*grpc.ClientConn
	busy        []bool
	cond        *sync.Cond
	mu          sync.Mutex
}

func GetConnection() *grpc.ClientConn{
	var c *grpc.ClientConn
	for c == nil {
		gcp.mu.Lock()
		for i := 0; i < len(gcp.busy); i++ {
			if !gcp.busy[i] {
				gcp.busy[i] = true
				c=gcp.connections[i]
			}
		}
		gcp.mu.Unlock()
		if c==nil {
			gcp.cond.Wait()
		}
	}
	return c
}

func Return(c *grpc.ClientConn) {
	gcp.mu.Lock()
	for i := 0; i < len(gcp.connections); i++ {
		if gcp.connections[i] == c {
			gcp.busy[i] = false
			gcp.cond.Signal()
			break
		}
	}
	gcp.mu.Unlock()
}

func ShutDownGcp() error {
	for i := 0; i < len(gcp.connections); i++ {
		err := gcp.connections[i].Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	size := 5
	serverAddr := "localhost:8000"

	gcp = GrpcConnPool{connections: make([]*grpc.ClientConn, size), busy: make([]bool, size), cond: sync.NewCond(&gcp.mu)}
	for i := 0; i < size; i++ {
		conn, _ := grpc.Dial(serverAddr, grpc.WithInsecure())
		gcp.connections[i] = conn
	}
}
