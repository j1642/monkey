Interpreter and compiler for the Monkey programming language from [Writing an Interpreter in Go](https://interpreterbook.com/)
and [Writing a Compiler in Go](https://compilerbook.com/) by Thorsten Ball.

```
let fibonacci = fn(x) {
    if (x == 0) {
        return 0;
    } else {
        if (x == 1) {
            return 1;
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};
fibonacci(15);
```
