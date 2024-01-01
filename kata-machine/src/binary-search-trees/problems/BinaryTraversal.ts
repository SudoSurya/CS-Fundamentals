class TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.left = (left === undefined ? null : left)
        this.right = (right === undefined ? null : right)
    }
}

function inorderTraversal(root: TreeNode | null): number[] {

    let result: number[] = [];
    traverse(root, result);
    return result;
};

function traverse(root: TreeNode | null, result: number[]): void {
    if (root === null) {
        return;
    }

    traverse(root.left, result);
    result.push(root.val);
    traverse(root.right, result);
}
