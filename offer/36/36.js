/**
 * // Definition for a Node.
 * function Node(val,left,right) {
 *    this.val = val;
 *    this.left = left;
 *    this.right = right;
 * };
 */
/**
 * @param {Node} root
 * @return {Node}
 */

function Node(val, left, right) {
    this.val = val;
    this.left = left;
    this.right = right;
};
let treeToDoublyList = function (root) {
    if (!root) {
        return;
    }
    let pre = head = null
    inorder(root)

    function inorder(node) {
        if (!node)
            return
        inorder(node.left)



        inorder(node.right)
    }
};



var buildTree = function (preorder, inorder) {
    if (preorder.length === 0) return null;
    const root = new Node(preorder[0]);
    const leftLength = inorder.indexOf(preorder[0]);

    root.left = buildTree(preorder.slice(1, leftLength + 1), inorder.slice(0, leftLength))
    root.right = buildTree(preorder.slice(leftLength + 1), inorder.slice(leftLength + 1))
    return root;
};

let root = buildTree([10, 6, 4, 8, 14, 12, 16], [4, 6, 8, 10, 12, 14, 16])
treeToDoublyList(root)