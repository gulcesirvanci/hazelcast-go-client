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
 * Returns whether the multimap contains the given key-value pair.
 */
//@Generated("644fd09cca4ad19bacd6b13cd1a9092a")
const (
    //hex: 0x020900
    MultiMapContainsEntryRequestMessageType = 133376
    //hex: 0x020901
    MultiMapContainsEntryResponseMessageType = 133377
    MultiMapContainsEntryRequestThreadIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapContainsEntryRequestInitialFrameSize = MultiMapContainsEntryRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    MultiMapContainsEntryResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapContainsEntryResponseInitialFrameSize = MultiMapContainsEntryResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapContainsEntryRequestParameters struct {

    /**
    * Name of the MultiMap
    */
name string

    /**
    * The key whose existence is checked.
    */
key serialization.Data

    /**
    * The value whose existence is checked.
    */
value serialization.Data

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation
    */
threadId int64
}

func MultiMapContainsEntryEncodeRequest(name string, key serialization.Data, value serialization.Data, threadId int64) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetOperationName("MultiMap.ContainsEntry")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapContainsEntryRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MultiMapContainsEntryRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    bufutil.DataCodecEncode(clientMessage, value)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapContainsEntryResponseParameters struct {
    /**
    * True if the multimap contains the key-value pair, false otherwise.
    */
response bool
}



func MultiMapContainsEntryDecodeResponse(clientMessage *bufutil.ClientMessage) *MultiMapContainsEntryResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MultiMapContainsEntryResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, MultiMapContainsEntryResponseResponseFieldOffset)
    return response
}

