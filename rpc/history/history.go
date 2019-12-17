package userssvc

import (
	"context"
	"fmt"

	pb "github.com/cagodoy/tenpo-challenge/lib/proto"
	history "github.com/cagodoy/tenpo-history-api"
	"github.com/cagodoy/tenpo-history-api/database"
	"github.com/cagodoy/tenpo-history-api/service"

	nats "github.com/nats-io/nats.go"
)

var _ pb.HistoryServiceServer = (*Service)(nil)

// Service ...
type Service struct {
	historySvc history.Service
}

// New ...
func New(store database.Store, conn *nats.EncodedConn) *Service {
	return &Service{
		historySvc: service.NewHistory(store, conn),
	}
}

// ListHistoryByUserId return a collection of users.
func (us *Service) ListHistoryByUserId(ctx context.Context, gr *pb.HistoryListByUserIdRequest) (*pb.HistoryListByUserIdResponse, error) {
	fmt.Println("[GRPC][HistoryService][ListHistoryByUserId][Request] empty = ")

	userID := gr.GetUserId()
	if userID == "" {
		fmt.Println("[GRPC][HistoryService][ListHistoryByUserId][Errir] invalid user_id req value")

		return &pb.HistoryListByUserIdResponse{
			Data: nil,
			Meta: nil,
			Error: &pb.Error{
				Message: "invalid user_id req value",
				Code:    500,
			},
		}, nil
	}

	listHistory, err := us.historySvc.ListByUserID(userID)
	if err != nil {
		fmt.Println(fmt.Sprintf("[GRPC][HistoryService][ListHistoryByUserId][Error] %v", err))

		return &pb.HistoryListByUserIdResponse{
			Data: nil,
			Meta: nil,
			Error: &pb.Error{
				Message: err.Error(),
				Code:    500,
			},
		}, nil
	}

	data := make([]*pb.History, 0)
	for _, history := range listHistory {
		data = append(data, history.ToProto())
	}

	res := &pb.HistoryListByUserIdResponse{
		Data:  data,
		Meta:  nil,
		Error: nil,
	}

	fmt.Println(fmt.Sprintf("[GRPC][HistoryService][ListHistoryByUserId][Response] %v", res))
	return res, nil
}
