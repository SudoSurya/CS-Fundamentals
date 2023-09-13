class Linter {
    constructor() {
        this.stack = [];
        this.opening = ['(', '[', '{'];
        this.closing = [')', ']', '}'];
    }

    lint(text) {
        for (let i = 0; i < text.length; i++) {
            if (this.#isOpening(text[i])) {
                this.stack.push(text[i]);
            } else if (this.#isClosing(text[i])) {
                let opening = this.stack.pop();
                if (!opening) {
                    return `Error! Expected ${this.opening[this.closing.indexOf(text[i])]} found end of input`;
                }
                if (!this.#isMatching(opening, text[i])) {
                    return `Error! Expected ${this.opening[this.closing.indexOf(text[i])]} found ${text[i]}`;
                }
            }
        }
        if (this.stack.length === 0) {
            return 'Success!';
        } else {
            return `Error! Expected
                ${this.opening[this.closing.indexOf(this.stack[this.stack.length - 1])]}
                found end of input`;
        }
    }

    #isOpening(char) {
        return this.opening.includes(char);
    }
    #isClosing(char) {
        return this.closing.includes(char);
    }
    #isMatching(opening, closing) {
        return this.opening.indexOf(opening) === this.closing.indexOf(closing);
    }

}

let linter = new Linter();

console.log(linter.lint('function add(a, b) ]{ return a + b; }'));
