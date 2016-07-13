package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/davyxu/pbmeta"
	pbprotos "github.com/davyxu/pbmeta/proto"
)

func fieldTypeString(fd *pbmeta.FieldDescriptor) string {

	var ret string

	switch fd.Type() {

	case pbprotos.FieldDescriptorProto_TYPE_INT32:
		ret = "int"
	case pbprotos.FieldDescriptorProto_TYPE_UINT32:
		ret = "uint"
	case pbprotos.FieldDescriptorProto_TYPE_BOOL:
		ret = "bool"
	case pbprotos.FieldDescriptorProto_TYPE_FLOAT:
		ret = "float"
	case pbprotos.FieldDescriptorProto_TYPE_DOUBLE:
		ret = "double"
	case pbprotos.FieldDescriptorProto_TYPE_STRING:
		ret = "string"
	case pbprotos.FieldDescriptorProto_TYPE_INT64:
		ret = "long"
	case pbprotos.FieldDescriptorProto_TYPE_UINT64:
		ret = "ulong"
	case pbprotos.FieldDescriptorProto_TYPE_BYTES:
		ret = "byte[]"
	case pbprotos.FieldDescriptorProto_TYPE_ENUM,
		pbprotos.FieldDescriptorProto_TYPE_MESSAGE:
		ret = fd.FullTypeName()
	default:
		ret = "unknown"
	}

	if fd.IsRepeated() {
		return fmt.Sprintf("global::System.Collections.Generic.List<%s>", ret)
	}

	return ret
}

func getDataFormat(fd *pbmeta.FieldDescriptor) string {
	switch fd.Type() {
	case pbprotos.FieldDescriptorProto_TYPE_STRING,
		pbprotos.FieldDescriptorProto_TYPE_MESSAGE,
		pbprotos.FieldDescriptorProto_TYPE_BOOL:
		return "Default"
	case pbprotos.FieldDescriptorProto_TYPE_FLOAT:
		return "FixedSize"
	}

	return "TwosComplement"
}

func wrapDefaultValue(fd *pbmeta.FieldDescriptor, typestr string) string {
	v := strings.TrimSpace(fd.DefaultValue())
	if v != "" {
		return fmt.Sprintf("(%s)%s", typestr, v)
	}

	return fmt.Sprintf("default(%s)", typestr)
}

func getDefaultValue(fd *pbmeta.FieldDescriptor) string {
	switch fd.Type() {
	case pbprotos.FieldDescriptorProto_TYPE_INT32:
		return wrapDefaultValue(fd, "int")
	case pbprotos.FieldDescriptorProto_TYPE_UINT32:
		return wrapDefaultValue(fd, "uint")
	case pbprotos.FieldDescriptorProto_TYPE_BOOL:
		return wrapDefaultValue(fd, "bool")
	case pbprotos.FieldDescriptorProto_TYPE_FLOAT:
		return wrapDefaultValue(fd, "float")
	case pbprotos.FieldDescriptorProto_TYPE_DOUBLE:
		return wrapDefaultValue(fd, "double")
	case pbprotos.FieldDescriptorProto_TYPE_INT64:
		return wrapDefaultValue(fd, "long")
	case pbprotos.FieldDescriptorProto_TYPE_UINT64:
		return wrapDefaultValue(fd, "ulong")
	case pbprotos.FieldDescriptorProto_TYPE_BYTES:
		return wrapDefaultValue(fd, "byte[]")
	case pbprotos.FieldDescriptorProto_TYPE_STRING:
		v := strings.TrimSpace(fd.DefaultValue())
		if v != "" {
			return "@" + strconv.Quote(fd.DefaultValue())
		}

		return strconv.Quote(fd.DefaultValue())
	case pbprotos.FieldDescriptorProto_TYPE_ENUM:
		ed := fd.EnumDesc()

		if ed == nil {
			return fd.DefaultValue()
		}

		if ed.ValueCount() > 0 {

			var defaultValue string

			// 有defaultvalue, 直接取值, 否则取枚举第一个
			if v := strings.TrimSpace(fd.DefaultValue()); v != "" {
				defaultValue = v
			} else {
				defaultValue = ed.Value(0).Name()
			}

			return fmt.Sprintf("%s.%s", fd.FullTypeName(), defaultValue)
		}
	case pbprotos.FieldDescriptorProto_TYPE_MESSAGE:
		return "null"
	}

	return "unknown"
}

func isRequired(fd *pbmeta.FieldDescriptor) string {
	if fd.IsRequired() {
		return "true"
	}

	return "false"
}

func printField(gen *Generator, fd *pbmeta.FieldDescriptor, msg *pbmeta.Descriptor, file *pbmeta.FileDescriptor) {
	typeStr := fieldTypeString(fd)
	memberVar := "_" + fd.Name()

	//  private int _Age = default(int);

	if fd.IsRepeated() {
		gen.Println("readonly ", typeStr, " ", memberVar, " = new ", typeStr, "();")

	} else {

		gen.Println(typeStr, " ", memberVar, " = ", getDefaultValue(fd), ";")
		gen.Println("bool _has", fd.Name(), " = false;")

	}

	// [global::ProtoBuf.ProtoMember(10, IsRequired = false, Name=@"Age", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
	gen.BeginLine()

	gen.Print("[global::ProtoBuf.ProtoMember(", fd.Define.GetNumber())

	if !fd.IsRepeated() {
		gen.Print(", IsRequired = ", isRequired(fd))
	}

	gen.Print(", Name=@\"", fd.Name(), "\", DataFormat = global::ProtoBuf.DataFormat.", getDataFormat(fd), ")]")

	gen.EndLine()

	if !fd.IsRepeated() {
		gen.Println("[global::System.ComponentModel.DefaultValue(", getDefaultValue(fd), ")]")
	}

	// 属性定义
	gen.Println("public ", typeStr, " ", fd.Name())
	gen.Println("{")
	gen.In()
	gen.Println("get { return ", memberVar, "; }")

	if !fd.IsRepeated() {
		gen.Println("set { ", memberVar, " = value; ")
		gen.Println("      _has", fd.Name(), " = true;")
		gen.Println("}")
	}

	gen.Out()
	gen.Println("}")

	gen.Println()

	if !fd.IsRepeated() {
		gen.Println("public bool Has", fd.Name())
		gen.Println("{")
		gen.In()
		gen.Println("get { return _has", fd.Name(), "; }")
		gen.Println("set { _has", fd.Name(), " = value; }")
		gen.Out()
		gen.Println("}")
	}

}
