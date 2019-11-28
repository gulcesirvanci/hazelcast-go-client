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
 * Associates the specified value with the specified key in this map. If the map previously contained a mapping for
 * the key, the old value is replaced by the specified value. This method is preferred to #put(Object, Object)
 * if the old value is not needed.
 * The object to be set will be accessible only in the current transaction context until the transaction is committed.
 */
//@Generated("da0f83e2391e9a2429a39e86b52c6a1f")
const (
    //hex: 0x0E0700
    TransactionalMapSetRequestMessageType = 919296
    //hex: 0x0E0701
    TransactionalMapSetResponseMessageType = 919297
    TransactionalMapSetRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMapSetRequestThreadIdFieldOffset = TransactionalMapSetRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalMapSetRequestInitialFrameSize = TransactionalMapSetRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalMapSetResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapSetRequestParameters struct {

    /**
    * Name of the Transactional Map
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
    * The specified key
    */
key serialization.Data

    /**
    * The value to associate with key
    */
value serialization.Data
}

func TransactionalMapSetEncodeRequest(name string, txnId UUID, threadId int64, key serialization.Data, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalMap.Set")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapSetRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalMapSetRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalMapSetRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, key)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func TransactionalMapSetDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalMapSetRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalMapSetRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalMapSetRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalMapSetRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.key = DataCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMapSetResponseParameters struct {
}

func TransactionalMapSetEncodeResponse() *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalMapSetResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMapSetResponseMessageType)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func TransactionalMapSetDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalMapSetResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalMapSetResponseParameters)
    //empty initial frame
    iterator.Next()
    return response
}

