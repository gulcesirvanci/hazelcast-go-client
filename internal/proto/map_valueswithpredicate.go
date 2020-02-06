
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
     "github.com/hazelcast/hazelcast-go-client/serialization"
)


/*
 * This file is auto-generated by the Hazelcast Client Protocol Code Generator.
 * To change this file, edit the templates or the protocol
 * definitions on the https://github.com/hazelcast/hazelcast-client-protocol
 * and regenerate it.
 */

/**
 * Queries the map based on the specified predicate and returns the values of matching entries.Specified predicate
 * runs on all members in parallel. The collection is NOT backed by the map, so changes to the map are NOT reflected
 * in the collection, and vice-versa. This method is always executed by a distributed query, so it may throw a
 * QueryResultSizeExceededException if query result size limit is configured.
 */
//@Generated("40ebe2facfcf6b47bb2ad186e927f6d0")
const (
    //hex: 0x012700
    MapValuesWithPredicateRequestMessageType = 75520
    //hex: 0x012701
    MapValuesWithPredicateResponseMessageType = 75521
    MapValuesWithPredicateRequestInitialFrameSize = PartitionIdFieldOffset + IntSizeInBytes
    MapValuesWithPredicateResponseInitialFrameSize = CorrelationIdFieldOffset + IntSizeInBytes

)

func MapValuesWithPredicateEncodeRequest(name string, predicate serialization.Data) *ClientMessage {
    clientMessage := CreateForEncode()
    clientMessage.SetRetryable( true )
    clientMessage.SetOperationName("Map.ValuesWithPredicate")
	initialFrame := &Frame{Content: make([]byte, ListAddAllResponseInitialFrameSize), Flags: UnfragmentedMessage}
    EncodeInt(initialFrame.Content, TypeFieldOffset, MapValuesWithPredicateRequestMessageType)
    clientMessage.Add(initialFrame)

    StringCodecEncode(clientMessage, name)


    DataCodecEncode(clientMessage, predicate)

    return clientMessage
}


func MapValuesWithPredicateDecodeResponse(clientMessage *ClientMessage) func() (/*** result value collection of the query.*/response []serialization.Data) {
    return func() (/*** result value collection of the query.*/response []serialization.Data) {
        iterator := clientMessage.FrameIterator()
        //empty initial frame
        iterator.Next()
        var result []serialization.Data
        //begin frame, list
        iterator.Next()
        for !NextFrameIsDataStructureEndFrame(iterator) {
        result = append(result, DataCodecDecode(iterator))
        }
        //end frame, list
        iterator.Next()
        response = result //0.1
        return
    }
}

