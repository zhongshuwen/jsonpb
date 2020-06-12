package jsonpb

import "github.com/golang/protobuf/proto"

func MarshalIndentToString(pb proto.Message, indent string) (string, error) {
	m := &Marshaler{
		Indent: indent,
	}
	return m.MarshalToString(pb)
}
