export function dfs_traversal(V: number, adjacencyList: number[][]) {
    let visited: boolean[] = new Array(V).fill(false);
    let result: number[] = [];

    visited[0] = true;
    dfs(0, adjacencyList, visited, result);
    return result;
}

function dfs(node: number, adjacencyList: number[][], visited: boolean[], result: number[]) {
    result.push(node);
    visited[node] = true;

    for (let i = 0; i < adjacencyList[node].length; i++) {
        if (!visited[adjacencyList[node][i]]) {
            dfs(adjacencyList[node][i], adjacencyList, visited, result);
        }
    }
}

