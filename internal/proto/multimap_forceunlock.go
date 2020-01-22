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
 * Releases the lock for the specified key regardless of the lock owner. It always successfully unlocks the key,
 * never blocks and returns immediately.
 */
//@Generated("3b7fbb688962466706ab8717432510dc")
const (
    //hex: 0x021400
    MultiMapForceUnlockRequestMessageType = 136192
    //hex: 0x021401
    MultiMapForceUnlockResponseMessageType = 136193
    MultiMapForceUnlockRequestReferenceIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapForceUnlockRequestInitialFrameSize = MultiMapForceUnlockRequestReferenceIdFieldOffset + bufutil.LongSizeInBytes
    MultiMapForceUnlockResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapForceUnlockRequestParameters struct {

    /**
    * Name of the MultiMap
    */
name string

    /**
    * The key to Lock
    */
key serialization.Data

    /**
    * The client-wide unique id for this request. It is used to make the request idempotent by sending the same reference id during retries.
    */
referenceId int64
}

func MultiMapForceUnlockEncodeRequest(name string, key serialization.Data, referenceId int64) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = true
    clientMessage.SetOperationName("MultiMap.ForceUnlock")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapForceUnlockRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MultiMapForceUnlockRequestReferenceIdFieldOffset, referenceId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapForceUnlockResponseParameters struct {
}



func MultiMapForceUnlockDecodeResponse(clientMessage *bufutil.ClientMessage) *MultiMapForceUnlockResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MultiMapForceUnlockResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

