
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
 * Associates a given value to the specified key and replicates it to the cluster. If there is an old value, it will
 * be replaced by the specified one and returned from the call. In addition, you have to specify a ttl and its TimeUnit
 * to define when the value is outdated and thus should be removed from the replicated map.
 */
//@Generated("e65e2b159c8e4e467bc0c2859228ffa9")
const (
    //hex: 0x0D0100
    ReplicatedMapPutRequestMessageType = 852224
    //hex: 0x0D0101
    ReplicatedMapPutResponseMessageType = 852225
    ReplicatedMapPutRequestTtlFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    ReplicatedMapPutRequestInitialFrameSize = ReplicatedMapPutRequestTtlFieldOffset + LongSizeInBytes
    ReplicatedMapPutResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func ReplicatedMapPutEncodeRequest(name string, key serialization.Data, value serialization.Data, ttl int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("ReplicatedMap.Put")
	initialFrame := &Frame{Content: make([]byte, ReplicatedMapPutResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, ReplicatedMapPutRequestMessageType)
    EncodeLong(initialFrame.Content, ReplicatedMapPutRequestTtlFieldOffset, ttl)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func ReplicatedMapPutDecodeResponse(clientMessage *ClientMessage) func() (/*** The old value if existed for the key.*//* @Nullable */response serialization.Data) {
    return func() (/*** The old value if existed for the key.*//* @Nullable */response serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        response = DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  
        return
    }
}

