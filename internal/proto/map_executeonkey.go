
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
 * Applies the user defined EntryProcessor to the entry mapped by the key. Returns the the object which is result of
 * the process() method of EntryProcessor.
 */
//@Generated("4770e09d1f4f420236b41d3f00c13fbb")
const (
    //hex: 0x012E00
    MapExecuteOnKeyRequestMessageType = 77312
    //hex: 0x012E01
    MapExecuteOnKeyResponseMessageType = 77313
    MapExecuteOnKeyRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapExecuteOnKeyRequestInitialFrameSize = MapExecuteOnKeyRequestThreadIdFieldOffset + LongSizeInBytes
    MapExecuteOnKeyResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapExecuteOnKeyEncodeRequest(name string, entryProcessor serialization.Data, key serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.ExecuteOnKey")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapExecuteOnKeyRequestMessageType)
    EncodeLong(initialFrame.Content, MapExecuteOnKeyRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, entryProcessor)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapExecuteOnKeyDecodeResponse(clientMessage *ClientMessage) func() (/*** result of entry process.*//* @Nullable */response serialization.Data) {
    return func() (/*** result of entry process.*//* @Nullable */response serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        response = DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  
        return
    }
}

