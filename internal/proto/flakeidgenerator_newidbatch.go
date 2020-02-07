
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
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Fetches a new batch of ids for the given flake id generator.
 */
//@Generated("c49f852903aad8ae579f0ce69358d780")
const (
    //hex: 0x1C0100
    FlakeIdGeneratorNewIdBatchRequestMessageType = 1835264
    //hex: 0x1C0101
    FlakeIdGeneratorNewIdBatchResponseMessageType = 1835265
    FlakeIdGeneratorNewIdBatchRequestBatchSizeFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    FlakeIdGeneratorNewIdBatchRequestInitialFrameSize = FlakeIdGeneratorNewIdBatchRequestBatchSizeFieldOffset + IntSizeInBytes
    FlakeIdGeneratorNewIdBatchResponseBaseFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    FlakeIdGeneratorNewIdBatchResponseIncrementFieldOffset = FlakeIdGeneratorNewIdBatchResponseBaseFieldOffset + LongSizeInBytes
    FlakeIdGeneratorNewIdBatchResponseBatchSizeFieldOffset = FlakeIdGeneratorNewIdBatchResponseIncrementFieldOffset + LongSizeInBytes
    FlakeIdGeneratorNewIdBatchResponseInitialFrameSize = FlakeIdGeneratorNewIdBatchResponseBatchSizeFieldOffset + IntSizeInBytes

)

func FlakeIdGeneratorNewIdBatchEncodeRequest(name string, batchSize int32) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("FlakeIdGenerator.NewIdBatch")
	initialFrame := &Frame{Content: make([]byte, FlakeIdGeneratorNewIdBatchResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, FlakeIdGeneratorNewIdBatchRequestMessageType)
    EncodeInt(initialFrame.Content, FlakeIdGeneratorNewIdBatchRequestBatchSizeFieldOffset, batchSize)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)

    return clientMessage
}


func FlakeIdGeneratorNewIdBatchDecodeResponse(clientMessage *ClientMessage) func() (/*** First id in the batch.*/base int64, /*** Increment for the next id in the batch.*/increment int64, /*** Number of ids in the batch.*/batchSize int32) {
    return func() (/*** First id in the batch.*/base int64, /*** Increment for the next id in the batch.*/increment int64, /*** Number of ids in the batch.*/batchSize int32) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        base = DecodeLong(initialFrame.Content, FlakeIdGeneratorNewIdBatchResponseBaseFieldOffset)
        increment = DecodeLong(initialFrame.Content, FlakeIdGeneratorNewIdBatchResponseIncrementFieldOffset)
        batchSize = DecodeInt(initialFrame.Content, FlakeIdGeneratorNewIdBatchResponseBatchSizeFieldOffset)
        return
    }
}

