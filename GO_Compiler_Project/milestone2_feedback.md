# Feedback for Milestone 2
Please check Lamont's comment on your compiler. Please feel free to reach out and come to office hours or set up a time if you have any more questions regarding implementations of your compiler, or if you want help with reviewing your code<br>
Below are some of common problems I've identified, since your code does not compile for this phase I did not get the chance to test them out, but below are some common mistakes to look out for when you're completing semantic analysis. 

1. You may want to check for the existance of a main function, since it is used as a entry point to the execution of your code. 
2. Redeclaration of global struct/variables: <br>
`type Point2D struct { ... }` <br>
`var Point2D int;` <br>
This should produce an error because type declarations and global variables have global scope so there cannot be a global variable declaration and type decalaration of the same scope. Check "Lanaguage Overview" page under "Language Semantics".
3. A control-flow check for all functions, you could do this by adding a return check function for the statements and check the last statement of a function for return statements (unless function is supposed to return nil), but exact implementation is up to you. Just make sure that it is implemented so that `Any function with a return type must return a valid value (of its return type) along all control flow paths.` as per language specification (also it'll be difficult to check dyanmically), meaning that the following functions should be considered erroneous.
```
func bar (num int) int {
    if (num == 0) {
        return 1;
    } else {
        if (num >= 1) {
            return 1;
        } else {
            // This inner else block is missing a return statement.
        }
    }
}

func foo (num int) int {
    // A for loop is not neccessarily executed, so this return statement is not guarenteed.
    for(false){
        return 0;
    }
}
```