package api

import (
	"context"
	"file_server/api/proto"
	"file_server/pkg/upload"
)

func (s *FileSvr) UploadLocal(ctx context.Context, req *proto.UploadLocalReq) (ret *proto.UploadLocalRes, err error) {
	ret, err = upload.Local(req.Name, req.FileBytes)
	res := &proto.R{Code: 600}
	if err != nil {
		r := new(proto.UploadLocalRes)
		res.Msg = err.Error()
		r.Res = res
		return
	}
	res.Code = 200
	ret.Res = res
	return
}
