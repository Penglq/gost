/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package math

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestFloat64(t *testing.T) {
	assert.True(t, CompareFloat64(12.3334, 12.3344, 0.01))
	assert.True(t, CompareFloat64(12.3334, 12.32981, 0.01))
	assert.False(t, CompareFloat64(12.3334, 12.0325, 0.01))
}

func TestCompareFloat32(t *testing.T) {
	assert.True(t, CompareFloat32(12.3334, 12.3344, 0.01))
	assert.True(t, CompareFloat32(12.3334, 12.32981, 0.01))
	assert.False(t, CompareFloat64(12.3334, 12.0325, 0.01))
}
