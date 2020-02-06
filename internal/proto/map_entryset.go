
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
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Returns a Set clone of the mappings contained in this map.
 * The collection is NOT backed by the map, so changes to the map are NOT reflected in the collection, and vice-versa.
 * This method is always executed by a distributed query, so it may throw a QueryResultSizeExceededException
 * if query result size limit is configured.
 */
//@Generated("c430b470aa757996581b2bebc7b3119b")
const (
    //hex: 0x012500
    MapEntrySetRequestMessageType = 75008
    //hex: 0x012501
    MapEntrySetResponseMessageType = 75009
    MapEntrySetRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapEntrySetResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapEntrySetEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.EntrySet")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapEntrySetRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapEntrySetDecodeResponse(clientMessage *ClientMessage) func() (/*** a set clone of the keys mappings in this map*/response []*Pair) {
    return func() (/*** a set clone of the keys mappings in this map*/response []*Pair) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        var result []*Pair
        //begin frame, list
        iterator.Next()
        for !NextFrameIsDataStructureEndFrame(iterator) {
            i := 0
            key := DataCodecDecode(iterator)
            value := DataCodecDecode(iterator)
            result[i] = NewPair(key,value)
            i++
        }

        //end frame, list
        iterator.Next()
        response = result //2


        return
    }
}

