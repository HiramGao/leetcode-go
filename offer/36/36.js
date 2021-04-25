/**
 * // Definition for a Node.
 * function Node(val,left,right) {
 *    this.val = val;
 *    this.left = left;
 *    this.right = right;
 * };
 **/
let pre, head
let treeToDoublyList = function (root) {
    if (!root) {
        return null;
    }
    dfs(root);
    head.left = pre;
    pre.right = head;
    return head;
};

function dfs(treeNode) {
    if (!treeNode) return;
    dfs(treeNode.left);
    if (pre){
        pre.right = treeNode;
    }
    else{
        head = treeNode;
    }
    treeNode.left = pre;
    pre =treeNode;
    dfs(treeNode.right);
}

