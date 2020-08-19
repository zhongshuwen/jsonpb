package jsonpb

import "github.com/golang/protobuf/proto"

func MarshalToString(pb proto.Message) (string, error) {
	m := &Marshaler{}
	return m.MarshalToString(pb)
}

func MarshalIndentToString(pb proto.Message, indent string) (string, error) {
	m := &Marshaler{
		Indent: indent,
	}
	return m.MarshalToString(pb)
}
