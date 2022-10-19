package video

import (
	"410proj/apps/app/api/internal/svc"
	"410proj/apps/app/api/internal/types"
	"410proj/apps/user/rpc/user"
	"410proj/apps/videos/rpc/video"
	"410proj/pkg/jwtx"
	"410proj/pkg/xerr"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	imageFileName = "data"
	bucketName    = "410proj"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddLogic) Add(req *types.AddVideoReq) (resp *types.AddVideoResp, err error) {
	// 发布视频
	logx.Infof("Add Video Start..........")
	data, header, err := l.r.FormFile(imageFileName)
	if err != nil {
		return &types.AddVideoResp{
			VideoResponse: types.VideoResponse{
				StatusCode: int32(xerr.FileUploadErr),
				Msg:        "上传视频失败" + err.Error(),
			},
		}, nil
	}
	defer data.Close()
	// 需要从token中获取 userId
	userid, err := jwtx.GetUserId(l.svcCtx.Config.JwtAuth.AccessSecret, req.Token)
	if err != nil {
		return nil, err
	}
	// 判断userid是否存在
	_, err = l.svcCtx.UserRPC.Get(l.ctx, &user.UserItemReq{UserId: userid})
	if err != nil {
		return nil, err
	}

	filename := filepath.Base(header.Filename)
	finalName := fmt.Sprintf("%d_%s_%d", userid, filename, time.Now().Unix())
	saveVideoFile := filepath.Join("./data/video", finalName+".mp4")
	saveCoverFile := filepath.Join("./data/cover", finalName+".jpg")
	// 写入本地文件，用于生成封面
	logx.Infof("Write Video Into Disk..........")
	f, err := os.OpenFile(saveVideoFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return &types.AddVideoResp{
			VideoResponse: types.VideoResponse{
				StatusCode: int32(xerr.FileUploadErr),
				Msg:        "上传视频失败" + err.Error(),
			},
		}, nil
	}
	// 需要拷贝一份
	io.Copy(f, data) // 写入本地
	f.Close()
	logx.Infof("Write Done..........")
	// 生成封面
	// 使用 ffmpeg 提取视频第1秒处的帧作为封面
	cmd := exec.Command(l.svcCtx.Config.FfmpegExecPath, "-i", saveVideoFile, "-ss", "00:00:01", saveCoverFile)
	cmd.Run()

	ossfunc := func() {
		// 上传视频
		fvideo, err := os.OpenFile(saveVideoFile, os.O_RDONLY, 0666)
		if err != nil {
			return
		}
		defer fvideo.Close()

		bucket, err := l.svcCtx.OssClient.Bucket(bucketName)
		if err != nil {
			return
		}
		objpath := "videos/" + finalName + ".mp4"
		if err = bucket.PutObject(objpath, fvideo); err != nil {
			return
		}
		// 上传封面
		fcover, err := os.OpenFile(saveCoverFile, os.O_RDONLY, 0666)
		if err != nil {
			return
		}
		defer fcover.Close()

		objpath = "images/" + finalName + ".jpg"
		if err = bucket.PutObject(objpath, fcover); err != nil {
			return
		}
	}
	var videourl, coverurl string
	if l.svcCtx.Config.UseOSS {
		go ossfunc()
		videourl = fmt.Sprintf("https://410proj.oss-cn-chengdu.aliyuncs.com/videos/%s", finalName+".mp4")
		coverurl = fmt.Sprintf("https://410proj.oss-cn-chengdu.aliyuncs.com/images/%s", finalName+".jpg")
	} else {
		videourl = fmt.Sprintf("%s:%d/data/video/%s",
			l.svcCtx.Config.ServerPath, l.svcCtx.Config.Port, finalName+".mp4")
		coverurl = fmt.Sprintf("%s:%d/data/cover/%s",
			l.svcCtx.Config.ServerPath, l.svcCtx.Config.Port, finalName+".jpg")

	}

	logx.Infof("Call RPC for add record..........")
	// 发起rpc
	in := video.AddVideoReq{
		AuthorId: userid,
		PlayUrl:  videourl,
		CoverUrl: coverurl,
	}

	_, err = l.svcCtx.VideoRPC.Add(l.ctx, &in)
	if err != nil {
		return &types.AddVideoResp{
			types.VideoResponse{
				StatusCode: 0,
				Msg:        "上传失败" + err.Error(),
			},
		}, nil
	}
	logx.Infof("All Done..........")
	return &types.AddVideoResp{
		types.VideoResponse{
			StatusCode: 0,
			Msg:        "上传成功",
		},
	}, nil

}
