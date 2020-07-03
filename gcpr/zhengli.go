package gcprs

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProdServiceClient interface {
	GetProdStk(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
}
type prodServiceClient struct {
	cc *grpc.ClientConn
}
func (c *prodServiceClient) GetProdStk(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/services.ProdService/getProdStk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ProdServiceServer interface {
	GetProdStk(context.Context, *ProdRequest) (*ProdResponse, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStk(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStk not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}