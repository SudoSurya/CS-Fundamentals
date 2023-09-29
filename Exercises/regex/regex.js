let sentence = "Surya"
let regex = /SuRyA/i

let test = regex.test(sentence)
console.log(test)

let testStr = "Repeat, REpeat, Repeat"
let outRegex = /Repeat/ig
console.log(testStr.match(outRegex))


let exampleStr = "  lun iiun kun Let's have fun with regular expressions!  "
let wsRegex = /.un/ig
let result = exampleStr.match(wsRegex)
console.log(result)

// Path: regex.js
