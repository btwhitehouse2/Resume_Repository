package ir

// <result> = [tail | musttail | notail ] call [fast-math flags] [cconv] [ret attrs] [addrspace(<num>)]
//            <ty>|<fnty> <fnptrval>(<function args>) [fn attrs] [ operand bundles ]

// %retval = call i32 @test(i32 %argc)
// call i32 (ptr, ...) @printf(ptr %msg, i32 12, i8 42)        ; yields i32
// %X = tail call i32 @foo()                                    ; yields i32
// %Y = tail call fastcc i32 @foo()  ; yields i32
// call void %foo(i8 signext 97)