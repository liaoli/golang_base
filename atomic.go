package main

/**
*@author: 廖理
*@date:2022/8/11
**/
//

//原子操作
//针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，
//通常直接使用原子操作比使用锁操作效率更高。
//Go语言中原子操作由内置的标准库sync/atomic提供。
//
//atomic包
//方法	解释
//func LoadInt32(addr *int32) (val int32)
//func LoadInt64(addr *int64) (val int64)
//func LoadUint32(addr *uint32) (val uint32)
//func LoadUint64(addr *uint64) (val uint64)
//func LoadUintptr(addr *uintptr) (val uintptr)
//func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)	读取操作
//func StoreInt32(addr *int32, val int32)
//func StoreInt64(addr *int64, val int64)
//func StoreUint32(addr *uint32, val uint32)
//func StoreUint64(addr *uint64, val uint64)
//func StoreUintptr(addr *uintptr, val uintptr)
//func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)	写入操作
//func AddInt32(addr *int32, delta int32) (new int32)
//func AddInt64(addr *int64, delta int64) (new int64)
//func AddUint32(addr *uint32, delta uint32) (new uint32)
//func AddUint64(addr *uint64, delta uint64) (new uint64)
//func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)	修改操作
//func SwapInt32(addr *int32, new int32) (old int32)
//func SwapInt64(addr *int64, new int64) (old int64)
//func SwapUint32(addr *uint32, new uint32) (old uint32)
//func SwapUint64(addr *uint64, new uint64) (old uint64)
//func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
//func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)	交换操作
//func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
//func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
//func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
//func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
//func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
//func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)	比较并交换操作
