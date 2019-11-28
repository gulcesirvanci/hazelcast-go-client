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
 * Returns statistics of the task
 */
//@Generated("4b4a54ba045747e379231e918d317c14")
const (
    //hex: 0x1A0600
    ScheduledExecutorGetStatsFromAddressRequestMessageType = 1705472
    //hex: 0x1A0601
    ScheduledExecutorGetStatsFromAddressResponseMessageType = 1705473
    ScheduledExecutorGetStatsFromAddressRequestInitialFrameSize = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    ScheduledExecutorGetStatsFromAddressResponseLastIdleTimeNanosFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    ScheduledExecutorGetStatsFromAddressResponseTotalIdleTimeNanosFieldOffset = ScheduledExecutorGetStatsFromAddressResponseLastIdleTimeNanosFieldOffset + bufutil.LongSizeInBytes
    ScheduledExecutorGetStatsFromAddressResponseTotalRunsFieldOffset = ScheduledExecutorGetStatsFromAddressResponseTotalIdleTimeNanosFieldOffset + bufutil.LongSizeInBytes
    ScheduledExecutorGetStatsFromAddressResponseTotalRunTimeNanosFieldOffset = ScheduledExecutorGetStatsFromAddressResponseTotalRunsFieldOffset + bufutil.LongSizeInBytes
    ScheduledExecutorGetStatsFromAddressResponseLastRunDurationNanosFieldOffset = ScheduledExecutorGetStatsFromAddressResponseTotalRunTimeNanosFieldOffset + bufutil.LongSizeInBytes
    ScheduledExecutorGetStatsFromAddressResponseInitialFrameSize = ScheduledExecutorGetStatsFromAddressResponseLastRunDurationNanosFieldOffset + bufutil.LongSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ScheduledExecutorGetStatsFromAddressRequestParameters struct {

    /**
    * The name of the scheduler.
    */
schedulerName string

    /**
    * The name of the task
    */
taskName string

    /**
    * The address of the member where the task will get scheduled.
    */
address Address
}

func ScheduledExecutorGetStatsFromAddressEncodeRequest(schedulerName string, taskName string, address Address) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = true
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("ScheduledExecutor.GetStatsFromAddress")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ScheduledExecutorGetStatsFromAddressRequestMessageType)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, schedulerName)
    StringCodec.encode(clientMessage, taskName)
    AddressCodec.encode(clientMessage, address)
    return clientMessage
}

func ScheduledExecutorGetStatsFromAddressDecodeRequest(clientMessage *bufutil.ClientMessagex) *ScheduledExecutorGetStatsFromAddressRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(ScheduledExecutorGetStatsFromAddressRequestParameters)
    //empty initial frame
    iterator.Next()
    request.schedulerName = StringCodec.decode(iterator)
    request.taskName = StringCodec.decode(iterator)
    request.address = AddressCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type ScheduledExecutorGetStatsFromAddressResponseParameters struct {
    /**
    * TODO DOC
    */
lastIdleTimeNanos int64
    /**
    * TODO DOC
    */
totalIdleTimeNanos int64
    /**
    * TODO DOC
    */
totalRuns int64
    /**
    * TODO DOC
    */
totalRunTimeNanos int64
    /**
    * TODO DOC
    */
lastRunDurationNanos int64
}

func ScheduledExecutorGetStatsFromAddressEncodeResponse(lastIdleTimeNanos int64 , totalIdleTimeNanos int64 , totalRuns int64 , totalRunTimeNanos int64 , lastRunDurationNanos int64 ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, ScheduledExecutorGetStatsFromAddressResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, ScheduledExecutorGetStatsFromAddressResponseMessageType)
    bufutil.EncodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseLastIdleTimeNanosFieldOffset, lastIdleTimeNanos)
    bufutil.EncodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseTotalIdleTimeNanosFieldOffset, totalIdleTimeNanos)
    bufutil.EncodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseTotalRunsFieldOffset, totalRuns)
    bufutil.EncodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseTotalRunTimeNanosFieldOffset, totalRunTimeNanos)
    bufutil.EncodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseLastRunDurationNanosFieldOffset, lastRunDurationNanos)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func ScheduledExecutorGetStatsFromAddressDecodeResponse(clientMessage *bufutil.ClientMessagex) *ScheduledExecutorGetStatsFromAddressResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(ScheduledExecutorGetStatsFromAddressResponseParameters)
    initialFrame := iterator.Next()
    response.lastIdleTimeNanos = bufutil.DecodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseLastIdleTimeNanosFieldOffset)
    response.totalIdleTimeNanos = bufutil.DecodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseTotalIdleTimeNanosFieldOffset)
    response.totalRuns = bufutil.DecodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseTotalRunsFieldOffset)
    response.totalRunTimeNanos = bufutil.DecodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseTotalRunTimeNanosFieldOffset)
    response.lastRunDurationNanos = bufutil.DecodeLong(initialFrame.Content, ScheduledExecutorGetStatsFromAddressResponseLastRunDurationNanosFieldOffset)
    return response
}

