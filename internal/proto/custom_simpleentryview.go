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

type SimpleEntryView struct {
key serialization.Data
value serialization.Data
cost int64
creationTime int64
expirationTime int64
hits int64
lastAccessTime int64
lastStoredTime int64
lastUpdateTime int64
version int64
ttl int64
maxIdle int64
}

func (s SimpleEntryView) Key() serialization.Data {
    return s.key
}

func (s SimpleEntryView) Value() serialization.Data {
    return s.value
}

func (s SimpleEntryView) ExpirationTime() int64 {
    return s.expirationTime
}

func (s SimpleEntryView) Hits() int64 {
    return s.hits
}

func (s SimpleEntryView) LastAccessTime() int64 {
    return s.lastAccessTime
}

func (s SimpleEntryView) LastStoredTime() int64 {
    return s.lastStoredTime
}

func (s SimpleEntryView) LastUpdateTime() int64 {
    return s.lastUpdateTime
}

func (s SimpleEntryView) Version() int64 {
    return s.version
}

func (s SimpleEntryView) Ttl() int64 {
    return s.ttl
}

func (s SimpleEntryView) MaxIdle() int64 {
    return s.maxIdle
}

func (s SimpleEntryView) Cost() int64 {
    return s.cost
}

func (s SimpleEntryView) CreationTime() int64 {
    return s.creationTime
}

//@Generated("99733ff6560195f784d32f45e6c68ba0")
const (
    SimpleEntryViewCostFieldOffset = 0
    SimpleEntryViewCreationTimeFieldOffset = SimpleEntryViewCostFieldOffset + LongSizeInBytes
    SimpleEntryViewExpirationTimeFieldOffset = SimpleEntryViewCreationTimeFieldOffset + LongSizeInBytes
    SimpleEntryViewHitsFieldOffset = SimpleEntryViewExpirationTimeFieldOffset + LongSizeInBytes
    SimpleEntryViewLastAccessTimeFieldOffset = SimpleEntryViewHitsFieldOffset + LongSizeInBytes
    SimpleEntryViewLastStoredTimeFieldOffset = SimpleEntryViewLastAccessTimeFieldOffset + LongSizeInBytes
    SimpleEntryViewLastUpdateTimeFieldOffset = SimpleEntryViewLastStoredTimeFieldOffset + LongSizeInBytes
    SimpleEntryViewVersionFieldOffset = SimpleEntryViewLastUpdateTimeFieldOffset + LongSizeInBytes
    SimpleEntryViewTtlFieldOffset = SimpleEntryViewVersionFieldOffset + LongSizeInBytes
    SimpleEntryViewMaxIdleFieldOffset = SimpleEntryViewTtlFieldOffset + LongSizeInBytes
    SimpleEntryViewInitialFrameSize = SimpleEntryViewMaxIdleFieldOffset + LongSizeInBytes
)

func SimpleEntryViewCodecEncode(clientMessage *ClientMessage, simpleEntryView SimpleEntryView) {
        clientMessage.Add(BeginFrame)
        initialFrame := &Frame{Content: make([]byte, SimpleEntryViewInitialFrameSize), Flags: UnfragmentedMessage}
        EncodeLong(initialFrame.Content, SimpleEntryViewCostFieldOffset, simpleEntryView.cost)
        EncodeLong(initialFrame.Content, SimpleEntryViewCreationTimeFieldOffset, simpleEntryView.creationTime)
        EncodeLong(initialFrame.Content, SimpleEntryViewExpirationTimeFieldOffset, simpleEntryView.expirationTime)
        EncodeLong(initialFrame.Content, SimpleEntryViewHitsFieldOffset, simpleEntryView.hits)
        EncodeLong(initialFrame.Content, SimpleEntryViewLastAccessTimeFieldOffset, simpleEntryView.lastAccessTime)
        EncodeLong(initialFrame.Content, SimpleEntryViewLastStoredTimeFieldOffset, simpleEntryView.lastStoredTime)
        EncodeLong(initialFrame.Content, SimpleEntryViewLastUpdateTimeFieldOffset, simpleEntryView.lastUpdateTime)
        EncodeLong(initialFrame.Content, SimpleEntryViewVersionFieldOffset, simpleEntryView.version)
        EncodeLong(initialFrame.Content, SimpleEntryViewTtlFieldOffset, simpleEntryView.ttl)
        EncodeLong(initialFrame.Content, SimpleEntryViewMaxIdleFieldOffset, simpleEntryView.maxIdle)
        clientMessage.Add(initialFrame)
        DataCodecEncode(clientMessage, simpleEntryView.key)
        DataCodecEncode(clientMessage, simpleEntryView.value)

        clientMessage.Add(EndFrame)
    }

func SimpleEntryViewCodecDecode(iterator *ForwardFrameIterator)  SimpleEntryView  {
        // begin frame
        iterator.Next()
        initialFrame := iterator.Next()
        cost := DecodeLong(initialFrame.Content, SimpleEntryViewCostFieldOffset)
        creationTime := DecodeLong(initialFrame.Content, SimpleEntryViewCreationTimeFieldOffset)
        expirationTime := DecodeLong(initialFrame.Content, SimpleEntryViewExpirationTimeFieldOffset)
        hits := DecodeLong(initialFrame.Content, SimpleEntryViewHitsFieldOffset)
        lastAccessTime := DecodeLong(initialFrame.Content, SimpleEntryViewLastAccessTimeFieldOffset)
        lastStoredTime := DecodeLong(initialFrame.Content, SimpleEntryViewLastStoredTimeFieldOffset)
        lastUpdateTime := DecodeLong(initialFrame.Content, SimpleEntryViewLastUpdateTimeFieldOffset)
        version := DecodeLong(initialFrame.Content, SimpleEntryViewVersionFieldOffset)
        ttl := DecodeLong(initialFrame.Content, SimpleEntryViewTtlFieldOffset)
        maxIdle := DecodeLong(initialFrame.Content, SimpleEntryViewMaxIdleFieldOffset)
        key := DataCodecDecode(iterator)
        value := DataCodecDecode(iterator)
        FastForwardToEndFrame(iterator)
        return SimpleEntryView { key, value, cost, creationTime, expirationTime, hits, lastAccessTime, lastStoredTime, lastUpdateTime, version, ttl, maxIdle }
    }