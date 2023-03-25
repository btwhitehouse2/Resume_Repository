package ir

//<result> = getelementptr <ty>, ptr <ptrval>{, [inrange] <ty> <idx>}*

// %aptr = getelementptr {i32, [12 x i8]}, ptr %saptr, i64 0, i32 1
// %vptr = getelementptr {i32, <2 x i8>}, ptr %svptr, i64 0, i32 1, i32 1
