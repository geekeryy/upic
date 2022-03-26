# upic
图床工具

### 使用方式
go install 将工具安装到$GOPATH/bin目录下，需要将其添加到PATH方可直接使用，否则使用`$GOPATH/bin/upic up xxx.jpg`

可将环境变量写入.bash_profile方便后续直接使用upic命令
```
➜ go install github.com/comeonjy/upic@main
➜ export QINIU_CDN=http://xxx.xxx.com
➜ export QINIU_AK=xxxxxx
➜ export QINIU_SK=xxxxxx
➜ upic up ~/demo.png
2022/03/26 18:00:52 上传成功 http://xxx.xxx.com/upic/670a0bf2c54812d90d628141461936bd
2022/03/26 18:00:52 Markdown ![](http://xxx.xxx.com/upic/670a0bf2c54812d90d628141461936bd)
```

