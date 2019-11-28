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
 * Stores a key-value pair in the multimap.
 */
//@Generated("a6a8e29167e8c6c0d5bffc2c732d898b")
const (
    //hex: 0x0F0100
    TransactionalMultiMapPutRequestMessageType = 983296
    //hex: 0x0F0101
    TransactionalMultiMapPutResponseMessageType = 983297
    TransactionalMultiMapPutRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMultiMapPutRequestThreadIdFieldOffset = TransactionalMultiMapPutRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalMultiMapPutRequestInitialFrameSize = TransactionalMultiMapPutRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalMultiMapPutResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMultiMapPutResponseInitialFrameSize = TransactionalMultiMapPutResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMultiMapPutRequestParameters struct {

    /**
    * Name of the Transactional Multi Map
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

    /**
    * The key to be stored
    */
key serialization.Data

    /**
    * The value to be stored
    */
value serialization.Data
}

func TransactionalMultiMapPutEncodeRequest(name string, txnId UUID, threadId int64, key serialization.Data, value serialization.Data) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("TransactionalMultiMap.Put")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMultiMapPutRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalMultiMapPutRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalMultiMapPutRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    DataCodec.encode(clientMessage, key)
    DataCodec.encode(clientMessage, value)
    return clientMessage
}

func TransactionalMultiMapPutDecodeRequest(clientMessage *bufutil.ClientMessagex) *TransactionalMultiMapPutRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(TransactionalMultiMapPutRequestParameters)
    initialFrame := iterator.Next()
    request.txnId = bufutil.DecodeUUID(initialFrame.Content, TransactionalMultiMapPutRequestTxnIdFieldOffset)
    request.threadId = bufutil.DecodeLong(initialFrame.Content, TransactionalMultiMapPutRequestThreadIdFieldOffset)
    request.name = StringCodec.decode(iterator)
    request.key = DataCodec.decode(iterator)
    request.value = DataCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMultiMapPutResponseParameters struct {
    /**
    * True if the size of the multimap is increased, false if the multimap already contains the key-value pair.
    */
response bool
}

func TransactionalMultiMapPutEncodeResponse(response bool ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, TransactionalMultiMapPutResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMultiMapPutResponseMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, TransactionalMultiMapPutResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func TransactionalMultiMapPutDecodeResponse(clientMessage *bufutil.ClientMessagex) *TransactionalMultiMapPutResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalMultiMapPutResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, TransactionalMultiMapPutResponseResponseFieldOffset)
    return response
}

