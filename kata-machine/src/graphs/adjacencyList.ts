
export function constructAdjList(n: number, edges: number[][]): number[][] {
    const adjList: number[][] = new Array(n).fill(0).map(() => [])
    for (let i = 0; i < edges.length; i++) {
        if (edges[i].length == 0) {
            continue
        }
        for (let j = 0; j < edges[i].length; j++) {
            adjList[i].push(edges[i][j])
        }
    }
    return adjList
}
