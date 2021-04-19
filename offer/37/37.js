var serialize = function (root) {
    if (!root) return '';
    let q = [root];
    let res = []

    while (q.length) {
        let node = q.splice(0, 1)
        console.log(node)
        node.left && node.push(node.left)
        node.right && node.push(node.right)
        res.push(node.val)
    }
    console.log(res.join(","))
    return res.join(",")
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
var deserialize = function (data) {
    console.log(data)
};

console.log(deserialize("[1,2,3,null,null,4,5]"))