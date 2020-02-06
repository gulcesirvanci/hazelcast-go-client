
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
 * Returns the set of keys in the multimap.The collection is NOT backed by the map, so changes to the map are NOT
 * reflected in the collection, and vice-versa.
 */
//@Generated("01d48880f2c6ae9508d32fde448bbebb")
const (
    //hex: 0x020400
    MultiMapKeySetRequestMessageType = 132096
    //hex: 0x020401
    MultiMapKeySetResponseMessageType = 132097
    MultiMapKeySetRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MultiMapKeySetResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MultiMapKeySetEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("MultiMap.KeySet")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MultiMapKeySetRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MultiMapKeySetDecodeResponse(clientMessage *ClientMessage) func() ( /*** The set of keys in the multimap. The returned set might be modifiable but it has no effect on the multimap.*/response []serialization.Data) {
    return func() (/*** The set of keys in the multimap. The returned set might be modifiable but it has no effect on the multimap.*/response []serialization.Data) {
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

