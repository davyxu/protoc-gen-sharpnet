package main

import (
	"github.com/davyxu/pbmeta"
)

func printMessage(gen *Generator, msg *pbmeta.Descriptor, file *pbmeta.FileDescriptor) {
	gen.Println("[global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@\"", msg.Name(), "\")]")
	gen.Println("public partial class ", msg.Name(), " : global::ProtoBuf.IExtensible")
	gen.Println("{")
	gen.In()

	for i := 0; i < msg.NestedMsg.MessageCount(); i++ {

		nestedMsg := msg.NestedMsg.Message(i)
		printMessage(gen, nestedMsg, file)
	}

	gen.Println("public ", msg.Name(), "() {}")
	gen.Println()

	for i := 0; i < msg.FieldCount(); i++ {

		fd := msg.Field(i)
		printField(gen, fd, msg, file)
		gen.Println()
	}

	for i := 0; i < msg.EnumCount(); i++ {

		enum := msg.Enum(i)
		printEnum(gen, enum)
	}

	gen.Println("private global::ProtoBuf.IExtension extensionObject;")
	gen.Println("global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)")
	gen.Println("{ return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }")

	gen.Out()
	gen.Println("}")

}
