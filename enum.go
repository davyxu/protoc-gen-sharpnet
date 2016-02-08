package main

import (
	"github.com/davyxu/pbmeta"
)

func printEnum(gen *Generator, enum *pbmeta.EnumDescriptor) {
	gen.Println("[global::ProtoBuf.ProtoContract(Name=@\"", enum.Name(), "\")]")
	gen.Println("public enum ", enum.Name())
	gen.Println("{")
	gen.In()

	for i := 0; i < enum.ValueCount(); i++ {

		evd := enum.Value(i)

		gen.Println("[global::ProtoBuf.ProtoEnum(Name=@\"", evd.Name(), "\", Value=", evd.Value(), ")]")
		gen.BeginLine()
		gen.Print(evd.Name(), " = ", evd.Value())

		if i < enum.ValueCount()-1 {
			gen.Print(",")
		}

		gen.EndLine()
		gen.Println()
	}

	gen.Out()
	gen.Println("}")

}
