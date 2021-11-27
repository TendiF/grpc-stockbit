package chat

import (
	"context"
	"log"

	"github.com/TendiF/grpc-stockbit/proto"
)

type Server struct {
	proto.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *proto.Message) (*proto.Message, error) {

	log.Print("Message: ", message.Body)

	return message, nil
}
