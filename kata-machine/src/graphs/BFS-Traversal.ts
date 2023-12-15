
export function bfs_traversal(list: number[][]): number[] {
    let visited = new Array(list.length).fill(false)
    let queue: number[] = []
    let result: number[] = []
    queue.push(0)

    while (queue.length > 0) {
        console.log("------------start of loop------------")
        console.log("queue",queue)
        console.log("result",result)
        console.log("visited",visited)
        let curr = queue.shift()
        if (visited[curr!]) {
            continue
        }
        if (curr == undefined) {
            break
        }
        result.push(curr)
        visited[curr!] = true
        for (let i = 0; i < list[curr!].length; i++) {
            queue.push(list[curr!][i])
        }

        console.log("------------mid of loop------------")
        console.log("queue",queue)
        console.log("result",result)
        console.log("visited",visited)
        console.log("------------end of loop------------")
    }
    return result
}



