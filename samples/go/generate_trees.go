package main

import (
	. "./list"
)

type Integer int

var n int

func copyMap(m map[int]struct{}) map[int]struct{} {
	newMap := make(map[int]struct{}, n)
	for k := range m {
		newMap[k] = struct{}{}
	}
	return newMap
}

func fill(used map[int]struct{}, currentTree []*Integer) [][]*Integer {
	results := make([][]*Integer, 0, n)
	if currentTree[(len(currentTree)-1)/2] != nil {
		parentVal := int(*currentTree[(len(currentTree)-1)/2])
		if len(currentTree)%2 == 0 { // left
			for i := 0; i < parentVal; i++ {
				if _, ok := used[i]; !ok {
					newTree := make([]*Integer, 0)
					newItem := Integer(i)
					copy(newTree, append(currentTree, &newItem))
					results = append(results, newTree)
				}
			}
		} else { // right
			for i := parentVal + 1; i <= n; i++ {
				if _, ok := used[i]; !ok {
					newTree := make([]*Integer, 0)
					newItem := Integer(i)
					copy(newTree, append(currentTree, &newItem))
					results = append(results, newTree)
				}
			}
		}
	}
	newTree := make([]*Integer, 0)
	copy(newTree, append(currentTree, nil))
	results = append(results, newTree)
	if len(used)+1 == n {
		return results
	}
	finalResults := make([][]*Integer, 0, len(results))
	for treeIdx := range results {
		newUsed := copyMap(used)
		if len(results[treeIdx]) != 0 {
			lastVal := results[treeIdx][len(results[treeIdx])-1]
			newUsed[int(*lastVal)] = struct{}{}
			finalResults = append(finalResults, fill(newUsed, results[treeIdx])...)
		}
	}
	return finalResults
}

type ParentWithSide struct {
	parent *TreeNode
	left   bool
}

func buildTree(tree []*Integer) *TreeNode {
	head := &TreeNode{}
	toServe := make([]*TreeNode, 0)
	toServe = append(toServe, head)
	parentMap := make(map[*TreeNode]ParentWithSide)
	for i := range tree {
		if tree[i] != nil {
			toServe[0].Val = int(*tree[i])
		} else {
			parent, ok := parentMap[toServe[0]]
			if ok {
				if parent.left {
					parent.parent.Left = nil
				} else {
					parent.parent.Right = nil
				}
			}
		}
		toServe[0].Left = &TreeNode{}
		toServe[0].Right = &TreeNode{}
		parentMap[toServe[0].Left] = ParentWithSide{toServe[0], true}
		parentMap[toServe[0].Right] = ParentWithSide{toServe[0], false}
		toServe = append(toServe, toServe[0].Left, toServe[0].Right)
		toServe = toServe[1:]
	}
	for i := range toServe {
		parent := parentMap[toServe[i]]
		if parent.left {
			parent.parent.Left = nil
		} else {
			parent.parent.Right = nil
		}
	}
	return head
}

func getTrees(trees [][]*Integer) []*TreeNode {
	heads := make([]*TreeNode, 0, len(trees))
	for treeIdx := range trees {
		heads = append(heads, buildTree(trees[treeIdx]))
	}
	return heads
}

func generateTrees(n int) []*TreeNode {
	n = n
	results := make([]*TreeNode, 0, n)
	for i := 0; i < n; i++ {
		used := make(map[int]struct{}, n)
		used[i] = struct{}{}
		newElem := Integer(i)
		results = append(results, getTrees(fill(used, []*Integer{&newElem}))...)
	}
	return results
}
