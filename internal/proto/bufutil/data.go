package bufutil

import (
	"github.com/hazelcast/hazelcast-go-client/serialization"
	"github.com/hazelcast/hazelcast-go-client/serialization/spi"
)

type DataCodec struct {
}
//TODO
func DataCodecEncode(iterator *ClientMessage, data serialization.Data)  {
	iterator.Add(&Frame{Content:data.(serialization.Data).Buffer()})
}


func DataCodecDecode(iterator *ForwardFrameIterator) serialization.Data {
	return spi.NewData(iterator.Next().Content)
}
