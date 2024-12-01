// Package protoiter provides generic iterator functions for Protocol Buffers reflection.
package protoiter

import (
	"iter"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// Descriptors is an interface that abstracts the methods required to create an iterator over [protoreflect.Descriptors].
//
// It defines a generic way to access a collection of descriptors with a length and retrieval method.
// Descriptor is a generic type parameter representing the specific descriptor type (e.g. FileDescriptor, MessageDescriptor).
type Descriptors[Descriptor protoreflect.Descriptor] interface {
	Len() int
	Get(int) Descriptor
}

// Each creates a sequential iterator over a collection of descriptors.
// It allows iterating through descriptors with their indices.
//
// Parameters:
//   - dd: A collection of descriptors implementing the [Descriptors] interface
//
// Returns:
//   - An iterator sequence that yields the index and descriptor for each item
func Each[DD Descriptors[D], D protoreflect.Descriptor](dd DD) iter.Seq2[int, D] {
	return func(yield func(int, D) bool) {
		for i := range dd.Len() {
			if !yield(i, dd.Get(i)) {
				break
			}
		}
	}
}

// Files is an interface that abstracts the methods required to create an iterator over [google.golang.org/protobuf/reflect/protoregistry.Files].
type Files interface {
	RangeFiles(f func(protoreflect.FileDescriptor) bool)
	RangeFilesByPackage(name protoreflect.FullName, f func(protoreflect.FileDescriptor) bool)
}

// EachFile creates a sequential iterator over all file descriptors.
//
// It returns an iterator of calling [google.golang.org/protobuf/reflect/protoregistry.Files.RangeFiles].
//
//	RangeFiles iterates over all registered files while f returns true. If multiple files have the same name, RangeFiles iterates over all of them. The iteration order is undefined.
//
// Parameters:
//   - files: A Files implementation providing access to file descriptors
//
// Returns:
//   - An iterator sequence that yields each file descriptor
func EachFile(files Files) iter.Seq[protoreflect.FileDescriptor] {
	return func(yield func(protoreflect.FileDescriptor) bool) {
		files.RangeFiles(yield)
	}
}

// EachFileByPackage creates a sequential iterator over file descriptors in a specific package.
//
// It returns an iterator of calling [google.golang.org/protobuf/reflect/protoregistry.Files.RangeFilesByPackage].
//
//	RangeFilesByPackage iterates over all registered files in a given proto package while f returns true. The iteration order is undefined.
//
// Parameters:
//   - files: A Files implementation providing access to file descriptors
//   - name: The full package name to filter file descriptors
//
// Returns:
//   - An iterator sequence that yields file descriptors within the specified package
func EachFileByPackage(files Files, name protoreflect.FullName) iter.Seq[protoreflect.FileDescriptor] {
	return func(yield func(protoreflect.FileDescriptor) bool) {
		files.RangeFilesByPackage(name, yield)
	}
}

// Types is an interface that abstracts the methods required to create an iterator over [google.golang.org/protobuf/reflect/protoregistry.Types].
//
// It provides methods to range over different types of protocol buffer descriptors.
type Types interface {
	RangeEnums(f func(protoreflect.EnumType) bool)
	RangeMessages(f func(protoreflect.MessageType) bool)
	RangeExtensions(f func(protoreflect.ExtensionType) bool)
	RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool)
}

// EachEnum creates a sequential iterator over enum types.
//
// It returns an iterator of calling [google.golang.org/protobuf/reflect/protoregistry.Types.RangeEnums].
//
//	RangeEnums iterates over all registered enums while f returns true. Iteration order is undefined.
//
// Parameters:
//   - types: A Types implementation providing access to enum types
//
// Returns:
//   - An iterator sequence that yields each enum type
func EachEnum(types Types) iter.Seq[protoreflect.EnumType] {
	return func(yield func(protoreflect.EnumType) bool) {
		types.RangeEnums(yield)
	}
}

// EachMessage creates a sequential iterator over message types.
//
// It returns an iterator of calling [google.golang.org/protobuf/reflect/protoregistry.Types.RangeMessages].
//
//	RangeMessages iterates over all registered messages while f returns true. Iteration order is undefined.
//
// Parameters:
//   - types: A Types implementation providing access to message types
//
// Returns:
//   - An iterator sequence that yields each message type
func EachMessage(types Types) iter.Seq[protoreflect.MessageType] {
	return func(yield func(protoreflect.MessageType) bool) {
		types.RangeMessages(yield)
	}
}

// EachExtension creates a sequential iterator over extension types.
//
// It returns an iterator of calling [google.golang.org/protobuf/reflect/protoregistry.Types.RangeExtensions].
//
//	RangeExtensions iterates over all registered extensions while f returns true. Iteration order is undefined.
//
// Parameters:
//   - types: A Types implementation providing access to extension types
//
// Returns:
//   - An iterator sequence that yields each extension type
func EachExtension(types Types) iter.Seq[protoreflect.ExtensionType] {
	return func(yield func(protoreflect.ExtensionType) bool) {
		types.RangeExtensions(yield)
	}
}

// EachExtensionByMessage creates a sequential iterator over extension types for a specific message.
//
// It returns an iterator of calling [google.golang.org/protobuf/reflect/protoregistry.Types.RangeExtensionsByMessage].
//
//	RangeExtensionsByMessage iterates over all registered extensions filtered by a given message type while f returns true. Iteration order is undefined.
//
// Parameters:
//   - types: A Types implementation providing access to extension types
//   - message: The full name of the message to filter extension types
//
// Returns:
//   - An iterator sequence that yields extension types for the specified message
func EachExtensionByMessage(types Types, message protoreflect.FullName) iter.Seq[protoreflect.ExtensionType] {
	return func(yield func(protoreflect.ExtensionType) bool) {
		types.RangeExtensionsByMessage(message, yield)
	}
}

// EachField creates a sequential iterator over fields in a protocol buffer message.
//
// It returns an iterator of calling [protoreflect.Message.Range].
//
//	Range iterates over every populated field in an undefined order,
//	calling f for each field descriptor and value encountered.
//	Range returns immediately if f returns false.
//	While iterating, mutating operations may only be performed
//	on the current field descriptor.
//
// Parameters:
//   - message: The protocol buffer message to iterate over
//
// Returns:
//   - An iterator sequence that yields each field descriptor and its corresponding value
func EachField(message protoreflect.Message) iter.Seq2[protoreflect.FieldDescriptor, protoreflect.Value] {
	return func(yield func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
		message.Range(yield)
	}
}
