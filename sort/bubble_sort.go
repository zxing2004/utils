package sort

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

// BubbleSort bubble sorting, rule parameter refers to the incoming sorting rule.
// from big to small, please pass the parameter "gt";
// from small to large, please pass the parameter "lt";
// BubbleSort 冒泡排序，rule参数指传入的排序规则
// 从大到小请传参数“gt”，从小到大请传参数“lt”.
func BubbleSort(arr *[]int, rule string) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if rule == "gt" && (*arr)[j] < (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
			if rule == "lt" && (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
