
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
 * Returns a set clone of the keys contained in this map. The set is NOT backed by the map, so changes to the map
 * are NOT reflected in the set, and vice-versa. This method is always executed by a distributed query, so it may
 * throw a QueryResultSizeExceededException if query result size limit is configured.
 */
//@Generated("a1bf8573700e7372d189ce0a42f62669")
const (
    //hex: 0x012200
    MapKeySetRequestMessageType = 74240
    //hex: 0x012201
    MapKeySetResponseMessageType = 74241
    MapKeySetRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapKeySetResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapKeySetEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.KeySet")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapKeySetRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapKeySetDecodeResponse(clientMessage *ClientMessage) func() ( /*** a set clone of the keys contained in this map.*/response []serialization.Data) {
    return func() (/*** a set clone of the keys contained in this map.*/response []serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        var result []serialization.Data
        //begin frame, list
        iterator.Next()
        for !NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, DataCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        response = result //0.1
        return
    }
}

