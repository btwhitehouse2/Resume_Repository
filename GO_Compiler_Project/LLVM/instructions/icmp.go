package ir

//<result> = icmp <cond> <ty> <op1>, <op2>   ; yields i1 or <N x i1>:result

// <result> = icmp eq i32 4, 5          ; yields: result=false
// <result> = icmp ne ptr %X, %X        ; yields: result=false
// <result> = icmp ult i16  4, 5        ; yields: result=true