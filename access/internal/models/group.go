package models

import pb "app/access/pkg/proto/gen"

type Group struct {
	ID      uint64
	Name    string
	Desc    string
	Users   []uint64
	Methods []Method
}

func (g *Group) Proto() *pb.Group {
	group := &pb.Group{
		Id:    g.ID,
		Name:  g.Name,
		Desc:  g.Desc,
		Users: g.Users,
	}

	var methods []*pb.Method
	for _, m := range g.Methods {
		methods = append(methods, m.Proto())
	}
	group.Methods = methods

	return group
}
