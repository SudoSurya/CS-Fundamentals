class stack {
    #items = [];
    constructor() {
        this.items = [];
    }

    push(element) {
        this.items.push(element);
        return "Pushed"
    }

    pop() {
        if (this.items.length == 0) return "Underflow";
        return this.items.pop();
    }
    length() {
        return this.items.length;
    }
    peek() {
        return this.items[this.items.length - 1];
    }

}
