package ir
//branch to another location
//br i1 <cond>, label <iftrue>, label <iffalse>

// Test:
//   %cond = icmp eq i32 %a, %b
//   br i1 %cond, label %IfEqual, label %IfUnequal
// IfEqual:
//   ret i32 1
// IfUnequal:
//   ret i32 0