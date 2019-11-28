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
 * Removes at most the given number of available elements from this queue and adds them to the given collection.
 * A failure encountered while attempting to add elements to collection may result in elements being in neither,
 * either or both collections when the associated exception is thrown. Attempts to drain a queue to itself result in
 * ILLEGAL_ARGUMENT. Further, the behavior of this operation is undefined if the specified collection is
 * modified while the operation is in progress.
 */
//@Generated("ca21b5aee8f0505d79b28992e993b13f")
const (
    //hex: 0x030A00
    QueueDrainToMaxSizeRequestMessageType = 199168
    //hex: 0x030A01
    QueueDrainToMaxSizeResponseMessageType = 199169
    QueueDrainToMaxSizeRequestMaxSizeFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    QueueDrainToMaxSizeRequestInitialFrameSize = QueueDrainToMaxSizeRequestMaxSizeFieldOffset + bufutil.IntSizeInBytes
    QueueDrainToMaxSizeResponseInitialFrameSize = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueDrainToMaxSizeRequestParameters struct {

    /**
    * Name of the Queue
    */
name string

    /**
    * The maximum number of elements to transfer
    */
maxSize int
}

func QueueDrainToMaxSizeEncodeRequest(name string, maxSize int) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Queue.DrainToMaxSize")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueDrainToMaxSizeRequestMessageType)
    bufutil.EncodeInt(initialFrame.Content, QueueDrainToMaxSizeRequestMaxSizeFieldOffset, maxSize)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func QueueDrainToMaxSizeDecodeRequest(clientMessage *bufutil.ClientMessagex) *QueueDrainToMaxSizeRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(QueueDrainToMaxSizeRequestParameters)
    initialFrame := iterator.Next()
    request.maxSize = bufutil.DecodeInt(initialFrame.Content, QueueDrainToMaxSizeRequestMaxSizeFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueDrainToMaxSizeResponseParameters struct {
    /**
    * list of all removed data in result of this method
    */
response []serialization.Data
}

func QueueDrainToMaxSizeEncodeResponse(response []serialization.Data ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, QueueDrainToMaxSizeResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueDrainToMaxSizeResponseMessageType)
    clientMessage.Add(initialFrame)

    ListMultiFrameCodec.encode(clientMessage, response, DataCodecEncode)
    return clientMessage
}

func QueueDrainToMaxSizeDecodeResponse(clientMessage *bufutil.ClientMessagex) *QueueDrainToMaxSizeResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(QueueDrainToMaxSizeResponseParameters)
    //empty initial frame
    iterator.Next()
    response.response = ListMultiFrameCodec.decode(iterator, DataCodecDecode)
    return response
}

