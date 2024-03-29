package main

/**
* @project: go-leetcode
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/19/21 9:26 AM
**/

// 给定一个二叉树和一个节点，找到中序遍历的下一个节点。
// 如果这个节点有右子树，那么中序遍历的下一个节点就是它右子树的最左节点；
// 如果这个节点没有右子树，这里又分两种情况：1、它是它父节点的左子节点，那么中序遍历的下一个节点是它父节点；
// 2、不是第一种情况，需要从它父节点开始依次查找祖先节点，直到发现这个节点位于某个祖先节点的左子树上，这个祖先节点就是中序遍历的下一个节点。
