package services

import (
	"context"
	"fmt"
)

type ProdService struct {

}

func (this *ProdService)GetProdStk(ctx context.Context, request *ProdRequest) (*ProdResponse, error){

	fmt.Println(request.ProdId,999999)
	fmt.Println(request.ProdArea,88888)


	stock :=int32(0)
	if request.ProdArea == ProdArea_A{
		stock=1024
	}else if request.ProdArea==ProdArea_B{
		stock =2048
	}else{
		stock=4096
	}

	//fmt.Println((&ProdRequest{ProdId:2223}).GetProdId())
	//fmt.Println((ProdRequest).GetProdId())

	//fmt.Println((*ProdRequest))

	return &ProdResponse{
		ProdStock: stock ,
	},nil
}


func (this *ProdService)GetProdStks(ctx context.Context, in *QuerySize) (*ProdStkList, error){
	//res := make([]ProdResponse,10)


	//fmt.Println(QuerySize.GetSize())
	res := []*ProdResponse{
		{ProdStock: 123},
		{ProdStock: 233},
		{ProdStock: 333},
		{ProdStock: 433},
		{ProdStock: 553},
	}

	return &ProdStkList{
		Prodres: res ,
	},nil
}

func (this *ProdService)GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error){
	return &ProdModel{
		ProdId: 12,
		ProdName: "asdut8 实打实的撒",
		ProdPrice: 12.2,
	},nil
}