const fs = require('fs');

fs.readFile('./input', 'utf8', (err, data) => {
    const res = data
        .split('\n\n')
        .map(block => block.split('\n').reduce((acc, cur) => acc + Number(cur), 0))
        .reduce((max, sum) => Math.max(max, sum), -1);

    console.log(res);
});
