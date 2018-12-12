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

import (
	"errors"
	"regexp"
)

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

// EmailCheck EmailCheck
func EmailCheck(email string) (bool, error) {
	T, err := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`, email)
	return T, err
}

//VehicleNumberCheck the vehicle Number inspection for China area
func VehicleNumberCheck(VehicleNumber string) (bool, error) {
	T, err := regexp.MatchString(`^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领A-Z]{1}[A-Z]{1}[A-Z0-9]{4}[A-Z0-9挂学警港澳]{1}$`, VehicleNumber)
	return T, err
}

//IPv4Check IPv4 format Check
func IPv4Check(IPv4 string) (bool, error) {
	T, err := regexp.MatchString(`^((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))$`, IPv4)
	return T, err
}

// URLCheck Support HTTP, HTTPS, FTP url detection ,example: "https://github.com/zxing2004/utils"
func URLCheck(url string) (bool, error) {
	T, err := regexp.MatchString(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`, url)
	return T, err
}

// IDCardCheck Id verification supports 15 and 18 bit lengths
func IDCardCheck(idcard string) (bool, error) {
	err := errors.New("The length of the id card is illegal")
	if len(idcard) == 18 {
		T, err := regexp.MatchString(`^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`, idcard)
		return T, err
	}
	if len(idcard) == 15 {
		T, err := regexp.MatchString(`^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}$`, idcard)
		return T, err
	}
	return false, err
}
