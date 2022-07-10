# android-jni-go (ajni)

This library contains types and functions for interoperation Go and JVM on Android.

Most of the code is ported from [xlab/android-go](https://github.com/xlab/android-go), which seems not being maintained for a quite a while. This library extracts a small portion from android-go's auto-generated code, and fixes several critical issues.

## Differences from xlab/android-go

**This library contains only JNI interfaces.** Instead, xlab/android-go is a self-contained project to create Android app using Go.

**This library uses `uintptr` as underlying representation of JNI reference types,** while xlab/android-go uses `unsafe.Pointer`. The latter could be bugged when, for example, a JNI function returns a local ref. The address value of local ref is usually small (`<= 0xFF`) and could fail Go's pointer checking when goruntime moves the stack.

It's also worth noted that `xlab/android-go` uses `*JbyteArray` to represent a JVM reference type, while in this library we use simply `JbyteArray`. Apart from the difference, this library can be used as a drop-in replacement for xlab/android-go.
