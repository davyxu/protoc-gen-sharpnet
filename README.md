# protoc-gen-sharpnet

## protobuf-net提供的protogen代码生成器缺陷

* 自解析proto 闭源 非标准proto

* 无法向后兼容 不支持proto3

* 解析速度慢


## 本品特性

* 使用官方protoc编译器插件架构编写 标准proto

* 向后兼容 支持proto3的基本语法特性( 暂时不支持 map, oneof等proto3特有特性 )

* 解析速度快

# 功能实现及限制

* 不建议使用嵌套结构及枚举, 不保证导出结果的正确性

# 安装方法

	go get github.com/davyxu/protoc-gen-sharpnet
	
	go install github.com/davyxu/protoc-gen-sharpnet

# 使用方法

* 兼容protogen输出的格式
	
	protoc --plugin=protoc-gen-sharpnet=path\to\protoc-gen-sharpnet --sharpnet_out . --proto_path "." PROTO_FILE
	
* 扩展支持hasfield

	protoc --plugin=protoc-gen-sharpnet=path\to\protoc-gen-sharpnet --sharpnet_out use_hasfield:. --proto_path "." PROTO_FILE

	P.S. HasField特性是C++库支持功能, 用于判断某字段是否被序列化(设置)过.
	性能上有bool赋值的轻度损耗以及内存轻度损耗

# 使用情况

本品在商业项目中已大规模使用, 请放心使用. 

若发现bug请邮件sunicdavy@qq.com

# 链接

* protobuf-net运行库

	https://github.com/mgravell/protobuf-net
	
	
# 备注

感觉不错请star, 谢谢!

博客: http://www.cppblog.com/sunicdavy

知乎: http://www.zhihu.com/people/xu-bo-62-87

邮箱: sunicdavy@qq.com
