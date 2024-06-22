
export function constructAdjMatrix(n: number, edges: number[][]): number[][] {
    const adjMatrix: number[][] = new Array(n).fill(0).map(() => new Array(n).fill(0))
    for (let i = 0; i < edges.length; i++) {
        const [src, dest] = edges[i]
        adjMatrix[src][dest] = 1
        adjMatrix[dest][src] = 1
    }
    return adjMatrix
}
