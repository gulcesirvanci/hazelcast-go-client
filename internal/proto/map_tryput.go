
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
 * Tries to put the given key and value into this map within a specified timeout value. If this method returns false,
 * it means that the caller thread could not acquire the lock for the key within the timeout duration,
 * thus the put operation is not successful.
 */
//@Generated("32031bdaf2989b9068955c157bb2ac16")
const (
    //hex: 0x010C00
    MapTryPutRequestMessageType = 68608
    //hex: 0x010C01
    MapTryPutResponseMessageType = 68609
    MapTryPutRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapTryPutRequestTimeoutFieldOffset = MapTryPutRequestThreadIdFieldOffset + LongSizeInBytes
    MapTryPutRequestInitialFrameSize = MapTryPutRequestTimeoutFieldOffset + LongSizeInBytes
    MapTryPutResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapTryPutResponseInitialFrameSize = MapTryPutResponseResponseFieldOffset + BooleanSizeInBytes

)

func MapTryPutEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64, timeout int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.TryPut")
	initialFrame := &Frame{Content: make([]byte, MapTryPutResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapTryPutRequestMessageType)
    EncodeLong(initialFrame.Content, MapTryPutRequestThreadIdFieldOffset, threadId)
    EncodeLong(initialFrame.Content, MapTryPutRequestTimeoutFieldOffset, timeout)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)


    DataCodecEncode(clientMessage, value)

    return clientMessage
}


func MapTryPutDecodeResponse(clientMessage *ClientMessage) func() (/*** Returns true if successful, otherwise returns false*/response bool) {
    return func() (/*** Returns true if successful, otherwise returns false*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MapTryPutResponseResponseFieldOffset)
        return
    }
}

