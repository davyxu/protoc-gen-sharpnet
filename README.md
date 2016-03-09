# protoc-gen-sharpnet

## protobuf-net提供的protogen代码生成器缺陷

* 自解析proto 闭源 非标准proto

* 无法向后兼容 不支持proto3

* 解析速度慢


## 本品特性

* 使用官方protoc编译器插件架构编写 标准proto

* 向后兼容 支持proto3的基本语法特性( 暂时不支持 map, oneof等proto3特有特性 )

* 解析速度快


# 安装方法

	go get github.com/davyxu/protoc-gen-sharpnet
	
	go install github.com/davyxu/protoc-gen-sharpnet

# 使用方法

protoc --plugin=protoc-gen-sharpnet=path\to\protoc-gen-sharpnet --sharpnet_out . --proto_path "." PROTO_FILE

# 使用情况

本品在商业项目中已大规模使用, 请放心使用. 

若发现bug请邮件sunicdavy@qq.com

# 链接

* protobuf-net运行库

	https://github.com/mgravell/protobuf-net