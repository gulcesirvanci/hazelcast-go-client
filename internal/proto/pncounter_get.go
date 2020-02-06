
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
    "github.com/hazelcast/hazelcast-go-client/core"




)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Query operation to retrieve the current value of the PNCounter.
 * <p>
 * The invocation will return the replica timestamps (vector clock) which
 * can then be sent with the next invocation to keep session consistency
 * guarantees.
 * The target replica is determined by the {@code targetReplica} parameter.
 * If smart routing is disabled, the actual member processing the client
 * message may act as a proxy.
 */
//@Generated("7a1156b95e31b1dbb7973c306753b451")
const (
    //hex: 0x1D0100
    PNCounterGetRequestMessageType = 1900800
    //hex: 0x1D0101
    PNCounterGetResponseMessageType = 1900801
    PNCounterGetRequestTargetReplicaUUIDFieldOffset = PartitionIdFieldOffset + IntSizeInBytes
    PNCounterGetRequestInitialFrameSize = PNCounterGetRequestTargetReplicaUUIDFieldOffset + UUIDSizeInBytes
    PNCounterGetResponseValueFieldOffset = CorrelationIdFieldOffset + IntSizeInBytes
    PNCounterGetResponseReplicaCountFieldOffset = PNCounterGetResponseValueFieldOffset + LongSizeInBytes
    PNCounterGetResponseInitialFrameSize = PNCounterGetResponseReplicaCountFieldOffset + IntSizeInBytes

)

func PNCounterGetEncodeRequest(name string, replicaTimestamps []  *Pair, targetReplicaUUID core.Uuid) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("PNCounter.Get")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, PNCounterGetRequestMessageType)
    EncodeUUID(initialFrame.Content, PNCounterGetRequestTargetReplicaUUIDFieldOffset, targetReplicaUUID)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    EntryListUUIDLongCodecEncode(clientMessage, replicaTimestamps)

    return clientMessage
}


func PNCounterGetDecodeResponse(clientMessage *ClientMessage) func() ( /*** Value of the counter.*/value int64, /*** last observed replica timestamps (vector clock)*/replicaTimestamps []  *Pair, /*** Number of replicas that keep the state of this counter.*/replicaCount int32) {
    return func() (/*** Value of the counter.*/value int64, /*** last observed replica timestamps (vector clock)*/replicaTimestamps []  *Pair, /*** Number of replicas that keep the state of this counter.*/replicaCount int32) {
        iterator := clientMessage.FrameIterator()
        initialFrame := iterator.Next()
        value = DecodeLong(initialFrame.Content, PNCounterGetResponseValueFieldOffset)
        replicaCount = DecodeInt(initialFrame.Content, PNCounterGetResponseReplicaCountFieldOffset)
    //check line for no nullable decode
        return
    }
}

