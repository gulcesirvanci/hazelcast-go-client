
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
 * Applies the user defined EntryProcessor to the entry mapped by the key. Returns immediately with a Future
 * representing that task.EntryProcessor is not cancellable, so calling Future.cancel() method won't cancel the
 * operation of EntryProcessor.
 */
//@Generated("80e0e88a8a4b0aacdc60f1d8400d242c")
const (
    //hex: 0x012F00
    MapSubmitToKeyRequestMessageType = 77568
    //hex: 0x012F01
    MapSubmitToKeyResponseMessageType = 77569
    MapSubmitToKeyRequestThreadIdFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    MapSubmitToKeyRequestInitialFrameSize = MapSubmitToKeyRequestThreadIdFieldOffset + LongSizeInBytes
    MapSubmitToKeyResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapSubmitToKeyEncodeRequest(name string, entryProcessor serialization.Data, key serialization.Data, threadId int64) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( false )
    clientMessage.SetOperationName("Map.SubmitToKey")
	initialFrame := &Frame{Content: make([]byte, MapSubmitToKeyResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapSubmitToKeyRequestMessageType)
    EncodeLong(initialFrame.Content, MapSubmitToKeyRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, entryProcessor)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapSubmitToKeyDecodeResponse(clientMessage *ClientMessage) func() (/*** result of entry process.*//* @Nullable */response serialization.Data) {
    return func() (/*** result of entry process.*//* @Nullable */response serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        response = DecodeNullable(iterator, DataCodecDecode).(serialization.Data) // 2  
        return
    }
}

