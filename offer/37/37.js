function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

var serialize = function (root) {
    if (!root) return '';
    let q = [root];
    let res = []

    while (q.length) {
        let node = q.shift();
        if (node){
            q.push(node.left)
            q.push(node.right)
            res.push(node.val)
        }else {
            res.push('null')
        }
    }
    return res.join(",")
};

/**
 * Decodes your encoded data to tree.
 *
 * @param {string} data
 * @return {TreeNode}
 */
var deserialize = function (data) {
    if (data==='') return null
    let nodeVal = data.split(",").map(val => {
        let v = parseInt(val)
        return isNaN(v) ? null : v
    });
    let i = 1
    let root = new TreeNode(nodeVal[0]);
    let q = [root]

    while (q.length) {
        let node = q.shift();
        if (nodeVal[i] !== null) {
            node.left = new TreeNode(nodeVal[i]);
            q.push(node.left)
        }
        i++;
        if (nodeVal[i] !== null) {
            node.right = new TreeNode(nodeVal[i]);
            q.push(node.right)
        }
        i++;
    }

    return root
};

console.log(serialize(deserialize("1,2,3,null,null,4,5")))