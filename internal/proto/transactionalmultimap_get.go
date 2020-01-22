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
 * Returns the collection of values associated with the key.
 */
//@Generated("b1ba93a3e077327daa4e4608efdd71e5")
const (
    //hex: 0x0F0200
    TransactionalMultiMapGetRequestMessageType = 983552
    //hex: 0x0F0201
    TransactionalMultiMapGetResponseMessageType = 983553
    TransactionalMultiMapGetRequestTxnIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    TransactionalMultiMapGetRequestThreadIdFieldOffset = TransactionalMultiMapGetRequestTxnIdFieldOffset + bufutil.UUIDSizeInBytes
    TransactionalMultiMapGetRequestInitialFrameSize = TransactionalMultiMapGetRequestThreadIdFieldOffset + bufutil.LongSizeInBytes
    TransactionalMultiMapGetResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMultiMapGetRequestParameters struct {

    /**
    * Name of the Transactional Multi Map
    */
name string

    /**
    * ID of the transaction
    */
txnId string

    /**
    * The id of the user thread performing the operation. It is used to guarantee that only the lock holder thread (if a lock exists on the entry) can perform the requested operation.
    */
threadId int64

    /**
    * The key whose associated values are returned
    */
key serialization.Data
}

func TransactionalMultiMapGetEncodeRequest(name string, txnId string, threadId int64, key serialization.Data) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("TransactionalMultiMap.Get")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, TransactionalMultiMapGetRequestMessageType)
    bufutil.EncodeUUID(initialFrame.Content, TransactionalMultiMapGetRequestTxnIdFieldOffset, txnId)
    bufutil.EncodeLong(initialFrame.Content, TransactionalMultiMapGetRequestThreadIdFieldOffset, threadId)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.DataCodecEncode(clientMessage, key)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type TransactionalMultiMapGetResponseParameters struct {
    /**
    * The collection of the values associated with the key
    */
response []serialization.Data
}



func TransactionalMultiMapGetDecodeResponse(clientMessage *bufutil.ClientMessage) *TransactionalMultiMapGetResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(TransactionalMultiMapGetResponseParameters)
    //empty initial frame
    iterator.Next()
    var result []serialization.Data
        //begin frame, list
        iterator.Next()
        for !bufutil.NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, bufutil.DataCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        response.response = result
    return response
}

