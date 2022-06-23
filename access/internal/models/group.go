package models

import pb "app/access/pkg/proto/gen"

type Group struct {
	ID       uint64
	Name     string
	Desc     string
	Users    []uint64
	Services []*Service
}

func (g *Group) Proto() *pb.Group {
	group := &pb.Group{
		Id:    g.ID,
		Name:  g.Name,
		Desc:  g.Desc,
		Users: g.Users,
	}

	var services []*pb.Service
	for _, s := range g.Services {
		services = append(services, s.Proto())
	}
	group.Services = services

	return group
}
