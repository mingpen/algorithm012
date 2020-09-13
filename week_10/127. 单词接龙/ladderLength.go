package main

//https://leetcode-cn.com/problems/word-ladder/

// 使用bfs
// 从begin开始，被确认过的word无需再次访问，
// 从中心开始的,因为都会与endword对比

import (
	"container/list"
	"log"
)

func main() {
	log.Println(
		ladderLength("hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
}

type solution struct {
	beginWord, endWord string
	wordList           []string
	n                  int

	listNode *list.List
	charSize int

	confirmed []bool
}

// Node  ...
type Node struct {
	level     int
	beginWord string
}

func (s *solution) bfs(node *Node) (int, bool) {
	//
	for i, ok := range s.confirmed {
		if ok {
			continue
		}
		if s.oneChar(node.beginWord, s.wordList[i]) {
			if s.endWord == s.wordList[i] {
				return node.level + 1, true
			}
			node2 := &Node{
				level:     node.level + 1,
				beginWord: s.wordList[i],
			}
			s.listNode.PushBack(node2)
			s.confirmed[i] = true
		}
	}
	return 0, false
}

func (s *solution) ladderLength() int {
	s.n = len(s.wordList)
	s.charSize = len(s.beginWord)
	s.listNode = list.New()
	s.confirmed = make([]bool, s.n)
	node := &Node{level: 1, beginWord: s.beginWord}
	s.listNode.PushBack(node)
	for ele := s.listNode.Front(); nil != ele; ele = s.listNode.Front() {
		s.listNode.Remove(ele)
		if cnt, ok := s.bfs(ele.Value.(*Node)); ok {
			return cnt
		}
	}

	return 0
}

func (s *solution) oneChar(src string, dst string) bool {
	var diff int
	for i := 0; i < s.charSize; i++ {
		if src[i] != dst[i] {
			diff++
		}
	}
	return diff == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	s := &solution{beginWord: beginWord, endWord: endWord, wordList: wordList}
	return s.ladderLength()
}
