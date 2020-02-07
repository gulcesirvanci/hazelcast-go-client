
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
 * Removes the mapping for a key from this map if it is present.Unlike remove(Object), this operation does not return
 * the removed value, which avoids the serialization cost of the returned value.If the removed value will not be used,
 * a delete operation is preferred over a remove operation for better performance. The map will not contain a mapping
 * for the specified key once the call returns.
 * This method breaks the contract of EntryListener. When an entry is removed by delete(), it fires an EntryEvent
 * with a null oldValue. Also, a listener with predicates will have null values, so only keys can be queried via predicates
 */
//@Generated("b99b965276ddaec789b4fef736603054")
const (
    //hex: 0x010900
    MapDeleteRequestMessageType = 67840
    //hex: 0x010901
    MapDeleteResponseMessageType = 67841
    MapDeleteRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapDeleteRequestInitialFrameSize = MapDeleteRequestThreadIdFieldOffset + LongSizeInBytes
    MapDeleteResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapDeleteEncodeRequest(name string, key serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Delete")
	initialFrame := &Frame{Content: make([]byte, MapDeleteResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapDeleteRequestMessageType)
    EncodeLong(initialFrame.Content, MapDeleteRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapDeleteDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}

