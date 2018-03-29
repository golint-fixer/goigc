// Copyright ©2017 The ezgliding Authors.
//
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package igc

import (
	"fmt"
)

// NewDijkstraOptimizer returns a DijkstraOptimizer with the given characteristics.
//
func NewDijkstraOptimizer(cache bool) Optimizer {
	return &dijkstraOptimizer{cache: cache}
}

type dijkstraOptimizer struct {
	cache bool
}

func (b *dijkstraOptimizer) Optimize(track Track, nPoints int, score Score) (Task, error) {
	for i := 0; i < len(track.Points); i++ {
		fmt.Printf("%v", track.Points[i])
	}
	return Task{}, nil
}
