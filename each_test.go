package protoiter_test

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/goaux/protoiter"
	"github.com/goaux/results"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Example() {
	var _ protoiter.Files = protoregistry.GlobalFiles
	var _ protoiter.Types = protoregistry.GlobalTypes
	var _ descriptorpb.DescriptorProto

	for file := range protoiter.EachFile(protoregistry.GlobalFiles) {
		var _ protoreflect.FileDescriptor = file
		for field, value := range protoiter.EachField(file.Options().ProtoReflect()) {
			var _ protoreflect.FieldDescriptor = field
			var _ protoreflect.Value = value
		}
		for i, enum := range protoiter.Each(file.Enums()) {
			var _ int = i
			var _ protoreflect.EnumDescriptor = enum
			for field, value := range protoiter.EachField(enum.Options().ProtoReflect()) {
				var _ protoreflect.FieldDescriptor = field
				var _ protoreflect.Value = value
			}
		}
		for i, message := range protoiter.Each(file.Messages()) {
			var _ int = i
			var _ protoreflect.MessageDescriptor = message
			for field, value := range protoiter.EachField(message.Options().ProtoReflect()) {
				var _ protoreflect.FieldDescriptor = field
				var _ protoreflect.Value = value
			}
			for i, oneof := range protoiter.Each(message.Oneofs()) {
				var _ int = i
				var _ protoreflect.OneofDescriptor = oneof
				for field, value := range protoiter.EachField(oneof.Options().ProtoReflect()) {
					var _ protoreflect.FieldDescriptor = field
					var _ protoreflect.Value = value
				}
			}
		}
		for i, extension := range protoiter.Each(file.Extensions()) {
			var _ int = i
			var _ protoreflect.ExtensionDescriptor = extension
			for field, value := range protoiter.EachField(extension.Options().ProtoReflect()) {
				var _ protoreflect.FieldDescriptor = field
				var _ protoreflect.Value = value
			}
		}
		for i, service := range protoiter.Each(file.Services()) {
			var _ int = i
			var _ protoreflect.ServiceDescriptor = service
			for field, value := range protoiter.EachField(service.Options().ProtoReflect()) {
				var _ protoreflect.FieldDescriptor = field
				var _ protoreflect.Value = value
			}
			for i, method := range protoiter.Each(service.Methods()) {
				var _ int = i
				var _ protoreflect.MethodDescriptor = method
				for field, value := range protoiter.EachField(method.Options().ProtoReflect()) {
					var _ protoreflect.FieldDescriptor = field
					var _ protoreflect.Value = value
				}
			}
		}
	}

	for file := range protoiter.EachFileByPackage(protoregistry.GlobalFiles, "") {
		var _ protoreflect.FileDescriptor = file
	}

	for enum := range protoiter.EachEnum(protoregistry.GlobalTypes) {
		var _ protoreflect.EnumType = enum
	}
	for message := range protoiter.EachMessage(protoregistry.GlobalTypes) {
		var _ protoreflect.MessageType = message
	}
	for extension := range protoiter.EachExtension(protoregistry.GlobalTypes) {
		var _ protoreflect.ExtensionType = extension
	}
	for extension := range protoiter.EachExtensionByMessage(protoregistry.GlobalTypes, "") {
		var _ protoreflect.ExtensionType = extension
	}
	// Output:
}

func ExampleEach() {
	var _ timestamppb.Timestamp
	file := results.Must1(protoregistry.GlobalFiles.FindFileByPath("google/protobuf/timestamp.proto"))
	for i, message := range protoiter.Each(file.Messages()) {
		fmt.Println(i, message.FullName())
	}
	// Output:
	// 0 google.protobuf.Timestamp
}

func ExampleEachField() {
	now := timestamppb.New(time.Unix(123, 456))
	for field, value := range protoiter.EachField(now.ProtoReflect()) {
		fmt.Println(field.FullName(), value, reflect.TypeOf(value.Interface()))
	}
	// Unordered output:
	// google.protobuf.Timestamp.seconds 123 int64
	// google.protobuf.Timestamp.nanos 456 int32
}

func ExampleEachEnum() {
	for enumType := range protoiter.EachEnum(protoregistry.GlobalTypes) {
		var _ protoreflect.EnumType = enumType
	}
	// Output:
}

func ExampleEachMessage() {
	for messageType := range protoiter.EachMessage(protoregistry.GlobalTypes) {
		var _ protoreflect.MessageType = messageType
	}
	// Output:
}

func ExampleEachExtension() {
	for extensionType := range protoiter.EachExtension(protoregistry.GlobalTypes) {
		var _ protoreflect.ExtensionType = extensionType
	}
	// Output:
}

func ExampleEachFile() {
	for file := range protoiter.EachFile(protoregistry.GlobalFiles) {
		var _ protoreflect.FileDescriptor = file
	}
	// Output:
}

type testDescriptor struct {
	protoreflect.Descriptor
	index int
}

func (d testDescriptor) Index() int { return d.index }

type testDescriptors struct{}

func (testDescriptors) Len() int { return 5 }

func (testDescriptors) Get(i int) testDescriptor { return testDescriptor{index: i + 1} }

func TestEach(t *testing.T) {
	var ii []int
	var di []int
	for i, desc := range protoiter.Each(testDescriptors{}) {
		ii = append(ii, i)
		di = append(di, desc.Index())
		if i == 2 {
			break
		}
	}
	if !slices.Equal(ii, []int{0, 1, 2}) {
		t.Errorf("index must be []int{0, 1, 2} got %v", ii)
	}
	if !slices.Equal(di, []int{1, 2, 3}) {
		t.Errorf("index must be []int{1, 2, 3} got %v", di)
	}
}

func TestEachField(t *testing.T) {
	got := make(map[string]any)
	now := timestamppb.New(time.Unix(123, 456))
	for field, value := range protoiter.EachField(now.ProtoReflect()) {
		var _ protoreflect.FieldDescriptor = field
		var _ protoreflect.Value = value
		got[string(field.FullName())] = value.Interface()
	}
	want := map[string]any{
		"google.protobuf.Timestamp.seconds": int64(123),
		"google.protobuf.Timestamp.nanos":   int32(456),
	}
	if !maps.Equal(got, want) {
		t.Errorf("must be equal\ngot\t%#v\nwant\t%#v", got, want)
	}
}
