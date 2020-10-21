package common

import (
  "fmt"
  "reflect"

  "github.com/gogo/protobuf/proto"
)

type Serializer interface {
  Serialize(msg interface{}) ([]byte, error)
  Deserialize(typeName string, bytes []byte) (interface{}, error)
}

type protoSerializer struct{}

func NewSerializer() Serializer {
  return &protoSerializer{}
}

func (protoSerializer) Serialize(msg interface{}) ([]byte, error) {
  if message, ok := msg.(proto.Message); ok {
    bytes, err := proto.Marshal(message)
    if err != nil {
      return nil, err
    }

    return bytes, nil
  }
  return nil, fmt.Errorf("msg must be proto.Message")
}

func (protoSerializer) Deserialize(typeName string, bytes []byte) (interface{}, error) {
  protoType := proto.MessageType(typeName)
  if protoType == nil {
    return nil, fmt.Errorf("unknown message type %v", typeName)
  }
  t := protoType.Elem()

  intPtr := reflect.New(t)
  instance := intPtr.Interface().(proto.Message)
  err := proto.Unmarshal(bytes, instance)
  if err != nil {
    return nil, err
  }

  return instance, nil
}

func (protoSerializer) GetTypeName(msg interface{}) (string, error) {
  if message, ok := msg.(proto.Message); ok {
    typeName := proto.MessageName(message)

    return typeName, nil
  }
  return "", fmt.Errorf("msg must be proto.Message")
}
