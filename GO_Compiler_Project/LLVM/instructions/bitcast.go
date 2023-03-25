package ir

//convert an object to a different type 
//<result> = bitcast <ty> <value> to <ty2>             ; yields ty2

// %X = bitcast i8 255 to i8          ; yields i8 :-1
// %Y = bitcast i32* %x to i16*       ; yields i16*:%x