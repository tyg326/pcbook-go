package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/techschool/pcbook/pb"
	"gitlab.com/techschool/pcbook/sample"
	"gitlab.com/techschool/pcbook/serializer"
	"gitlab.com/techschool/pcbook/service"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	// 并行运行
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer(t)

	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedID := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, expectedID, res.Id)

	// check that the laptop is saved to the store
	other, err := laptopServer.LaptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotEmpty(t, other)

	// check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)

}

// StartTestLaptopServer
func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random avaliable port
	require.NoError(t, err)

	// grpcServer.Serve(listener) // block call
	go grpcServer.Serve(listener) // non block call

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	// require.Equal(t, laptop1, laptop2)

	json1 := serializer.ProtobufToJSON(laptop1)
	require.NotNil(t, json1)

	json2 := serializer.ProtobufToJSON(laptop2)
	require.NotNil(t, json2)

	require.Equal(t, json1, json2)

}
