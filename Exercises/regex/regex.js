/* let sentence = "Surya"
let regex = /SuRyA/i

let test = regex.test(sentence)
console.log(test)

let testStr = "Repeat, REpeat, Repeat"
let outRegex = /Repeat/ig
console.log(testStr.match(outRegex))


let exampleStr = "  lun iiun kun Let's have fun with regular expressions!  "
let wsRegex = /.un/ig
let result = exampleStr.match(wsRegex)
console.log(result) */

// Path: regex.js


function grep(str, regex) {
    let SpittedStr = str.split("\n")
    for (let i = 0; i < SpittedStr.length; i++) {
        if (regex.test(SpittedStr[i])) {
            console.log(SpittedStr[i])
        }
    }
}

let str = `
Ah Love! could you and I with Fate conspire\n
To grasp this sorry Scheme of Things entire,\n
    Would not we shatter it to bits-- and then\n
Re - mould it nearer to the Heart's Desire!\n
    `

let regex = /\b(it)\b/g
grep(str, regex)
