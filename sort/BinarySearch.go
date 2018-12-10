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

import (
	"fmt"
)

func BinarySearch(arr []int, khey int) (index int) {
	L := 0
	R := len(arr) - 1
	for {
		M := (L + (R - 1)) / 2
		if L > R {
			return -1
		}
		if arr[M] == khey {
			fmt.Println("Found", arr[M])
			return M
		} else if arr[M] > khey {
			R = M - 1
		} else {
			L = M + 1
		}
	}

	return -1
}
