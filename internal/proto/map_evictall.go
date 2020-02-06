
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
 * Evicts all keys from this map except the locked ones. If a MapStore is defined for this map, deleteAll is not
 * called by this method. If you do want to deleteAll to be called use the clear method. The EVICT_ALL event is
 * fired for any registered listeners.
 */
//@Generated("8fb60d3b4004571352bf1101420c62a0")
const (
    //hex: 0x011F00
    MapEvictAllRequestMessageType = 73472
    //hex: 0x011F01
    MapEvictAllResponseMessageType = 73473
    MapEvictAllRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapEvictAllResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapEvictAllEncodeRequest(name string) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.EvictAll")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapEvictAllRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func MapEvictAllDecodeResponse(clientMessage *ClientMessage) func() () {
    return func() () {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        return
    }
}

