package ir

//<result> = alloca [inalloca] <type> [, <ty> <NumElements>] [, align <alignment>] [, addrspace(<num>)]     ; yields type addrspace(num)*:result
//%ptr = alloca i32                             ; yields ptr
//%ptr = alloca i32, i32 4                      ; yields ptr
//%ptr = alloca i32, i32 4, align 1024          ; yields ptr
//%ptr = alloca i32, align 1024                 ; yields ptr

