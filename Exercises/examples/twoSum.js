function twoSum(array) {
    for (let i = 0; i < array.length; i++) {
        for (let j = 0; j < array.length; j++) {
            if (i !== j && array[i] + array[j] === 10) {
                return true;
            }
        }
    }
    return false;

}

function contains(params) {
    for (let i = 0; i < params.length; i++) {
        if (params[i] === 'X') {
            return true;
        }
    }
    return false;
}

console.log(twoSum([1, 2, 3, 4, 5, 6]))

let str = 'XOX'
console.log(contains(str))
