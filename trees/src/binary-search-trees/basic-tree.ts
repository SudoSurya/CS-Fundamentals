export type Node<T> = {
    val: T;
    left: Node<T> | undefined;
    right: Node<T> | undefined;
}
export class TreeNode<T>{
    val: T;
    left?: Node<T>;
    right?: Node<T>;
    constructor(val: T, left?: Node<T>, right?: Node<T>) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
    insert(val: T, rootnode: Node<T>) {
        if (val < rootnode.val) {
            if (rootnode.left == undefined) {
                rootnode.left = new TreeNode(val) as Node<T>;
            } else {
                this.insert(val, rootnode.left);
            }
        } else if (val > rootnode.val) {
            if (rootnode.right == undefined) {
                rootnode.right = new TreeNode(val) as Node<T>;
            } else {
                this.insert(val, rootnode.right);
            }
        }
    }
}

export function search<T>(val: T, rootnode: Node<T>): Node<T> | undefined {
    if (rootnode?.val === val) {
        return rootnode;
    } else if (val < rootnode?.val) {
        return search(val, rootnode?.left as Node<T>);
    } else if (val > rootnode?.val) {
        return search(val, rootnode?.right as Node<T>);
    } else {
        return undefined;
    }
}
