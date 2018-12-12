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

package vaild

import "regexp"

// NameCheck Non-chinese name check
func NameCheck(name string) (bool, error) {
	T, err := regexp.MatchString("^\\p{Han}+$", name)
	return T, err
}

// PhoneNumberCheck Mobile phone number check in China area
func PhoneNumberCheck(PhoneNumber string) (bool, error) {
	T, err := regexp.MatchString(`^1([358][0-9]|4[579]|66|7[0135678]|9[89])[0-9]{8}$`, PhoneNumber)
	return T, err
}

//VehicleNumberCheck the vehicle Number inspection for China area
func VehicleNumberCheck(VehicleNumber string) (bool, error) {
	T, err := regexp.MatchString(`^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}[A-Z0-9]{4}[A-Z0-9挂学警港澳]{1}$`, VehicleNumber)
	return T, err

}

//IPv4Check IPv4 Check
func IPv4Check(IPv4 string) (bool, error) {
	T, err := regexp.MatchString(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`, IPv4)
	return T, err

}
