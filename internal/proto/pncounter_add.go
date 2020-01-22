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
)

/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Adds a delta to the PNCounter value. The delta may be negative for a
 * subtraction.
 * <p>
 * The invocation will return the replica timestamps (vector clock) which
 * can then be sent with the next invocation to keep session consistency
 * guarantees.
 * The target replica is determined by the {@code targetReplica} parameter.
 * If smart routing is disabled, the actual member processing the client
 * message may act as a proxy.
 */
//@Generated("9747aab0727b6b44accf767b9d48a3bb")
const (
    //hex: 0x1D0200
    PNCounterAddRequestMessageType = 1901056
    //hex: 0x1D0201
    PNCounterAddResponseMessageType = 1901057
    PNCounterAddRequestDeltaFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    PNCounterAddRequestGetBeforeUpdateFieldOffset = PNCounterAddRequestDeltaFieldOffset + bufutil.LongSizeInBytes
    PNCounterAddRequestTargetReplicaUUIDFieldOffset = PNCounterAddRequestGetBeforeUpdateFieldOffset + bufutil.BooleanSizeInBytes
    PNCounterAddRequestInitialFrameSize = PNCounterAddRequestTargetReplicaUUIDFieldOffset + bufutil.UUIDSizeInBytes
    PNCounterAddResponseValueFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    PNCounterAddResponseReplicaCountFieldOffset = PNCounterAddResponseValueFieldOffset + bufutil.LongSizeInBytes
    PNCounterAddResponseInitialFrameSize = PNCounterAddResponseReplicaCountFieldOffset + bufutil.IntSizeInBytes

)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type PNCounterAddRequestParameters struct {

    /**
    * the name of the PNCounter
    */
name string

    /**
    * the delta to add to the counter value, can be negative
    */
delta int64

    /**
    * {@code true} if the operation should return the
    * counter value before the addition, {@code false}
    * if it should return the value after the addition
    */
getBeforeUpdate bool

    /**
    * last observed replica timestamps (vector clock)
    */
replicaTimestamps []map[uuid]int64

    /**
    * the target replica
    */
targetReplicaUUID string
}

func PNCounterAddEncodeRequest(name string, delta int64, getBeforeUpdate bool, replicaTimestamps []map[uuid]int64, targetReplicaUUID string) *bufutil.ClientMessage {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.Is_Retryable = false
    clientMessage.SetOperationName("PNCounter.Add")
	initialFrame := &bufutil.Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, PNCounterAddRequestMessageType)
    bufutil.EncodeLong(initialFrame.Content, PNCounterAddRequestDeltaFieldOffset, delta)
    bufutil.EncodeBoolean(initialFrame.Content, PNCounterAddRequestGetBeforeUpdateFieldOffset, getBeforeUpdate)
    bufutil.EncodeUUID(initialFrame.Content, PNCounterAddRequestTargetReplicaUUIDFieldOffset, targetReplicaUUID)
    clientMessage.Add(initialFrame)
    bufutil.StringCodecEncode(clientMessage, name)
    bufutil.EntryListUUIDLongCodecEncode(clientMessage, replicaTimestamps)
    return clientMessage
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type PNCounterAddResponseParameters struct {
    /**
    * Value of the counter.
    */
value int64
    /**
    * last observed replica timestamps (vector clock)
    */
replicaTimestamps []map[uuid]int64
    /**
    * Number of replicas that keep the state of this counter.
    */
replicaCount int32
}



func PNCounterAddDecodeResponse(clientMessage *bufutil.ClientMessage) *PNCounterAddResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(PNCounterAddResponseParameters)
    initialFrame := iterator.Next()
    response.value = bufutil.DecodeLong(initialFrame.Content, PNCounterAddResponseValueFieldOffset)
    response.replicaCount = bufutil.DecodeInt(initialFrame.Content, PNCounterAddResponseReplicaCountFieldOffset)
    response.replicaTimestamps = bufutil.EntryListUUIDLongCodecDecode(iterator)
    return response
}

