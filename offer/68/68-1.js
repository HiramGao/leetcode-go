var lowestCommonAncestor = function(root, p, q) {
    let res = root
    while(true){
        if(p.val < res.val && q.val < res.val){
            res = res.left
        }else if (p.val > res.val && q.val > res.val){
            res = res.right
        }else {
            break
        }
    }
    return res
};


