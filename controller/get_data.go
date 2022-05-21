package controller

import (
	"context"
	"github.com/zqb-knight/gsq_zqb_ios/utils"
)

func GetImageData(ctx context.Context) string {
	return utils.GetKey(ctx, "img")
}

func GetTitleData(ctx context.Context) string {
	return utils.GetKey(ctx, "title")
}
