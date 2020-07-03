package services

import (
	"context"
	"fmt"
)

type ProdService struct {

}

func (this *ProdService)GetProdStk(context.Context, *ProdRequest) (*ProdResponse, error){

	fmt.Println((*ProdRequest).GetProdId,999999)

	//fmt.Println((&ProdRequest{ProdId:2223}).GetProdId())
	//fmt.Println((ProdRequest).GetProdId())

	//fmt.Println((*ProdRequest))

	return &ProdResponse{
		ProdStock: 128 ,
	},nil
}