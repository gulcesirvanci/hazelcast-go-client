
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
 * Puts an entry into this map with a given ttl (time to live) value.Entry will expire and get evicted after the ttl
 * If ttl is 0, then the entry lives forever.This method returns a clone of the previous value, not the original
 * (identically equal) value previously put into the map.Time resolution for TTL is seconds. The given TTL value is
 * rounded to the next closest second value.
 */
//@Generated("dea3bb96ccdaf8e99b8036950ca3db71")
const (
    //hex: 0x010100
    MapPutRequestMessageType = 65792
    //hex: 0x010101
    MapPutResponseMessageType = 65793
    MapPutRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapPutRequestTtlFieldOffset = MapPutRequestThreadIdFieldOffset + LongSizeInBytes
    MapPutRequestInitialFrameSize = MapPutRequestTtlFieldOffset + LongSizeInBytes
    MapPutResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapPutEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64, ttl int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.Put")
	initialFrame := &Frame{Content: make([]byte, MapPutResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapPutRequestMessageType)
    EncodeLong(initialFrame.Content, MapPutRequestThreadIdFieldOffset, threadId)
    EncodeLong(initialFrame.Content, MapPutRequestTtlFieldOffset, ttl)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func MapPutDecodeResponse(clientMessage *ClientMessage) func() (/*** old value of the entry*//* @Nullable */response serialization.Data) {
    return func() (/*** old value of the entry*//* @Nullable */response serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        response = DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  
        return
    }
}

