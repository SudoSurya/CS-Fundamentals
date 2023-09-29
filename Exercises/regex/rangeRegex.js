let quoteSample = "The five boxing wizards jump quickly."
let alphabetRegexV2 = /[a-z]/ig

let result = alphabetRegexV2.test(quoteSample)
console.log(result)

let RangeoFNumbers = ""
let myRegex = /[^0-9aeiou ]/ig

let result2 = RangeoFNumbers.match(myRegex)
let testREs = myRegex.test(RangeoFNumbers)
console.log(testREs)
console.log(result2)

let password = "abcA"

let checkPass = /t[a-z]*?i/

let result3 = checkPass.test(password)
console.log(result3)

let StringPassword = "abc123"

let checkPass2 = /(?=\w{3,6})(?=\D*\d)/
