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
    "github.com/hazelcast/hazelcast-go-client/internal/proto/bufutil"
    "github.com/hazelcast/hazelcast-go-client/serialization"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Removes all the entries with the given key. The collection is NOT backed by the map, so changes to the map are
 * NOT reflected in the collection, and vice-versa.
 */
//@Generated("816cf10acf31d96ecc0ee47e45b96bc7")
const (
    //hex: 0x021500
    MultiMapRemoveEntryRequestMessageType = 136448
    //hex: 0x021501
    MultiMapRemoveEntryResponseMessageType = 136449
    MultiMapRemoveEntryRequestThreadIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapRemoveEntryRequestInitialFrameSize = MultiMapRemoveEntryRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    MultiMapRemoveEntryResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapRemoveEntryResponseInitialFrameSize = MultiMapRemoveEntryResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapRemoveEntryRequestParameters struct {

    /**
    * Name of the MultiMap
    */
name string

    /**
    * The key of the entry to remove
    */
key serialization.Data

    /**
    * The value of the entry to remove
    */
value serialization.Data

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation
    */
threadId int64
}

func MultiMapRemoveEntryEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("MultiMap.RemoveEntry")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapRemoveEntryRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MultiMapRemoveEntryRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    bufutil.DataCodecEncode(clientMessage, value)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapRemoveEntryResponseParameters struct {
    /**
    * True if the size of the multimap changed after the remove operation, false otherwise.
    */
response bool
}



func MultiMapRemoveEntryDecodeResponse(clientMessage *bufutil.ClientMessage) *MultiMapRemoveEntryResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MultiMapRemoveEntryResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, MultiMapRemoveEntryResponseResponseFieldOffset)
    return response
}

