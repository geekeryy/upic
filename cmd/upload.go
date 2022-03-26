package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/comeonjy/upic/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var QiniuAK = os.Getenv("QINIU_AK")
var QiniuSK = os.Getenv("QINIU_SK")
var QiniuCDN = os.Getenv("QINIU_CDN")

var uploadCmd = &cobra.Command{
	Use:   "up",
	Short: "upload pic",
	Long:  "上传图片",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请输入文件名")
			return
		}
		var url string
		var err error
		switch viper.GetString("oss") {
		case "qiniu":
			client := pkg.NewQiNiu(QiniuAK, QiniuSK, QiniuCDN)
			url, err = client.Upload(args[0])
		}
		if err != nil {
			log.Println("Update pic err:", err)
			return
		}
		log.Println("上传成功", url)
		log.Println("Markdown", fmt.Sprintf("![](%s)", url))
	},
}

func init() {
	uploadCmd.PersistentFlags().StringP("oss", "s", "qiniu", "选择一个对象存储")
	viper.BindPFlag("oss", uploadCmd.PersistentFlags().Lookup("oss"))
}
