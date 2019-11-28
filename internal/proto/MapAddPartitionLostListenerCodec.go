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
 * Adds a MapPartitionLostListener. The addPartitionLostListener returns a register-id. This id is needed to remove
 * the MapPartitionLostListener using the removePartitionLostListener(String) method.
 * There is no check for duplicate registrations, so if you register the listener twice, it will get events twice.
 * IMPORTANT: Please see com.hazelcast.partition.PartitionLostListener for weaknesses.
 * IMPORTANT: Listeners registered from HazelcastClient may miss some of the map partition lost events due
 * to design limitations.
 */
//@Generated("74e41798538f0673692fed9f39fe4fee")
const (
    //hex: 0x011B00
    MapAddPartitionLostListenerRequestMessageType = 72448
    //hex: 0x011B01
    MapAddPartitionLostListenerResponseMessageType = 72449
    MapAddPartitionLostListenerRequestLocalOnlyFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapAddPartitionLostListenerRequestInitialFrameSize = MapAddPartitionLostListenerRequestLocalOnlyFieldOffset + bufutil.BooleanSizeInBytes
    MapAddPartitionLostListenerResponseResponseFieldOffset = bufutil.CorrelationIdFieldOffset + bufutil.IntSizeInBytes
    MapAddPartitionLostListenerResponseInitialFrameSize = MapAddPartitionLostListenerResponseResponseFieldOffset + bufutil.UUIDSizeInBytes
    MapAddPartitionLostListenerEventMapPartitionLostPartitionIdFieldOffset = bufutil.PartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapAddPartitionLostListenerEventMapPartitionLostUuidFieldOffset = MapAddPartitionLostListenerEventMapPartitionLostPartitionIdFieldOffset + bufutil.IntSizeInBytes
    MapAddPartitionLostListenerEventMapPartitionLostInitialFrameSize = MapAddPartitionLostListenerEventMapPartitionLostUuidFieldOffset + bufutil.UUIDSizeInBytes
    //hex: 0x011B02
    MapAddPartitionLostListenerEventMapPartitionLostMessageType = 72450


)

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapAddPartitionLostListenerRequestParameters struct {

    /**
    * name of map
    */
name string

    /**
    * if true fires events that originated from this node only, otherwise fires all events
    */
localOnly bool
}

func MapAddPartitionLostListenerEncodeRequest(name string, localOnly bool) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
    clientMessage.IsRetryable = false
    clientMessage.SetAcquiresResource(false)
    clientMessage.SetOperationName("Map.AddPartitionLostListener")
	initialFrame := bufutil.Frame{make([]byte, ListAddAllResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapAddPartitionLostListenerRequestMessageType)
    bufutil.EncodeBoolean(initialFrame.Content, MapAddPartitionLostListenerRequestLocalOnlyFieldOffset, localOnly)
    clientMessage.Add(initialFrame)
    StringCodec.encode(clientMessage, name)
    return clientMessage
}

func MapAddPartitionLostListenerDecodeRequest(clientMessage *bufutil.ClientMessagex) *MapAddPartitionLostListenerRequestParameters {
    iterator := clientMessage.FrameIterator()
    request := new(MapAddPartitionLostListenerRequestParameters)
    initialFrame := iterator.Next()
    request.localOnly = bufutil.DecodeBoolean(initialFrame.Content, MapAddPartitionLostListenerRequestLocalOnlyFieldOffset)
    request.name = StringCodec.decode(iterator)
    return request
}

/* tslint:disabled:URF-UNREAD-PUBLIC-OR-PROTECTED-FIELD */
type MapAddPartitionLostListenerResponseParameters struct {
    /**
    * returns the registration id for the MapPartitionLostListener.
    */
response UUID
}

func MapAddPartitionLostListenerEncodeResponse(response UUID ) *bufutil.ClientMessagex {
    clientMessage := bufutil.CreateForEncode()
	initialFrame := bufutil.Frame{make([]byte, MapAddPartitionLostListenerResponseInitialFrameSize), bufutil.UnfragmentedMessage}
    bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapAddPartitionLostListenerResponseMessageType)
    bufutil.EncodeUUID(initialFrame.Content, MapAddPartitionLostListenerResponseResponseFieldOffset, response)
    clientMessage.Add(initialFrame)

    return clientMessage
}

func MapAddPartitionLostListenerDecodeResponse(clientMessage *bufutil.ClientMessagex) *MapAddPartitionLostListenerResponseParameters {
    iterator := clientMessage.FrameIterator()
    response := new(MapAddPartitionLostListenerResponseParameters)
    initialFrame := iterator.Next()
    response.response = bufutil.DecodeUUID(initialFrame.Content, MapAddPartitionLostListenerResponseResponseFieldOffset)
    return response
}

    func encodeMapAddPartitionLostListenerEvent(partitionId int, uuid UUID) *bufutil.ClientMessagex {
        clientMessage := bufutil.createForEncode()
    	initialFrame := bufutil.Frame{make([]byte, MapAddPartitionLostListenerEventMapPartitionLostInitialFrameSize), bufutil.UNFRAGMENTED_MESSAGE}
        initialFrame.flags |= ClientMessage.IS_EventFLAG
        bufutil.EncodeInt(initialFrame.Content, bufutil.TypeFieldOffset, MapAddPartitionLostListenerEventMapPartitionLostMessageType)
        encodeInt(initialFrame.Content, MapAddPartitionLostListenerEventMapPartitionLostPartitionIdFieldOffset, partitionId)
        encodeUUID(initialFrame.Content, MapAddPartitionLostListenerEventMapPartitionLostUuidFieldOffset, uuid)
        clientMessage.Add(initialFrame)

        return clientMessage
    }

    type MapAddPartitionLostListenerAbstractEventItemFunc func(partitionId int, uuid UUID)

     func MapAddPartitionLostListenerHandle(clientMessage *bufutil.ClientMessagex, handleEventItem MapAddPartitionLostListenerAbstractEventItemFunc){
        messageType := clientMessage.getMessageType()
        iterator := clientMessage.FrameIterator()
         if messageType == MapAddPartitionLostListenerEventMapPartitionLostMessageType {
             initialFrame := iterator.Next()
             partitionId := bufutil.DecodeInt(initialFrame.Content, MapAddPartitionLostListenerEventMapPartitionLostPartitionIdFieldOffset)
             uuid := bufutil.DecodeUUID(initialFrame.Content, MapAddPartitionLostListenerEventMapPartitionLostUuidFieldOffset)
             handleEventItem(partitionId, uuid)
             return
          }
         Logger.getLogger(super.getClass()).finest("Unknown message type received on event handler :" + messageType)
     }

