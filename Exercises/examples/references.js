function reference(array) {
    array.push(4);
    return array;
}

let myArray = [1, 2, 3];
let res = reference(myArray);
console.log(res); // [1, 2, 3, 4]
console.log(myArray); // [1, 2, 3, 4]

function noReference(array) {
    array = []
    return array;
}

let yArray = [1, 2, 3];

let res2 = noReference(yArray);

console.log(res2); // [1, 2, 3, 4]
console.log(yArray); // [1, 2, 3]
