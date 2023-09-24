
function qs(arr, left, right) {
    if (left >= right) {
        return
    }
    let partitionIdx = partition(arr, left, right);
    qs(arr, left, partitionIdx - 1);
    qs(arr, partitionIdx + 1, right);
}
function quickSelect(arr, left, right, k) {
    if (right - left <= 0) {
        return arr[left];
    }

    let partitionIdx = partition(arr, left, right);
    if (partitionIdx === k) {
        return arr[partitionIdx];
    } else if (partitionIdx > k) {
        return quickSelect(arr, left, partitionIdx - 1, k);
    } else {
        return quickSelect(arr, partitionIdx + 1, right, k);
    }
}

function partition(arr, left, right) {
    let pivot = arr[right];
    let idx = left - 1;

    for (let i = left; i < right; i++) {
        if (arr[i] < pivot) {
            idx++;
            let temp = arr[i];
            arr[i] = arr[idx];
            arr[idx] = temp;
        }
    }

    idx++;
    arr[right] = arr[idx];
    arr[idx] = pivot;
    return idx;
}


let arr = [5, 4, 3, 2, 1];
console.log(arr);
qs(arr, 0, arr.length - 1);
console.log(quickSelect(arr, 0, arr.length - 1, 1));
console.log(arr);
