package menublock

// Copyright 2016, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

import (
	"sync"
)

type MenuBlock struct {
	items        int
	height       int
	item         int
	current_line int
	mu           sync.RWMutex
}

func Create(
	items int,
) *MenuBlock {
	if items < 1 {
		return nil
	}
	return &MenuBlock{
		items: items,
	}
}
func (self *MenuBlock) GetItems() int {
	return self.items
}
func (self *MenuBlock) GetItem() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.item
}
func (self *MenuBlock) GetLine() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.current_line
}
func (self *MenuBlock) GetItemLine() (
	int, // item
	int, // line
) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.item, self.current_line
}
func (self *MenuBlock) GetHeight() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.height
}
func (self *MenuBlock) Reset() {
	self.mu.Lock()
	self.resetMaybe(-1)
	self.mu.Unlock()
}
func (self *MenuBlock) ResetMaybe(
	height int,
) bool {
	self.mu.Lock()
	defer self.mu.Unlock()
	return self.resetMaybe(height)
}
func (self *MenuBlock) resetMaybe(
	height int,
) bool {
	if height <= 0 {
		// reset to unusable
		self.height = -1
		self.item = 0
		self.current_line = 0
		return true
	}
	if self.height != height {
		// reset
		self.height = height
		self.item = 0
		self.current_line = 0
		return true
	}
	return false
}
func (self *MenuBlock) HasAbove() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.current_line != 0
}
func (self *MenuBlock) HasBelow() bool {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return self.current_line < self.lowestLine()
}
func (self *MenuBlock) Top() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.height <= 0 {
		return
	}
	self.current_line = 0
	self.item = 0
}
func (self *MenuBlock) Up() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.height <= 0 {
		return
	}
	self.item--
	self.current_line--
	if self.item < 0 {
		// wrap around to the bottom
		self.item = self.items - 1
		self.current_line = self.lowestLine()
	}
	if self.current_line < 0 {
		self.current_line = 0
	}
}
func (self *MenuBlock) Down() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.height <= 0 {
		return
	}
	self.item++
	self.current_line++
	if self.item >= self.items {
		// wrap around to the top
		self.item = 0
		self.current_line = 0
	}
	lowest_line := self.lowestLine()
	if self.current_line > lowest_line {
		self.current_line = lowest_line
	}
}
func (self *MenuBlock) Bottom() {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.height <= 0 {
		return
	}
	self.item = self.items - 1
	self.current_line = self.lowestLine()
}
func (self *MenuBlock) lowestLine() int {
	lowest_line := self.items - self.height
	if lowest_line < 0 {
		lowest_line = 0
	}
	return lowest_line
}
func (self *MenuBlock) jump() {
	if self.item < 0 {
		// top
		self.item = 0
		self.current_line = 0
		return
	}
	lowest_line := self.lowestLine()
	if self.item >= self.items {
		// bottom
		self.item = self.items - 1
		self.current_line = lowest_line
		return
	}
	if self.current_line > lowest_line {
		self.current_line = lowest_line
	}
	if self.current_line < 0 {
		self.current_line = 0
	}
}
func (self *MenuBlock) JumpUp(
	i int,
) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.height <= 0 {
		return
	}
	self.item -= i
	self.current_line -= i
	self.jump()
}
func (self *MenuBlock) JumpDown(
	i int,
) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.height <= 0 {
		return
	}
	self.item += i
	self.current_line += i
	self.jump()
}
