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

    delete(val: T, rootnode: Node<T>): Node<T> | undefined {
        if (rootnode == undefined) {
            return rootnode;
        } else if (val < rootnode.val) {
            rootnode.left = this.delete(val, rootnode.left as Node<T>);
            return rootnode;
        } else if (val > rootnode.val) {
            rootnode.right = this.delete(val, rootnode.right as Node<T>);
            return rootnode;
        } else if (val === rootnode.val) {
            if (rootnode.left == undefined) {
                return rootnode.right;
            } else if (rootnode.right == undefined) {
                return rootnode.left;
            } else {
                rootnode.right = this.lift(rootnode.right, rootnode);
                return rootnode;
            }
        }
    }
    lift(node: Node<T>, val: Node<T>): Node<T> {
        if (node.left) {
            node.left = this.lift(node.left, val);
            return node;
        }
        else {
            val.val = node.val;
            return node.right as Node<T>;
        }
    }
    traverse(node: Node<T>) {
        if(node == undefined){
            return;
        }
        this.traverse(node.left as Node<T>);
        console.log(node.val);
        this.traverse(node.right as Node<T>);
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
