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
 * Remove item from transactional set.
 */
//@Generated("03e2a9e9317836768ebe7e6578fcab1c")
const (
    //hex: 0x100200
    TransactionalSetRemoveRequestMessageType = 1049088
    //hex: 0x100201
    TransactionalSetRemoveResponseMessageType = 1049089
    TransactionalSetRemoveRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalSetRemoveRequestThreadIdFieldOffset = TransactionalSetRemoveRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalSetRemoveRequestInitialFrameSize = TransactionalSetRemoveRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalSetRemoveResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalSetRemoveResponseInitialFrameSize = TransactionalSetRemoveResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalSetRemoveRequestParameters struct {

    /**
    * Name of the Transactional Set
    */
name string

    /**
    * ID of the this transaction operation
    */
txnId string

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64

    /**
    * Item removed from Transactional Set
    */
item serialization.Data
}

func TransactionalSetRemoveEncodeRequest(name string, txnId string, threadId int64, item serialization.Data) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("TransactionalSet.Remove")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalSetRemoveRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalSetRemoveRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalSetRemoveRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, item)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalSetRemoveResponseParameters struct {
    /**
    * True if item is remove successfully
    */
response bool
}



func TransactionalSetRemoveDecodeResponse(clientMessage *bufutil.ClientMessage) *TransactionalSetRemoveResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalSetRemoveResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, TransactionalSetRemoveResponseResponseFieldOffset)
    return response
}

