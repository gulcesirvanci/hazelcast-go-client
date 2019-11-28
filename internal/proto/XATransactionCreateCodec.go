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
 * TODO DOC
 */
//@Generated("9c25e435e7b4d62e2e34a89939796bd5")
const (
    //hex: 0x140500
    XATransactionCreateRequestMessageType = 1312000
    //hex: 0x140501
    XATransactionCreateResponseMessageType = 1312001
    XATransactionCreateRequestTimeoutFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    XATransactionCreateRequestInitialFrameSize = XATransactionCreateRequestTimeoutFieldOffset + bufutil.LongSizeInBytes
    XATransactionCreateResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    XATransactionCreateResponseInitialFrameSize = XATransactionCreateResponseResponseFieldOffset + bufutil.UUIDSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type XATransactionCreateRequestParameters struct {

    /**
    * Java XA transaction id as defined in interface javax.transaction.xa.Xid.
    */
xid SerializableXID

    /**
    * The timeout in seconds for XA operations such as prepare, commit, rollback.
    */
timeout int64
}

func XATransactionCreateEncodeRequest(xid Xid, timeout int64) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(true)
    clientMessage.SetOperationName("XATransaction.Create")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, XATransactionCreateRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, XATransactionCreateRequestTimeoutFieldOffset, timeout)
    clientMessage.Add(initialFrame)
    XidCodec.encode(clientMessage, xid)
    return clientMessage
}

func XATransactionCreateDecodeRequest(clientMessage *bufutil.ClientMessagex) *XATransactionCreateRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(XATransactionCreateRequestParameters)
    initialFrame := iterator.Next()
    request.timeout = bufutil.DecodeLong(initialFrame.Content, XATransactionCreateRequestTimeoutFieldOffset)
    request.xid = XidCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type XATransactionCreateResponseParameters struct {
    /**
    * The transaction unique identifier.
    */
response UUID
}

func XATransactionCreateEncodeResponse(response UUID ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, XATransactionCreateResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, XATransactionCreateResponseMessageType)
    bufutil.EncodeUUID(initialFrame.Content, XATransactionCreateResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func XATransactionCreateDecodeResponse(clientMessage *bufutil.ClientMessagex) *XATransactionCreateResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(XATransactionCreateResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeUUID(initialFrame.Content, XATransactionCreateResponseResponseFieldOffset)
    return response
}

