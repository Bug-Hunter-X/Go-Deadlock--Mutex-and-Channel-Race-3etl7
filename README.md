# Go Deadlock Example

This repository demonstrates a common deadlock scenario in Go involving mutexes and channels.

## Description
The `bug.go` file contains a program that creates two goroutines.  The first goroutine acquires a mutex, closes a channel, and then releases the mutex. The second goroutine waits on the channel, and after receiving the signal, it tries to acquire the same mutex.  This leads to a deadlock because goroutine 1 holds the mutex, and goroutine 2 is blocked waiting for it. 

## Solution
The `bugSolution.go` file provides a solution by carefully ordering the operations to avoid the deadlock. The solution demonstrates proper usage of mutexes and channels to prevent race conditions and deadlocks.