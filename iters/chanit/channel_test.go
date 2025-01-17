// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chanit_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/chanit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
)

func TestChannel(t *testing.T) {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	s := sliceit.To(chanit.Channel(ch))
	assertpkg.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
}
