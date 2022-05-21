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

func SetImageData(ctx context.Context, value string) string {
	return utils.SetKey(ctx, "img", value)
}

func SetTitleData(ctx context.Context, value string) string {
	return utils.SetKey(ctx, "title", value)
}
