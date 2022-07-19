package server

import (
	pb "app/auth/pkg/proto/gen"
	"log"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateStream(_ *emptypb.Empty,
	stream pb.Auth_UpdateStreamServer) error {
	client := NewClient(stream)
	s.u.addClient(client)

	select {
	case <-client.Done:
		return nil
	}
}

type Client struct {
	Done   chan interface{}
	Stream pb.Auth_UpdateStreamServer
}

func NewClient(s pb.Auth_UpdateStreamServer) *Client {
	return &Client{
		Done:   make(chan interface{}),
		Stream: s,
	}
}

type Updater struct {
	mu    sync.RWMutex
	wg    sync.WaitGroup
	close chan interface{}

	nextClientID uint64
	clients      map[uint64]*Client
	events       <-chan *pb.UpdateStreamResponse
}

func NewUpdater(event <-chan *pb.UpdateStreamResponse) *Updater {

	u := &Updater{
		mu:           sync.RWMutex{},
		wg:           sync.WaitGroup{},
		close:        make(chan interface{}),
		nextClientID: 0,
		clients:      make(map[uint64]*Client),
		events:       event,
	}

	go u.Run()

	return u
}

func (u *Updater) Run() {
	u.wg.Add(1)
	for {
		select {
		case event, ok := <-u.events:
			if ok {
				u.sendUpdate(event)
			} else {
				log.Println("error")
				u.wg.Done()
				return
			}
		case <-u.close:
			for _, client := range u.clients {
				close(client.Done)
			}
			u.wg.Done()
			return
		}
	}
}

func (u *Updater) sendUpdate(msg *pb.UpdateStreamResponse) {
	removeIDs := make([]uint64, 0)
	defer u.removeClient(removeIDs)

	for id, client := range u.clients {
		err := client.Stream.Send(msg)
		if err != nil {
			removeIDs = append(removeIDs, id)
			close(client.Done)
		}
	}

}

func (u *Updater) Stop() {
	close(u.close)
	u.wg.Wait()
}

func (u *Updater) addClient(client *Client) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.clients[u.nextClientID] = client
}

func (u *Updater) removeClient(ids []uint64) {
	u.mu.Lock()
	defer u.mu.Unlock()

	newMapClients := make(map[uint64]*Client)
	newID := uint64(0)

	for _, client := range u.clients {
		if uint64(len(u.clients)) < newID {
			break
		}

		newMapClients[newID] = client
		newID++
	}

	u.clients = newMapClients
	u.nextClientID = uint64(len(newMapClients))
}
