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
 * Add new item to transactional set.
 */
//@Generated("ed27dba8aeaa102be8091d081d9f5872")
const (
    //hex: 0x100100
    TransactionalSetAddRequestMessageType = 1048832
    //hex: 0x100101
    TransactionalSetAddResponseMessageType = 1048833
    TransactionalSetAddRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalSetAddRequestThreadIdFieldOffset = TransactionalSetAddRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalSetAddRequestInitialFrameSize = TransactionalSetAddRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalSetAddResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalSetAddResponseInitialFrameSize = TransactionalSetAddResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalSetAddRequestParameters struct {

    /**
    * Name of the Transactional Set
    */
name string

    /**
    * ID of the this transaction operation
    */
txnId UUID

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64

    /**
    * Item added to transactional set
    */
item serialization.Data
}

func TransactionalSetAddEncodeRequest(name string, txnId UUID, threadId int64, item serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalSet.Add")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalSetAddRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalSetAddRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalSetAddRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, item)
    return clientMessage
}

func TransactionalSetAddDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalSetAddRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalSetAddRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalSetAddRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalSetAddRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.item = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalSetAddResponseParameters struct {
    /**
    * True if item is added successfully
    */
response bool
}

func TransactionalSetAddEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalSetAddResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalSetAddResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, TransactionalSetAddResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func TransactionalSetAddDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalSetAddResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalSetAddResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, TransactionalSetAddResponseResponseFieldOffset)
    return response
}

