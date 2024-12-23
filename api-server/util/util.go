/*
Copyright 2023-2024 Omnissa, LLC.
SPDX-License-Identifier: Apache-2.0

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

package util

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
)

func RandomString(n int) string {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(fmt.Errorf("Error generating random: %s", err))
	}
	return hex.EncodeToString(bytes)
}

func ToJson(data interface{}) string {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(b)
}

func DeepCopyMap(from map[string]interface{}) map[string]interface{} {
	to := make(map[string]interface{})

	for k, v := range from {
		if v != nil && reflect.TypeOf(v).Kind() == reflect.Map {
			valueMap := v.(map[string]interface{})
			to[k] = DeepCopyMap(valueMap)
		} else {
			to[k] = v
		}
	}
	return to
}
