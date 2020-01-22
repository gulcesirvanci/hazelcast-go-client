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
 * Removes all of this collection's elements that are also contained in the specified collection (optional operation).
 * After this call returns, this collection will contain no elements in common with the specified collection.
 */
//@Generated("7ef3299a6237655813b8e65634d8f4b8")
const (
    //hex: 0x030D00
    QueueCompareAndRemoveAllRequestMessageType = 199936
    //hex: 0x030D01
    QueueCompareAndRemoveAllResponseMessageType = 199937
    QueueCompareAndRemoveAllRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    QueueCompareAndRemoveAllResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    QueueCompareAndRemoveAllResponseInitialFrameSize = QueueCompareAndRemoveAllResponseResponseFieldOffset + bufutil.BooleanSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueCompareAndRemoveAllRequestParameters struct {

    /**
    * Name of the Queue
    */
name string

    /**
    * Collection containing elements to be removed from this collection
    */
dataList []serialization.Data
}

func QueueCompareAndRemoveAllEncodeRequest(name string, dataList []serialization.Data) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("Queue.CompareAndRemoveAll")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, QueueCompareAndRemoveAllRequestMessageType)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    clientMessage.Add(bufutil.BeginFrame)
        for i := 0; i < len(dataList) ; i++ {
            bufutil.DataCodecEncode(clientMessage, dataList[i])
        }
        clientMessage.Add(bufutil.EndFrame)


    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type QueueCompareAndRemoveAllResponseParameters struct {
    /**
    * <tt>true</tt> if this collection changed as a result of the call
    */
response bool
}



func QueueCompareAndRemoveAllDecodeResponse(clientMessage *bufutil.ClientMessage) *QueueCompareAndRemoveAllResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(QueueCompareAndRemoveAllResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeBoolean(initialFrame.Content, QueueCompareAndRemoveAllResponseResponseFieldOffset)
    return response
}

