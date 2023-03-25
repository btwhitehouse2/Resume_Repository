## General Comments 
We already talked via Ed about the structure for the LLVM IR translation so I don't have many comments. A few notes 

1. You may want to think about storing register information in the symbol table of the VarEntries because you will need this later on when you get to code generation. 
2. The Operand interface looks great; however, you could possibly also just define it as a struct where it has a flag that indicates the operand is either an immediate value or register value. You could also just keep the current state where it's just an interface and define distinct structs for Immediate and Register. 
3. Make sure to make a distinct Label struct where each time you call the constructor: "NewLabel" it will automatically use the label counter to make a new label that is unique. 
