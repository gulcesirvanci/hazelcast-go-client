
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
 * Checks the lock for the specified key.If the lock is acquired then returns true, else returns false.
 */
//@Generated("1e053870b9cad0017b3cf81ce41ccbdb")
const (
    //hex: 0x011200
    MapIsLockedRequestMessageType = 70144
    //hex: 0x011201
    MapIsLockedResponseMessageType = 70145
    MapIsLockedRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapIsLockedResponseResponseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    MapIsLockedResponseInitialFrameSize = MapIsLockedResponseResponseFieldOffset + BooleanSizeInBytes

)

func MapIsLockedEncodeRequest(name string, key serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.IsLocked")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapIsLockedRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, key)

    return clientMessage
}


func MapIsLockedDecodeResponse(clientMessage *ClientMessage) func() ( /*** Returns true if the entry is locked, otherwise returns false*/response bool) {
    return func() (/*** Returns true if the entry is locked, otherwise returns false*/response bool) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        response = DecodeBoolean(initialFrame.Content, MapIsLockedResponseResponseFieldOffset)
        return
    }
}

