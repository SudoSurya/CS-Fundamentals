
function isSubset(array1, array2) {
    let largerArray = array1.length > array2.length ? array1 : array2;
    let smallerArray = array1.length > array2.length ? array2 : array1;

    let largerArrayHash = {};

    for (const value of largerArray) {
        largerArrayHash[value] = true;
    }

    for (const value of smallerArray) {
        if (!largerArrayHash[value]) {
            return false;
        }
    }
    return true;
}

console.log(isSubset([1, 2, 3, 4, 5], [1, 4, 5])); // true
console.log(isSubset([1, 2, 3, 4, 5], [1, 6])); // false
