--idl：接口定义 
source /etc/profile
kitex安装(不要安装高版本,会出现兼容问题): go install code.byted.org/kite/kitex/tool/cmd/kitex@v1.15.3
代码生成命令：kitexgit -service exclusive_base_qz idl/exclusive_base_qz.thrift


go env -w GOPROXY="https://go-mod-proxy.byted.org,https://goproxy.cn,https://proxy.golang.org,direct"
go env -w GOPRIVATE="*.byted.org,*.everphoto.cn,git.smartisan.com"
go env -w GOSUMDB="sum.golang.google.cn"
