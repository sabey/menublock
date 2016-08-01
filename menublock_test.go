package menublock

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"fmt"
	"sabey.co/unittest"
	"testing"
)

func TestMenuBlock(t *testing.T) {
	fmt.Println("TestMenuBlock")

	for _, mt := range []*menutest{
		// 1-1h
		&menutest{
			items:  1,
			height: 1,
		},
		&menutest{
			items:  1,
			height: 1,
			up:     1,
		},
		&menutest{
			items:  1,
			height: 1,
			down:   1,
		},
		&menutest{
			items:  1,
			height: 1,
			jumpup: 1,
		},
		&menutest{
			items:  1,
			height: 1,
			jumpup: 2,
		},
		&menutest{
			items:    1,
			height:   1,
			jumpdown: 1,
		},
		&menutest{
			items:    1,
			height:   1,
			jumpdown: 2,
		},
		&menutest{
			items:  1,
			height: 1,
			top:    true,
		},
		&menutest{
			items:  1,
			height: 1,
			bottom: true,
		},
		// 1-2h
		&menutest{
			items:  1,
			height: 2,
		},
		&menutest{
			items:  1,
			height: 2,
			up:     1,
		},
		&menutest{
			items:  1,
			height: 2,
			down:   1,
		},
		&menutest{
			items:  1,
			height: 2,
			jumpup: 1,
		},
		&menutest{
			items:  1,
			height: 2,
			jumpup: 2,
		},
		&menutest{
			items:    1,
			height:   2,
			jumpdown: 1,
		},
		&menutest{
			items:    1,
			height:   2,
			jumpdown: 2,
		},
		&menutest{
			items:  1,
			height: 2,
			top:    true,
		},
		&menutest{
			items:  1,
			height: 2,
			bottom: true,
		},
		// 2-1h
		&menutest{
			items:     2,
			height:    1,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			up:        2,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			up:        1,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			up:        3,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			down:      1,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			down:      2,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			down:      3,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			jumpup:    1,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			jumpup:    2,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			jumpup:    3,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			jumpdown:  1,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			jumpdown:  2,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			jumpdown:  3,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     2,
			height:    1,
			top:       true,
			has_below: true,
		},
		&menutest{
			items:     2,
			height:    1,
			bottom:    true,
			item:      1,
			line:      1,
			has_above: true,
		},
		// 2-2h
		&menutest{
			items:  2,
			height: 2,
		},
		&menutest{
			items:  2,
			height: 2,
			up:     1,
			item:   1,
		},
		&menutest{
			items:  2,
			height: 2,
			up:     2,
		},
		&menutest{
			items:  2,
			height: 2,
			up:     3,
			item:   1,
		},
		&menutest{
			items:  2,
			height: 2,
			down:   1,
			item:   1,
		},
		&menutest{
			items:  2,
			height: 2,
			down:   2,
		},
		&menutest{
			items:  2,
			height: 2,
			down:   3,
			item:   1,
		},
		&menutest{
			items:  2,
			height: 2,
			jumpup: 1,
		},
		&menutest{
			items:  2,
			height: 2,
			jumpup: 2,
		},
		&menutest{
			items:  2,
			height: 2,
			jumpup: 3,
		},
		&menutest{
			items:    2,
			height:   2,
			jumpdown: 1,
			item:     1,
		},
		&menutest{
			items:    2,
			height:   2,
			jumpdown: 2,
			item:     1,
		},
		&menutest{
			items:    2,
			height:   2,
			jumpdown: 3,
			item:     1,
		},
		&menutest{
			items:  2,
			height: 2,
			top:    true,
		},
		&menutest{
			items:  2,
			height: 2,
			bottom: true,
			item:   1,
		},
		// 3-1h
		&menutest{
			items:     3,
			height:    1,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			up:        1,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    1,
			up:        2,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			up:        3,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			down:      1,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			down:      2,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    1,
			down:      3,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpup:    1,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpup:    2,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpup:    3,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpup:    4,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpdown:  1,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpdown:  2,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpdown:  3,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    1,
			jumpdown:  4,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    1,
			top:       true,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    1,
			bottom:    true,
			item:      2,
			line:      2,
			has_above: true,
		},
		// 3-2h
		&menutest{
			items:     3,
			height:    2,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			up:        1,
			item:      2,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    2,
			up:        2,
			item:      1,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			up:        3,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			down:      1,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    2,
			down:      2,
			item:      2,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    2,
			down:      3,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			jumpup:    1,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			jumpup:    2,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			jumpup:    3,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			jumpdown:  1,
			item:      1,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    2,
			jumpdown:  2,
			item:      2,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    2,
			jumpdown:  3,
			item:      2,
			line:      1,
			has_above: true,
		},
		&menutest{
			items:     3,
			height:    2,
			top:       true,
			has_below: true,
		},
		&menutest{
			items:     3,
			height:    2,
			bottom:    true,
			item:      2,
			line:      1,
			has_above: true,
		},
		// 3-3h
		&menutest{
			items:  3,
			height: 3,
		},
		&menutest{
			items:  3,
			height: 3,
			up:     1,
			item:   2,
		},
		&menutest{
			items:  3,
			height: 3,
			up:     2,
			item:   1,
		},
		&menutest{
			items:  3,
			height: 3,
			up:     3,
		},
		&menutest{
			items:  3,
			height: 3,
			down:   1,
			item:   1,
		},
		&menutest{
			items:  3,
			height: 3,
			down:   2,
			item:   2,
		},
		&menutest{
			items:  3,
			height: 3,
			down:   3,
		},
		&menutest{
			items:  3,
			height: 3,
			jumpup: 1,
		},
		&menutest{
			items:  3,
			height: 3,
			jumpup: 2,
		},
		&menutest{
			items:    3,
			height:   3,
			jumpdown: 1,
			item:     1,
		},
		&menutest{
			items:    3,
			height:   3,
			jumpdown: 2,
			item:     2,
		},
		&menutest{
			items:    3,
			height:   3,
			jumpdown: 3,
			item:     2,
		},
		&menutest{
			items:    3,
			height:   3,
			jumpdown: 4,
			item:     2,
		},
		&menutest{
			items:  3,
			height: 3,
			top:    true,
		},
		&menutest{
			items:  3,
			height: 3,
			bottom: true,
			item:   2,
		},
		// 4-1h
		&menutest{
			items:     4,
			height:    1,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			up:        1,
			item:      3,
			line:      3,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    1,
			up:        2,
			item:      2,
			line:      2,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			up:        3,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			up:        4,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			down:      1,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			down:      2,
			item:      2,
			line:      2,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			down:      3,
			item:      3,
			line:      3,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    1,
			down:      4,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpup:    1,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpup:    2,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpup:    3,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpup:    4,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpdown:  1,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpdown:  2,
			item:      2,
			line:      2,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpdown:  3,
			item:      3,
			line:      3,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    1,
			jumpdown:  4,
			item:      3,
			line:      3,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    1,
			top:       true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    1,
			bottom:    true,
			item:      3,
			line:      3,
			has_above: true,
		},
		// 4-1h
		&menutest{
			items:     4,
			height:    2,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			up:        1,
			item:      3,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    2,
			up:        2,
			item:      2,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			up:        3,
			item:      1,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			up:        4,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			down:      1,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			down:      2,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    2,
			down:      3,
			item:      3,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    2,
			down:      4,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpup:    1,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpup:    2,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpup:    3,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpup:    4,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpdown:  1,
			item:      1,
			line:      1,
			has_above: true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpdown:  2,
			item:      2,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpdown:  3,
			item:      3,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    2,
			jumpdown:  4,
			item:      3,
			line:      2,
			has_above: true,
		},
		&menutest{
			items:     4,
			height:    2,
			top:       true,
			has_below: true,
		},
		&menutest{
			items:     4,
			height:    2,
			bottom:    true,
			item:      3,
			line:      2,
			has_above: true,
		},
	} {
		m := Create(mt.items)
		//fmt.Printf("\nitems: %d height: %d top: %t bottom: %t jumpup: %d jumpdown: %d up: %d down: %d\n", mt.items, mt.height, mt.top, mt.bottom, mt.jumpup, mt.jumpdown, mt.up, mt.down)
		m.ResetMaybe(mt.height) // init
		// move after we set height
		if mt.bottom {
			m.Bottom()
		}
		if mt.jumpdown > 0 {
			m.JumpDown(mt.jumpdown)
		}
		if mt.jumpup > 0 {
			m.JumpUp(mt.jumpup)
		}
		if mt.down > 0 {
			for i := 0; i < mt.down; i++ {
				m.Down()
			}
		}
		if mt.up > 0 {
			for i := 0; i < mt.up; i++ {
				m.Up()
			}
		}
		if mt.top {
			m.Top()
		}
		// get output after movement
		//fmt.Printf("line: %d\n", m.GetLine())
		//fmt.Printf("item: %d\n", m.GetItem())
		unittest.Equals(t, m.GetHeight(), mt.height)
		unittest.Equals(t, m.GetLine(), mt.line)
		unittest.Equals(t, m.GetItem(), mt.item)
		unittest.Equals(t, m.HasAbove(), mt.has_above)
		unittest.Equals(t, m.HasBelow(), mt.has_below)
	}
}

type menutest struct {
	items     int
	height    int
	line      int
	item      int
	bottom    bool
	jumpdown  int
	jumpup    int
	down      int
	up        int
	top       bool
	has_above bool
	has_below bool
}
