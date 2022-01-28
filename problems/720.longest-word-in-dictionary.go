package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/longest-word-in-dictionary

type TrieNode struct {
	IsEnd    bool
	Children map[rune]*TrieNode
	Words    map[string]bool
}

func longestWord(words []string) string {
	root := &TrieNode{true, map[rune]*TrieNode{}, map[string]bool{}}

	for _, word := range words {
		node := root

		for i, k := range word {
			n, ok := node.Children[k]
			if !ok {
				n = &TrieNode{false, map[rune]*TrieNode{}, map[string]bool{}}
				node.Children[k] = n
			}

			if (len(word) - 1) == i {
				n.IsEnd = true
				n.Words[word] = true
			}

			node = n
		}
	}

	stack := []*TrieNode{}

	nodes := []*TrieNode{root}
	for len(nodes) > 0 {
		_nodes := []*TrieNode{}

		for _, node := range nodes {
			var keys []rune
			for k := range node.Children {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool {
				return keys[i] > keys[j]
			})

			for _, k := range keys {
				v := node.Children[k]

				if v.IsEnd {
					stack = append(stack, v)
					_nodes = append(_nodes, v)
				}
			}
		}

		nodes = _nodes
	}

	if len(stack) == 0 {
		return ""
	}

	for k := range stack[len(stack)-1].Words {
		return k
	}

	return ""
}

func main() {
	r := longestWord([]string{
		"o", "a", "ajd", "ajdpw", "ojowj", "okpnd", "okpn", "ef", "oetj", "ajdp", "ojo", "o", "ok", "oet", "o", "oj", "ojowjy", "e",
	})
	fmt.Println(r)
}

// COMMENT: 前缀树 + 深度遍历
