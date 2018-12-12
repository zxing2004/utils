// Copyright 2018 zxing2004 authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package sort

// InsertionSort The working principle is to build an ordered sequence,
// scan the unsorted data from back to front in the sorted sequence, find the corresponding position and insert
func InsertionSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		Index := i - 1
		current := (*arr)[i]
		for Index >= 0 && (*arr)[Index] > current {
			(*arr)[Index+1] = (*arr)[Index]
			Index--
		}
		(*arr)[Index+1] = current
	}
}
