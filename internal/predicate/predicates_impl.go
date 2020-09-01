// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package predicate

import (
	"fmt"
	"github.com/hazelcast/hazelcast-go-client/core"
	"github.com/hazelcast/hazelcast-go-client/serialization"
	"reflect"
)

type predicate struct {
	id int32
}

func newPredicate(id int32) *predicate {
	return &predicate{id}
}

func (p *predicate) ReadData(input serialization.DataInput) error {
	return nil
}

func (p *predicate) WriteData(output serialization.DataOutput) error {
	return nil
}

func (*predicate) FactoryID() int32 {
	return FactoryID
}

func (p *predicate) ClassID() int32 {
	return p.id
}

type SQL struct {
	*predicate
	sql string
}

func NewSQL(sql string) *SQL {
	return &SQL{newPredicate(sqlID), sql}
}

func (sp *SQL) ReadData(input serialization.DataInput) error {
	sp.predicate = newPredicate(sqlID)
	sp.sql = input.ReadUTF()
	return input.Error()
}

func (sp *SQL) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(sp.sql)
	return nil
}

type And struct {
	*predicate
	predicates []interface{}
}

func NewAnd(predicates []interface{}) *And {
	return &And{newPredicate(andID), predicates}
}

func (ap *And) ReadData(input serialization.DataInput) error {
	ap.predicate = newPredicate(andID)
	length := input.ReadInt32()
	ap.predicates = make([]interface{}, length)
	for i := 0; i < int(length); i++ {
		pred := input.ReadObject()
		ap.predicates[i] = pred
	}
	return input.Error()
}

func (ap *And) WriteData(output serialization.DataOutput) error {
	output.WriteInt32(int32(len(ap.predicates)))
	for _, pred := range ap.predicates {
		err := output.WriteObject(pred)
		if err != nil {
			return err
		}
	}
	return nil
}

type Between struct {
	*predicate
	field string
	from  interface{}
	to    interface{}
}

func NewBetween(field string, from interface{}, to interface{}) *Between {
	return &Between{newPredicate(betweenID), field, from, to}
}

func (bp *Between) ReadData(input serialization.DataInput) error {
	bp.predicate = newPredicate(betweenID)
	bp.field = input.ReadUTF()
	bp.to = input.ReadObject()
	bp.from = input.ReadObject()
	return input.Error()
}

func (bp *Between) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(bp.field)
	err := output.WriteObject(bp.to)
	if err != nil {
		return err
	}
	return output.WriteObject(bp.from)
}

type Equal struct {
	*predicate
	field string
	value interface{}
}

func NewEqual(field string, value interface{}) *Equal {
	return &Equal{newPredicate(equalID), field, value}
}

func (ep *Equal) ReadData(input serialization.DataInput) error {
	ep.predicate = newPredicate(equalID)
	ep.field = input.ReadUTF()
	ep.value = input.ReadObject()
	return input.Error()
}

func (ep *Equal) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(ep.field)
	return output.WriteObject(ep.value)
}

type GreaterLess struct {
	*predicate
	field string
	value interface{}
	equal bool
	less  bool
}

func NewGreaterLess(field string, value interface{}, equal bool, less bool) *GreaterLess {
	return &GreaterLess{newPredicate(greaterlessID), field, value, equal, less}
}

func (glp *GreaterLess) ReadData(input serialization.DataInput) error {
	glp.predicate = newPredicate(greaterlessID)
	glp.field = input.ReadUTF()
	glp.value = input.ReadObject()
	glp.equal = input.ReadBool()
	glp.less = input.ReadBool()
	return input.Error()
}

func (glp *GreaterLess) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(glp.field)
	err := output.WriteObject(glp.value)
	if err != nil {
		return err
	}
	output.WriteBool(glp.equal)
	output.WriteBool(glp.less)
	return nil
}

type Like struct {
	*predicate
	field string
	expr  string
}

func NewLike(field string, expr string) *Like {
	return &Like{newPredicate(likeID), field, expr}
}

func (lp *Like) ReadData(input serialization.DataInput) error {
	lp.predicate = newPredicate(likeID)
	lp.field = input.ReadUTF()
	lp.expr = input.ReadUTF()
	return input.Error()
}

func (lp *Like) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(lp.field)
	output.WriteUTF(lp.expr)
	return nil
}

type ILike struct {
	*Like
}

func NewILike(field string, expr string) *ILike {
	return &ILike{&Like{newPredicate(ilikeID), field, expr}}
}

func (ilp *ILike) ReadData(input serialization.DataInput) error {
	ilp.Like = &Like{predicate: newPredicate(ilikeID)}
	ilp.field = input.ReadUTF()
	ilp.expr = input.ReadUTF()
	return input.Error()
}

type In struct {
	*predicate
	field  string
	values []interface{}
}

func NewIn(field string, values []interface{}) *In {
	return &In{newPredicate(inID), field, values}
}

func (ip *In) ReadData(input serialization.DataInput) error {
	ip.predicate = newPredicate(inID)
	ip.field = input.ReadUTF()
	length := input.ReadInt32()
	ip.values = make([]interface{}, length)
	for i := int32(0); i < length; i++ {
		ip.values[i] = input.ReadObject()
	}
	return input.Error()
}

func (ip *In) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(ip.field)
	output.WriteInt32(int32(len(ip.values)))
	for _, value := range ip.values {
		err := output.WriteObject(value)
		if err != nil {
			return err
		}
	}
	return nil
}

type InstanceOf struct {
	*predicate
	className string
}

func NewInstanceOf(className string) *InstanceOf {
	return &InstanceOf{newPredicate(instanceOfID), className}
}

func (iop *InstanceOf) ReadData(input serialization.DataInput) error {
	iop.predicate = newPredicate(instanceOfID)
	iop.className = input.ReadUTF()
	return input.Error()
}

func (iop *InstanceOf) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(iop.className)
	return nil
}

type NotEqual struct {
	*Equal
}

func NewNotEqual(field string, value interface{}) *NotEqual {
	return &NotEqual{&Equal{newPredicate(notEqualID), field, value}}
}

func (nep *NotEqual) ReadData(input serialization.DataInput) error {
	nep.Equal = &Equal{predicate: newPredicate(notEqualID)}
	nep.field = input.ReadUTF()
	nep.value = input.ReadObject()
	return input.Error()
}

type Not struct {
	*predicate
	pred interface{}
}

func NewNot(pred interface{}) *Not {
	return &Not{newPredicate(notID), pred}
}

func (np *Not) ReadData(input serialization.DataInput) error {
	np.predicate = newPredicate(notID)
	i := input.ReadObject()
	np.pred = i.(interface{})
	return input.Error()
}

func (np *Not) WriteData(output serialization.DataOutput) error {
	return output.WriteObject(np.pred)
}

type Or struct {
	*predicate
	predicates []interface{}
}

func NewOr(predicates []interface{}) *Or {
	return &Or{newPredicate(orID), predicates}
}

func (or *Or) ReadData(input serialization.DataInput) error {
	or.predicate = newPredicate(orID)
	length := input.ReadInt32()
	or.predicates = make([]interface{}, length)
	for i := 0; i < int(length); i++ {
		pred := input.ReadObject()
		or.predicates[i] = pred.(interface{})
	}
	return input.Error()
}

func (or *Or) WriteData(output serialization.DataOutput) error {
	output.WriteInt32(int32(len(or.predicates)))
	for _, pred := range or.predicates {
		err := output.WriteObject(pred)
		if err != nil {
			return err
		}
	}
	return nil
}

type Regex struct {
	*predicate
	field string
	regex string
}

func NewRegex(field string, regex string) *Regex {
	return &Regex{newPredicate(regexID), field, regex}
}

func (rp *Regex) ReadData(input serialization.DataInput) error {
	rp.predicate = newPredicate(regexID)
	rp.field = input.ReadUTF()
	rp.regex = input.ReadUTF()
	return input.Error()
}

func (rp *Regex) WriteData(output serialization.DataOutput) error {
	output.WriteUTF(rp.field)
	output.WriteUTF(rp.regex)
	return nil
}

type False struct {
	*predicate
}

func NewFalse() *False {
	return &False{newPredicate(falseID)}
}

func (fp *False) ReadData(input serialization.DataInput) error {
	fp.predicate = newPredicate(falseID)
	return nil
}

func (fp *False) WriteData(output serialization.DataOutput) error {
	//Empty method
	return nil
}

type True struct {
	*predicate
}

func NewTrue() *True {
	return &True{newPredicate(trueID)}
}

func (tp *True) ReadData(input serialization.DataInput) error {
	tp.predicate = newPredicate(trueID)
	return nil
}

func (tp *True) WriteData(output serialization.DataOutput) error {
	//Empty method
	return nil
}


type Paging struct {
	*predicate
	comparator core.Comparator
	anchorList map[int32]map[interface{}]interface{}
	pageSize int32
	page int32
	iterationType byte
}

 var nilAnchor = map[interface{}]interface{} {
 	-1 : nil,
}

func PagingPredicate(predicate interface{}, pageSize int32, comparator core.Comparator) *Paging {
	if IsInstanceOf(predicate, Paging{}) == true {
		core.NewHazelcastIllegalArgumentError("Nested PagingPredicate is not supported!", nil)
	}
	if pageSize <= 0 {
		core.NewHazelcastIllegalArgumentError("pageSize should be greater than 0!", nil)
	}
	iterationType := byte(2) //entry
	comparator=nil
	var anchorList map[int32]map[interface{}]interface{} // List of pairs: (nearest page, (anchor key, anchor value))
	return &Paging{newPredicate(pageSize),comparator, anchorList, pageSize, 0, iterationType}
}

func NewPagingPredicate(predicate interface{}, pageSize int32) *Paging {
	return PagingPredicate(predicate, pageSize, nil)
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

 //repr 
 
func IterationTypeName(iterationType byte) string {
	switch iterationType {
	case byte(0):
		return "KEY"
	case byte(1):
		return "VALUE"
	case byte(2):
		return "ENTRY"
	default:
		return ""
	}
}

func (paging *Paging) WriteData(output serialization.DataOutput) error {
	output.WriteObject(paging.predicate)
	output.WriteObject(paging.comparator)
	output.WriteInt32(paging.page)
	output.WriteInt32(paging.pageSize)
	output.WriteUTF(IterationTypeName(paging.iterationType))
	output.WriteInt32(int32(len(paging.anchorList)))
	for nearestPage := range paging.anchorList {
		output.WriteInt32(nearestPage)
		for anchorK, anchorV := range paging.anchorList[nearestPage] {
			output.WriteObject(anchorK)
			output.WriteObject(anchorV)
		}
	}
	return nil
}

func (paging *Paging) NextPage() {
	paging.page++
}

func (paging *Paging) PreviousPage() {
	if paging.page != 0 {
		paging.page--
	}
}

func (paging *Paging) SetPage(newPage int32) {
	if paging.page < 0 {
		core.NewHazelcastIllegalArgumentError("New page should be positive or 0.", nil)
	}
	paging.page = newPage
}

func (paging *Paging) SetIterationType(newIterType byte) {
	paging.iterationType = newIterType
}

func (paging *Paging) SetAnchor(page int32, anchor map[interface{}]interface{}) {
	var anchorEntry map[int32]map[interface{}]interface{}
		anchorEntry[page] = anchor

	anchorCount := int32(len(paging.anchorList))

	if page < anchorCount {
		paging.anchorList[page] = anchor
	} else if page == anchorCount {
		paging.anchorList[page+1] = anchor
	} else {
		core.NewHazelcastIllegalArgumentError(fmt.Sprintf("Anchor index is not correct, expected: %d, found: %d", page, anchorCount), nil)
	}
}

func (paging *Paging) Reset() {
	paging.iterationType = byte(2)
	for k := range paging.anchorList {
		delete(paging.anchorList, k)
	}
	paging.page = 0
}

func (paging *Paging) Page() int32 {
	return paging.page
}

func (paging *Paging) PageSize() int32 {
	return paging.pageSize
}

func (paging *Paging) NearestAnchorEntry() map[interface{}]interface{} {
	anchorCount := int32(len(paging.anchorList))
	if paging.page == 0 || anchorCount == 0 {
		return nilAnchor
	}
	var anchoredEntry map[interface{}]interface{}
	if paging.page < anchorCount {
		anchoredEntry = paging.anchorList[paging.page - 1]
	} else {
		anchoredEntry = paging.anchorList[anchorCount - 1]
	}
	return anchoredEntry
}

func (paging *Paging) IterationType() byte {
	return paging.iterationType
}

func (paging *Paging) Comparator() core.Comparator {
	return paging.comparator
}

func (paging *Paging) Predicate() *predicate {
	return paging.predicate
}
