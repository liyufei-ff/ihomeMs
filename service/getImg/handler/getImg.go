package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"ihomeMs/service/getImg/model"
	getImg "ihomeMs/service/getImg/proto/getImg"
	"ihomeMs/service/getImg/utils"
	"image/color"
)

type GetImg struct{}

// MicroGetImg Call is a single request handler called via client.Call or the generated client code
func (e *GetImg) MicroGetImg(ctx context.Context, req *getImg.Request, rsp *getImg.Response) error {
	//初始化生成图片验证码对象
	cap := captcha.New()

	//设置字体
	cap.SetFont("./comic.ttf")

	//设置验证码大小
	cap.SetSize(128, 64)

	//设置干扰强度
	cap.SetDisturbance(captcha.HIGH)

	//设置颜色
	cap.SetFrontColor(color.RGBA{R: 255, A: 128})

	//设置背景se
	cap.SetBkgColor(color.RGBA{R: 128, A: 111})

	//生成验证码图片
	img, rnd := cap.Create(4, captcha.NUM)

	//存储验证码   redis
	err := model.SaveImgRnd(req.Uuid, rnd)
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}

	//传递图片信息给调用者
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	//json序列化
	imgJson, _ := json.Marshal(img)
	rsp.Data = imgJson
	return nil
}
