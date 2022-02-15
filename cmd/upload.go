package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/comeonjy/upic/config"
	"github.com/comeonjy/upic/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var uploadCmd = &cobra.Command{
	Use:   "up",
	Short: "upload pic",
	Long:  "上传图片",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请输入文件名")
			return
		}
		conf := config.NewConfig(context.Background())
		var url string
		var err error
		switch viper.GetString("oss") {
		case "qiniu":
			client := pkg.NewQiNiu(conf.Get().QiniuAK, conf.Get().QiniuSK, conf.Get().QiniuCDN)
			url, err = client.Upload(args[0])
		}
		if err != nil {
			log.Println(cmd.Context(), "Update pic err:%v", err)
		}
		log.Println("上传成功", url)
		log.Println("Markdown", fmt.Sprintf("![](%s)", url))
	},
}

func init() {
	uploadCmd.PersistentFlags().StringP("oss", "s", "qiniu", "选择一个对象存储")
	viper.BindPFlag("oss", uploadCmd.PersistentFlags().Lookup("oss"))
}
