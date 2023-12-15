export class Vertex<T>{
    val: T;
    adj: Vertex<T>[];
    constructor(val: T) {
        this.val = val;
        this.adj = [];
    }
    addEdge(vertex: Vertex<T>) {
        if (!this.adj.includes(vertex)) {
            this.adj.push(vertex);
            vertex.addEdge(this);
        }
    }
}

export function dfsTraversal<T>(vertex: Vertex<T>, visited: any = {}) {
    // Mark vertex as visited by adding it to the hash table:
    visited[vertex.val] = true;

    // Print the vertex's value, so we can make sure our traversal really works:
    console.log(vertex.val);

    // Iterate through the current vertex's adjacent vertices:
    for (const adjacentVertex of vertex.adj) {
        // Ignore an adjacent vertex if we've already visited it:
        if (!visited[adjacentVertex.val]) {
            // Recursively call this method on the adjacent vertex:
            dfsTraversal(adjacentVertex, visited);
        }
    }
}
