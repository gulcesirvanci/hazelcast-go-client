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
_"bytes"
"github.com/hazelcast/hazelcast-go-client/serialization"
_ "github.com/hazelcast/hazelcast-go-client"
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Increases or decreases the number of permits by the given value.
 */
//@Generated("1f77b9c829ef8987a68b19b86dda17e2")
const (
    //hex: 0x0C0500
    SemaphoreChangeRequestMessageType = 787712
    //hex: 0x0C0501
    SemaphoreChangeResponseMessageType = 787713
    SemaphoreChangeRequestSessionIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    SemaphoreChangeRequestThreadIdFieldOffset = SemaphoreChangeRequestSessionIdFieldOffset + bufutil.LongSizeInBytes
    SemaphoreChangeRequestInvocationUidFieldOffset = SemaphoreChangeRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    SemaphoreChangeRequestPermitsFieldOffset = SemaphoreChangeRequestInvocationUidFieldOffset + bufutil.UUIDSizeInBytes
    SemaphoreChangeRequestInitialFrameSize = SemaphoreChangeRequestPermitsFieldOffset + bufutil.IntSizeInBytes
    SemaphoreChangeResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    SemaphoreChangeResponseInitialFrameSize = SemaphoreChangeResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type SemaphoreChangeRequestParameters struct {

    /**
    * CP group id of this ISemaphore instance
    */
groupId RaftGroupId

    /**
    * Name of this ISemaphore instance
    */
name string

    /**
    * Session ID of the caller
    */
sessionId int64

    /**
    * ID of the caller thread
    */
threadId int64

    /**
    * UID of this invocation
    */
invocationUid UUID

    /**
    * number of permits to increase / decrease
    */
permits int
}

func SemaphoreChangeEncodeRequest(groupId RaftGroupId, name string, sessionId int64, threadId int64, invocationUid UUID, permits int) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Semaphore.Change")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, SemaphoreChangeRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, SemaphoreChangeRequestSessionIdFieldOffset, sessionId)
    bufutil.EncodeLong(initialFrame.Content, SemaphoreChangeRequestThreadIdFieldOffset, threadId)
    bufutil.EncodeUUID(initialFrame.Content, SemaphoreChangeRequestInvocationUidFieldOffset, invocationUid)
    bufutil.EncodeInt(initialFrame.Content, SemaphoreChangeRequestPermitsFieldOffset, permits)
    clientMessage.Add(initialFrame)
    RaftGroupIdCodec.encode(clientMessage, groupId)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func SemaphoreChangeDecodeRequest(clientMessage *bufutil.ClientMessagex) *SemaphoreChangeRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(SemaphoreChangeRequestParameters)
    initialFrame := iterator.Next()
    request.sessionId = bufutil.DecodeLong(initialFrame.Content, SemaphoreChangeRequestSessionIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, SemaphoreChangeRequestThreadIdFieldOffset)
    request.invocationUid = bufutil.DecodeUUID(initialFrame.Content, SemaphoreChangeRequestInvocationUidFieldOffset)
    request.permits = bufutil.DecodeInt(initialFrame.Content, SemaphoreChangeRequestPermitsFieldOffset)
    request.groupId = RaftGroupIdCodec.decode(iterator)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type SemaphoreChangeResponseParameters struct {
    /**
    * true
    */
response bool
}

func SemaphoreChangeEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, SemaphoreChangeResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, SemaphoreChangeResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, SemaphoreChangeResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func SemaphoreChangeDecodeResponse(clientMessage *bufutil.ClientMessagex) *SemaphoreChangeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(SemaphoreChangeResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, SemaphoreChangeResponseResponseFieldOffset)
    return response
}

