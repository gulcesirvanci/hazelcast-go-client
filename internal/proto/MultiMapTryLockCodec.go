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
 * Tries to acquire the lock for the specified key for the specified lease time. After lease time, the lock will be
 * released. If the lock is not available, then the current thread becomes disabled for thread scheduling purposes
 * and lies dormant until one of two things happens:the lock is acquired by the current thread, or the specified
 * waiting time elapses.
 */
//@Generated("bfd325402a7caa6bc743fb0a8b86c983")
const (
    //hex: 0x021100
    MultiMapTryLockRequestMessageType = 135424
    //hex: 0x021101
    MultiMapTryLockResponseMessageType = 135425
    MultiMapTryLockRequestThreadIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapTryLockRequestLeaseFieldOffset = MultiMapTryLockRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    MultiMapTryLockRequestTimeoutFieldOffset = MultiMapTryLockRequestLeaseFieldOffset + bufutil.LongSizeInBytes
    MultiMapTryLockRequestReferenceIdFieldOffset = MultiMapTryLockRequestTimeoutFieldOffset + bufutil.LongSizeInBytes
    MultiMapTryLockRequestInitialFrameSize = MultiMapTryLockRequestReferenceIdFieldOffset + bufutil.LongSizeInBytes
    MultiMapTryLockResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MultiMapTryLockResponseInitialFrameSize = MultiMapTryLockResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapTryLockRequestParameters struct {

    /**
    * Name of the MultiMap
    */
name string

    /**
    * Key to lock in this map.
    */
key serialization.Data

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation
    */
threadId int64

    /**
    * Time in milliseconds to wait before releasing the lock.
    */
lease int64

    /**
    * Maximum time to wait for the lock.
    */
timeout int64

    /**
    * The client-wide unique id for this request. It is used to make the request idempotent by sending the same reference id during retries.
    */
referenceId int64
}

func MultiMapTryLockEncodeRequest(name string, key serialization.Data, threadId int64, lease int64, timeout int64, referenceId int64) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(true)
    clientMessage.SetOperationName("MultiMap.TryLock")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapTryLockRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, MultiMapTryLockRequestThreadIdFieldOffset, threadId)
    bufutil.EncodeLong(initialFrame.Content, MultiMapTryLockRequestLeaseFieldOffset, lease)
    bufutil.EncodeLong(initialFrame.Content, MultiMapTryLockRequestTimeoutFieldOffset, timeout)
    bufutil.EncodeLong(initialFrame.Content, MultiMapTryLockRequestReferenceIdFieldOffset, referenceId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, key)
    return clientMessage
}

func MultiMapTryLockDecodeRequest(clientMessage *bufutil.ClientMessagex) *MultiMapTryLockRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MultiMapTryLockRequestParameters)
    initialFrame := iterator.Next()
    request.threadId = bufutil.DecodeLong(initialFrame.Content, MultiMapTryLockRequestThreadIdFieldOffset)
    request.lease = bufutil.DecodeLong(initialFrame.Content, MultiMapTryLockRequestLeaseFieldOffset)
    request.timeout = bufutil.DecodeLong(initialFrame.Content, MultiMapTryLockRequestTimeoutFieldOffset)
    request.referenceId = bufutil.DecodeLong(initialFrame.Content, MultiMapTryLockRequestReferenceIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.key = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MultiMapTryLockResponseParameters struct {
    /**
    * True if the lock was acquired and false if the waiting time elapsed before the lock acquired
    */
response bool
}

func MultiMapTryLockEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MultiMapTryLockResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MultiMapTryLockResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, MultiMapTryLockResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func MultiMapTryLockDecodeResponse(clientMessage *bufutil.ClientMessagex) *MultiMapTryLockResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MultiMapTryLockResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, MultiMapTryLockResponseResponseFieldOffset)
    return response
}

