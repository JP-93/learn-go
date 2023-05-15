package sever

import (
	"context"

	"m/v2/pb"
)

type Translation struct {
	pb.UnimplementedTranslationServer
}

func NewTranslation() *Translation {
	return &Translation{}
}

func (t *Translation) Translate(ctx context.Context, i *pb.TranslationInput) (*pb.TranslationOutput, error) {
	var c
}
