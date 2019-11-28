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
 * Adds a new item to the transactional list.
 */
//@Generated("2b88a70714be8922ae7237f9eac30b0e")
const (
    //hex: 0x110100
    TransactionalListAddRequestMessageType = 1114368
    //hex: 0x110101
    TransactionalListAddResponseMessageType = 1114369
    TransactionalListAddRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalListAddRequestThreadIdFieldOffset = TransactionalListAddRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalListAddRequestInitialFrameSize = TransactionalListAddRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalListAddResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalListAddResponseInitialFrameSize = TransactionalListAddResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalListAddRequestParameters struct {

    /**
    * Name of the Transactional List
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
    * The new item added to the transactionalList
    */
item serialization.Data
}

func TransactionalListAddEncodeRequest(name string, txnId UUID, threadId int64, item serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalList.Add")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalListAddRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalListAddRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalListAddRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, item)
    return clientMessage
}

func TransactionalListAddDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalListAddRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalListAddRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalListAddRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalListAddRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.item = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalListAddResponseParameters struct {
    /**
    * True if the item is added successfully, false otherwise
    */
response bool
}

func TransactionalListAddEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalListAddResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalListAddResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, TransactionalListAddResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func TransactionalListAddDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalListAddResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalListAddResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, TransactionalListAddResponseResponseFieldOffset)
    return response
}

