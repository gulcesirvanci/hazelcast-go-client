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
 * Returns the number of elements in this collection.If this collection contains more than Integer.MAX_VALUE
 * elements, returns Integer.MAX_VALUE.
 */
//@Generated("6ff78d219f47e7fe49b20be890bbac8c")
const (
    //hex: 0x120500
    TransactionalQueueSizeRequestMessageType = 1180928
    //hex: 0x120501
    TransactionalQueueSizeResponseMessageType = 1180929
    TransactionalQueueSizeRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalQueueSizeRequestThreadIdFieldOffset = TransactionalQueueSizeRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalQueueSizeRequestInitialFrameSize = TransactionalQueueSizeRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalQueueSizeResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalQueueSizeResponseInitialFrameSize = TransactionalQueueSizeResponseResponseFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalQueueSizeRequestParameters struct {

    /**
    * Name of the Transactional Queue
    */
name string

    /**
    * ID of the transaction
    */
txnId UUID

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64
}

func TransactionalQueueSizeEncodeRequest(name string, txnId UUID, threadId int64) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalQueue.Size")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalQueueSizeRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalQueueSizeRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalQueueSizeRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func TransactionalQueueSizeDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalQueueSizeRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalQueueSizeRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalQueueSizeRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalQueueSizeRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalQueueSizeResponseParameters struct {
    /**
    * The number of elements in this collection
    */
response int
}

func TransactionalQueueSizeEncodeResponse(response int ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalQueueSizeResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalQueueSizeResponseMessageType)
    bufutil.EncodeInt(initialFrame.Content, TransactionalQueueSizeResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func TransactionalQueueSizeDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalQueueSizeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalQueueSizeResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeInt(initialFrame.Content, TransactionalQueueSizeResponseResponseFieldOffset)
    return response
}

