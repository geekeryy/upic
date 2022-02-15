package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/comeonjy/go-kit/pkg/util"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiu struct {
	ak     string
	sk     string
	cdn    string
	bucket string
	prefix string
}

func NewQiNiu(ak, sk, cdn string) *QiNiu {
	return &QiNiu{
		ak:     ak,
		sk:     sk,
		cdn:    cdn,
		bucket: "public-11",
		prefix: "upic",
	}
}

func (q *QiNiu) Upload(localFile string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: q.bucket,
	}
	mac := qbox.NewMac(q.ak, q.sk)
	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	})
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:utm_source": "upic",
		},
	}
	key := util.Md5(localFile + time.Now().String())
	if err := formUploader.PutFile(context.Background(), &ret, upToken, q.prefix+"/"+key, localFile, &putExtra); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", q.cdn, ret.Key), nil
}
