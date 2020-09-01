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

package main

import (
	"fmt"
	"strconv"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/core/predicate"
)

func main() {
	config := hazelcast.NewConfig()
	config.NetworkConfig().AddAddress("127.0.0.1:5701")

	client, err := hazelcast.NewClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	mp, _ := client.GetMap("predicateExample")

	for i := 1; i < 20; i++ {
		if i % 2 == 0 {
			mp.Put("student-"+strconv.Itoa(i), "classA")
		}else{
			mp.Put("student-"+strconv.Itoa(i), "classB")
		}
	}

	equalPredicate := predicate.Equal("this", "classA" )

	values, err := mp.ValuesWithPredicate(equalPredicate)
	pagingPredicate := predicate.PagingPredicate(equalPredicate, 4)

	values, err = mp.ValuesWithPredicate(pagingPredicate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("page-1 -> ")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i], ", ")
	}

	pagingPredicate.NextPage()
	fmt.Println("page-2 -> ")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i],", ")
	}

	pagingPredicate.NextPage()
	fmt.Println("page-3 -> ")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i],", ")
	}

	pagingPredicate = predicate.PagingPredicatePage(3)

	pagingPredicate.SetPage(5)
	values, err = mp.ValuesWithPredicate(pagingPredicate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("page-6 -> ")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i],", ")
	}
	fmt.Println()

	mp.Clear()
	client.Shutdown()
}

type Student struct {
name string
id int
className string
}

func NewStudent(name string, id int) *Student {
	return &Student{name: name, id: id}
}

func (s *Student) Name() string {
	return s.name
}

func (s *Student) Id() int {
	return s.id
}

func (s *Student) ClassName() string {
	return s.className
}

func (s *Student) SetClassName(name string) {
s.className = name
}

func (s *Student) CompareTo(other Student) int{
	return s.id - other.id
}

