package main

import "fmt"

// https://leetcode-cn.com/problems/word-search

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			used[i][j] = false
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if check(board, i, j, used, word) {
				return true
			}
		}
	}

	return false
}

func check(board [][]byte, i int, j int, used [][]bool, next string) bool {
	if len(next) == 0 {
		return true
	}
	if board[i][j] != next[0] {
		return false
	}
	if len(next) == 1 {
		return true
	}

	used[i][j] = true

	if i+1 < len(board) && !used[i+1][j] && check(board, i+1, j, used, next[1:]) {
		return true
	}
	if i-1 >= 0 && !used[i-1][j] && check(board, i-1, j, used, next[1:]) {
		return true
	}

	boradj := board[i]
	if j+1 < len(boradj) && !used[i][j+1] && check(board, i, j+1, used, next[1:]) {
		return true
	}
	if j-1 >= 0 && !used[i][j-1] && check(board, i, j-1, used, next[1:]) {
		return true
	}

	used[i][j] = false

	return false
}

func main() {
	r := exist(
		[][]byte{
			{'a'},
		},
		"a",
	)
	fmt.Println(r)
}
