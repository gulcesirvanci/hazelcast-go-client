
/*
* Copyright (c) 2008-2019, Hazelcast, Inc. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License")
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package proto



import (
     "github.com/hazelcast/hazelcast-go-client/serialization"
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Copies all of the mappings from the specified map to this map (optional operation).The effect of this call is
 * equivalent to that of calling put(Object,Object) put(k, v) on this map once for each mapping from key k to value
 * v in the specified map.The behavior of this operation is undefined if the specified map is modified while the
 * operation is in progress.
 * Please note that all the keys in the request should belong to the partition id to which this request is being sent, all keys
 * matching to a different partition id shall be ignored. The API implementation using this request may need to send multiple
 * of these request messages for filling a request for a key set if the keys belong to different partitions.
 */
//@Generated("7955c3d12c5761bfeb3783097604dcd4")
const (
    //hex: 0x012C00
    MapPutAllRequestMessageType = 76800
    //hex: 0x012C01
    MapPutAllResponseMessageType = 76801
    MapPutAllRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapPutAllResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapPutAllEncodeRequest(name string, entries []*Pair) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.PutAll")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapPutAllRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


            clientMessage.Add(BeginFrame.Copy())

        for _, cp := range entries{
           DataCodecEncode(clientMessage, (*cp).Key().(serialization.Data))
           DataCodecEncode(clientMessage, (*cp).Value().(serialization.Data))
        }

        clientMessage.Add(EndFrame.Copy())



    return clientMessage
}


func MapPutAllDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}

