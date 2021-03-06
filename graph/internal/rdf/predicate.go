/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rdf

import "github.com/google/badwolf/triple/predicate"

var (
	ParentOfPredicate  *predicate.Predicate
	AppliesOnPredicate *predicate.Predicate
	HasTypePredicate   *predicate.Predicate
	DiffPredicate      *predicate.Predicate
	PropertyPredicate  *predicate.Predicate
	MetaPredicate      *predicate.Predicate
)

func init() {
	var err error
	if ParentOfPredicate, err = predicate.NewImmutable("parent_of"); err != nil {
		panic(err)
	}
	if AppliesOnPredicate, err = predicate.NewImmutable("applies_on"); err != nil {
		panic(err)
	}
	if MetaPredicate, err = predicate.NewImmutable("meta"); err != nil {
		panic(err)
	}
	if HasTypePredicate, err = predicate.NewImmutable("has_type"); err != nil {
		panic(err)
	}
	if DiffPredicate, err = predicate.NewImmutable("diff"); err != nil {
		panic(err)
	}
	if PropertyPredicate, err = predicate.NewImmutable("property"); err != nil {
		panic(err)
	}
	DefaultDiffer = &hierarchicDiffer{ParentOfPredicate}
}
