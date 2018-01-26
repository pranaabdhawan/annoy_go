/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: src/annoygomodule.i

package annoyindex

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _goslice_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _goslice_ swig_type_4;
typedef _goslice_ swig_type_5;
typedef _goslice_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef _goslice_ swig_type_9;
typedef _goslice_ swig_type_10;
extern void _wrap_Swig_free_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_annoyindex_acf5c19c77666e85(swig_intgo arg1);
extern void _wrap_delete_AnnoyIndex_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern void _wrap_AnnoyIndex_addItem_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_intgo arg2, swig_type_1 arg3);
extern void _wrap_AnnoyIndex_build_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_intgo arg2);
extern _Bool _wrap_AnnoyIndex_save_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_type_2 arg2);
extern void _wrap_AnnoyIndex_unload_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern _Bool _wrap_AnnoyIndex_load_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_type_3 arg2);
extern float _wrap_AnnoyIndex_getDistance_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern void _wrap_AnnoyIndex_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5, swig_voidp arg6);
extern void _wrap_AnnoyIndex_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_type_4 arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5, swig_voidp arg6);
extern void _wrap_AnnoyIndex_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndex_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_type_5 arg2, swig_intgo arg3, swig_intgo arg4, swig_voidp arg5);
extern swig_intgo _wrap_AnnoyIndex_getNItems_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern void _wrap_AnnoyIndex_verbose_annoyindex_acf5c19c77666e85(uintptr_t arg1, _Bool arg2);
extern void _wrap_AnnoyIndex_getItem_annoyindex_acf5c19c77666e85(uintptr_t arg1, swig_intgo arg2, swig_voidp arg3);
extern uintptr_t _wrap_new_AnnoyIndex_annoyindex_acf5c19c77666e85(void);
extern uintptr_t _wrap_new_AnnoyIndexAngular_annoyindex_acf5c19c77666e85(swig_intgo arg1);
extern void _wrap_delete_AnnoyIndexAngular_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern void _wrap_AnnoyIndexAngular_addItem_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_type_6 arg2);
extern void _wrap_AnnoyIndexAngular_build_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1);
extern _Bool _wrap_AnnoyIndexAngular_save_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_7 arg1);
extern void _wrap_AnnoyIndexAngular_unload_annoyindex_acf5c19c77666e85(uintptr_t _swig_base);
extern _Bool _wrap_AnnoyIndexAngular_load_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_8 arg1);
extern float _wrap_AnnoyIndexAngular_getDistance_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2);
extern void _wrap_AnnoyIndexAngular_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndexAngular_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4);
extern void _wrap_AnnoyIndexAngular_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_9 arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndexAngular_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_10 arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4);
extern swig_intgo _wrap_AnnoyIndexAngular_getNItems_annoyindex_acf5c19c77666e85(uintptr_t _swig_base);
extern void _wrap_AnnoyIndexAngular_verbose_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, _Bool arg1);
extern void _wrap_AnnoyIndexAngular_getItem_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_voidp arg2);
extern uintptr_t _wrap_new_AnnoyIndexEuclidean_annoyindex_acf5c19c77666e85(swig_intgo arg1);
extern void _wrap_delete_AnnoyIndexEuclidean_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern void _wrap_AnnoyIndexEuclidean_addItem_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_type_6 arg2);
extern void _wrap_AnnoyIndexEuclidean_build_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1);
extern _Bool _wrap_AnnoyIndexEuclidean_save_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_7 arg1);
extern void _wrap_AnnoyIndexEuclidean_unload_annoyindex_acf5c19c77666e85(uintptr_t _swig_base);
extern _Bool _wrap_AnnoyIndexEuclidean_load_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_8 arg1);
extern float _wrap_AnnoyIndexEuclidean_getDistance_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2);
extern void _wrap_AnnoyIndexEuclidean_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndexEuclidean_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4);
extern void _wrap_AnnoyIndexEuclidean_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_9 arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndexEuclidean_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_10 arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4);
extern swig_intgo _wrap_AnnoyIndexEuclidean_getNItems_annoyindex_acf5c19c77666e85(uintptr_t _swig_base);
extern void _wrap_AnnoyIndexEuclidean_verbose_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, _Bool arg1);
extern void _wrap_AnnoyIndexEuclidean_getItem_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_voidp arg2);
extern uintptr_t _wrap_new_AnnoyIndexManhattan_annoyindex_acf5c19c77666e85(swig_intgo arg1);
extern void _wrap_delete_AnnoyIndexManhattan_annoyindex_acf5c19c77666e85(uintptr_t arg1);
extern void _wrap_AnnoyIndexManhattan_addItem_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_type_6 arg2);
extern void _wrap_AnnoyIndexManhattan_build_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1);
extern _Bool _wrap_AnnoyIndexManhattan_save_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_7 arg1);
extern void _wrap_AnnoyIndexManhattan_unload_annoyindex_acf5c19c77666e85(uintptr_t _swig_base);
extern _Bool _wrap_AnnoyIndexManhattan_load_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_8 arg1);
extern float _wrap_AnnoyIndexManhattan_getDistance_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2);
extern void _wrap_AnnoyIndexManhattan_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndexManhattan_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4);
extern void _wrap_AnnoyIndexManhattan_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_9 arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4, swig_voidp arg5);
extern void _wrap_AnnoyIndexManhattan_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_type_10 arg1, swig_intgo arg2, swig_intgo arg3, swig_voidp arg4);
extern swig_intgo _wrap_AnnoyIndexManhattan_getNItems_annoyindex_acf5c19c77666e85(uintptr_t _swig_base);
extern void _wrap_AnnoyIndexManhattan_verbose_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, _Bool arg1);
extern void _wrap_AnnoyIndexManhattan_getItem_annoyindex_acf5c19c77666e85(uintptr_t _swig_base, swig_intgo arg1, swig_voidp arg2);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_annoyindex_acf5c19c77666e85(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrAnnoyIndex uintptr

func (p SwigcptrAnnoyIndex) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAnnoyIndex) SwigIsAnnoyIndex() {
}

func DeleteAnnoyIndex(arg1 AnnoyIndex) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_AnnoyIndex_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrAnnoyIndex) AddItem(arg2 int, arg3 []float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_AnnoyIndex_addItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), *(*C.swig_type_1)(unsafe.Pointer(&_swig_i_2)))
	if Swig_escape_always_false {
		Swig_escape_val = arg3
	}
}

func (arg1 SwigcptrAnnoyIndex) Build(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndex_build_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrAnnoyIndex) Save(arg2 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_AnnoyIndex_save_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrAnnoyIndex) Unload() {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndex_unload_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrAnnoyIndex) Load(arg2 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (bool)(C._wrap_AnnoyIndex_load_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrAnnoyIndex) GetDistance(arg2 int, arg3 int) (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (float32)(C._wrap_AnnoyIndex_getDistance_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrAnnoyIndex) GetNnsByItem__SWIG_0(arg2 int, arg3 int, arg4 int, arg5 *[]int, arg6 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	C._wrap_AnnoyIndex_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4), C.swig_voidp(_swig_i_5))
}

func (arg1 SwigcptrAnnoyIndex) GetNnsByVector__SWIG_0(arg2 []float32, arg3 int, arg4 int, arg5 *[]int, arg6 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	_swig_i_5 := arg6
	C._wrap_AnnoyIndex_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4), C.swig_voidp(_swig_i_5))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrAnnoyIndex) GetNnsByItem__SWIG_1(arg2 int, arg3 int, arg4 int, arg5 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndex_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4))
}

func (p SwigcptrAnnoyIndex) GetNnsByItem(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByItem__SWIG_1(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByItem__SWIG_0(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrAnnoyIndex) GetNnsByVector__SWIG_1(arg2 []float32, arg3 int, arg4 int, arg5 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndex_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), *(*C.swig_type_5)(unsafe.Pointer(&_swig_i_1)), C.swig_intgo(_swig_i_2), C.swig_intgo(_swig_i_3), C.swig_voidp(_swig_i_4))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (p SwigcptrAnnoyIndex) GetNnsByVector(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByVector__SWIG_1(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByVector__SWIG_0(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrAnnoyIndex) GetNItems() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_AnnoyIndex_getNItems_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrAnnoyIndex) Verbose(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndex_verbose_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrAnnoyIndex) GetItem(arg2 int, arg3 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_AnnoyIndex_getItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_voidp(_swig_i_2))
}

func NewAnnoyIndex() (_swig_ret AnnoyIndex) {
	var swig_r AnnoyIndex
	swig_r = (AnnoyIndex)(SwigcptrAnnoyIndex(C._wrap_new_AnnoyIndex_annoyindex_acf5c19c77666e85()))
	return swig_r
}

type AnnoyIndex interface {
	Swigcptr() uintptr
	SwigIsAnnoyIndex()
	AddItem(arg2 int, arg3 []float32)
	Build(arg2 int)
	Save(arg2 string) (_swig_ret bool)
	Unload()
	Load(arg2 string) (_swig_ret bool)
	GetDistance(arg2 int, arg3 int) (_swig_ret float32)
	GetNnsByItem(a ...interface{})
	GetNnsByVector(a ...interface{})
	GetNItems() (_swig_ret int)
	Verbose(arg2 bool)
	GetItem(arg2 int, arg3 *[]float32)
}

type SwigcptrAnnoyIndexAngular uintptr

func (p SwigcptrAnnoyIndexAngular) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAnnoyIndexAngular) SwigIsAnnoyIndexAngular() {
}

func NewAnnoyIndexAngular(arg1 int) (_swig_ret AnnoyIndexAngular) {
	var swig_r AnnoyIndexAngular
	_swig_i_0 := arg1
	swig_r = (AnnoyIndexAngular)(SwigcptrAnnoyIndexAngular(C._wrap_new_AnnoyIndexAngular_annoyindex_acf5c19c77666e85(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func DeleteAnnoyIndexAngular(arg1 AnnoyIndexAngular) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_AnnoyIndexAngular_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexAngular) AddItem(arg1 int, arg2 []float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndexAngular_addItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (_swig_base SwigcptrAnnoyIndexAngular) Build(arg1 int) {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndexAngular_build_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexAngular) Save(arg1 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AnnoyIndexAngular_save_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexAngular) Unload() {
	C._wrap_AnnoyIndexAngular_unload_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base))
}

func (_swig_base SwigcptrAnnoyIndexAngular) Load(arg1 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AnnoyIndexAngular_load_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_8)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetDistance(arg1 int, arg2 int) (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float32)(C._wrap_AnnoyIndexAngular_getDistance_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetNnsByItem__SWIG_0(arg1 int, arg2 int, arg3 int, arg4 *[]int, arg5 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndexAngular_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3), C.swig_voidp(_swig_i_4))
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetNnsByItem__SWIG_1(arg1 int, arg2 int, arg3 int, arg4 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_AnnoyIndexAngular_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3))
}

func (p SwigcptrAnnoyIndexAngular) GetNnsByItem(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByItem__SWIG_1(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByItem__SWIG_0(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetNnsByVector__SWIG_0(arg1 []float32, arg2 int, arg3 int, arg4 *[]int, arg5 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndexAngular_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3), C.swig_voidp(_swig_i_4))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetNnsByVector__SWIG_1(arg1 []float32, arg2 int, arg3 int, arg4 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_AnnoyIndexAngular_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_10)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func (p SwigcptrAnnoyIndexAngular) GetNnsByVector(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByVector__SWIG_1(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByVector__SWIG_0(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetNItems() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_AnnoyIndexAngular_getNItems_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base)))
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexAngular) Verbose(arg1 bool) {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndexAngular_verbose_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C._Bool(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexAngular) GetItem(arg1 int, arg2 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndexAngular_getItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (p SwigcptrAnnoyIndexAngular) SwigIsAnnoyIndex() {
}

func (p SwigcptrAnnoyIndexAngular) SwigGetAnnoyIndex() AnnoyIndex {
	return SwigcptrAnnoyIndex(p.Swigcptr())
}

type AnnoyIndexAngular interface {
	Swigcptr() uintptr
	SwigIsAnnoyIndexAngular()
	AddItem(arg1 int, arg2 []float32)
	Build(arg1 int)
	Save(arg1 string) (_swig_ret bool)
	Unload()
	Load(arg1 string) (_swig_ret bool)
	GetDistance(arg1 int, arg2 int) (_swig_ret float32)
	GetNnsByItem(a ...interface{})
	GetNnsByVector(a ...interface{})
	GetNItems() (_swig_ret int)
	Verbose(arg1 bool)
	GetItem(arg1 int, arg2 *[]float32)
	SwigIsAnnoyIndex()
	SwigGetAnnoyIndex() AnnoyIndex
}

type SwigcptrAnnoyIndexEuclidean uintptr

func (p SwigcptrAnnoyIndexEuclidean) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAnnoyIndexEuclidean) SwigIsAnnoyIndexEuclidean() {
}

func NewAnnoyIndexEuclidean(arg1 int) (_swig_ret AnnoyIndexEuclidean) {
	var swig_r AnnoyIndexEuclidean
	_swig_i_0 := arg1
	swig_r = (AnnoyIndexEuclidean)(SwigcptrAnnoyIndexEuclidean(C._wrap_new_AnnoyIndexEuclidean_annoyindex_acf5c19c77666e85(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func DeleteAnnoyIndexEuclidean(arg1 AnnoyIndexEuclidean) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_AnnoyIndexEuclidean_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) AddItem(arg1 int, arg2 []float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndexEuclidean_addItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) Build(arg1 int) {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndexEuclidean_build_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) Save(arg1 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AnnoyIndexEuclidean_save_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) Unload() {
	C._wrap_AnnoyIndexEuclidean_unload_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base))
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) Load(arg1 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AnnoyIndexEuclidean_load_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_8)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetDistance(arg1 int, arg2 int) (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float32)(C._wrap_AnnoyIndexEuclidean_getDistance_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetNnsByItem__SWIG_0(arg1 int, arg2 int, arg3 int, arg4 *[]int, arg5 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndexEuclidean_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3), C.swig_voidp(_swig_i_4))
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetNnsByItem__SWIG_1(arg1 int, arg2 int, arg3 int, arg4 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_AnnoyIndexEuclidean_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3))
}

func (p SwigcptrAnnoyIndexEuclidean) GetNnsByItem(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByItem__SWIG_1(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByItem__SWIG_0(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetNnsByVector__SWIG_0(arg1 []float32, arg2 int, arg3 int, arg4 *[]int, arg5 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndexEuclidean_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3), C.swig_voidp(_swig_i_4))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetNnsByVector__SWIG_1(arg1 []float32, arg2 int, arg3 int, arg4 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_AnnoyIndexEuclidean_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_10)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func (p SwigcptrAnnoyIndexEuclidean) GetNnsByVector(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByVector__SWIG_1(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByVector__SWIG_0(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetNItems() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_AnnoyIndexEuclidean_getNItems_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base)))
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) Verbose(arg1 bool) {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndexEuclidean_verbose_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C._Bool(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexEuclidean) GetItem(arg1 int, arg2 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndexEuclidean_getItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (p SwigcptrAnnoyIndexEuclidean) SwigIsAnnoyIndex() {
}

func (p SwigcptrAnnoyIndexEuclidean) SwigGetAnnoyIndex() AnnoyIndex {
	return SwigcptrAnnoyIndex(p.Swigcptr())
}

type AnnoyIndexEuclidean interface {
	Swigcptr() uintptr
	SwigIsAnnoyIndexEuclidean()
	AddItem(arg1 int, arg2 []float32)
	Build(arg1 int)
	Save(arg1 string) (_swig_ret bool)
	Unload()
	Load(arg1 string) (_swig_ret bool)
	GetDistance(arg1 int, arg2 int) (_swig_ret float32)
	GetNnsByItem(a ...interface{})
	GetNnsByVector(a ...interface{})
	GetNItems() (_swig_ret int)
	Verbose(arg1 bool)
	GetItem(arg1 int, arg2 *[]float32)
	SwigIsAnnoyIndex()
	SwigGetAnnoyIndex() AnnoyIndex
}

type SwigcptrAnnoyIndexManhattan uintptr

func (p SwigcptrAnnoyIndexManhattan) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrAnnoyIndexManhattan) SwigIsAnnoyIndexManhattan() {
}

func NewAnnoyIndexManhattan(arg1 int) (_swig_ret AnnoyIndexManhattan) {
	var swig_r AnnoyIndexManhattan
	_swig_i_0 := arg1
	swig_r = (AnnoyIndexManhattan)(SwigcptrAnnoyIndexManhattan(C._wrap_new_AnnoyIndexManhattan_annoyindex_acf5c19c77666e85(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func DeleteAnnoyIndexManhattan(arg1 AnnoyIndexManhattan) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_AnnoyIndexManhattan_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexManhattan) AddItem(arg1 int, arg2 []float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndexManhattan_addItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (_swig_base SwigcptrAnnoyIndexManhattan) Build(arg1 int) {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndexManhattan_build_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexManhattan) Save(arg1 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AnnoyIndexManhattan_save_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_7)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexManhattan) Unload() {
	C._wrap_AnnoyIndexManhattan_unload_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base))
}

func (_swig_base SwigcptrAnnoyIndexManhattan) Load(arg1 string) (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_AnnoyIndexManhattan_load_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_8)(unsafe.Pointer(&_swig_i_0))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetDistance(arg1 int, arg2 int) (_swig_ret float32) {
	var swig_r float32
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (float32)(C._wrap_AnnoyIndexManhattan_getDistance_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetNnsByItem__SWIG_0(arg1 int, arg2 int, arg3 int, arg4 *[]int, arg5 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndexManhattan_getNnsByItem__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3), C.swig_voidp(_swig_i_4))
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetNnsByItem__SWIG_1(arg1 int, arg2 int, arg3 int, arg4 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_AnnoyIndexManhattan_getNnsByItem__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3))
}

func (p SwigcptrAnnoyIndexManhattan) GetNnsByItem(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByItem__SWIG_1(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByItem__SWIG_0(a[0].(int), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetNnsByVector__SWIG_0(arg1 []float32, arg2 int, arg3 int, arg4 *[]int, arg5 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	_swig_i_4 := arg5
	C._wrap_AnnoyIndexManhattan_getNnsByVector__SWIG_0_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_9)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3), C.swig_voidp(_swig_i_4))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetNnsByVector__SWIG_1(arg1 []float32, arg2 int, arg3 int, arg4 *[]int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	_swig_i_3 := arg4
	C._wrap_AnnoyIndexManhattan_getNnsByVector__SWIG_1_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), *(*C.swig_type_10)(unsafe.Pointer(&_swig_i_0)), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2), C.swig_voidp(_swig_i_3))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
}

func (p SwigcptrAnnoyIndexManhattan) GetNnsByVector(a ...interface{}) {
	argc := len(a)
	if argc == 4 {
		p.GetNnsByVector__SWIG_1(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int))
		return
	}
	if argc == 5 {
		p.GetNnsByVector__SWIG_0(a[0].([]float32), a[1].(int), a[2].(int), a[3].(*[]int), a[4].(*[]float32))
		return
	}
	panic("No match for overloaded function call")
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetNItems() (_swig_ret int) {
	var swig_r int
	swig_r = (int)(C._wrap_AnnoyIndexManhattan_getNItems_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base)))
	return swig_r
}

func (_swig_base SwigcptrAnnoyIndexManhattan) Verbose(arg1 bool) {
	_swig_i_0 := arg1
	C._wrap_AnnoyIndexManhattan_verbose_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C._Bool(_swig_i_0))
}

func (_swig_base SwigcptrAnnoyIndexManhattan) GetItem(arg1 int, arg2 *[]float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_AnnoyIndexManhattan_getItem_annoyindex_acf5c19c77666e85(C.uintptr_t(_swig_base), C.swig_intgo(_swig_i_0), C.swig_voidp(_swig_i_1))
}

func (p SwigcptrAnnoyIndexManhattan) SwigIsAnnoyIndex() {
}

func (p SwigcptrAnnoyIndexManhattan) SwigGetAnnoyIndex() AnnoyIndex {
	return SwigcptrAnnoyIndex(p.Swigcptr())
}

type AnnoyIndexManhattan interface {
	Swigcptr() uintptr
	SwigIsAnnoyIndexManhattan()
	AddItem(arg1 int, arg2 []float32)
	Build(arg1 int)
	Save(arg1 string) (_swig_ret bool)
	Unload()
	Load(arg1 string) (_swig_ret bool)
	GetDistance(arg1 int, arg2 int) (_swig_ret float32)
	GetNnsByItem(a ...interface{})
	GetNnsByVector(a ...interface{})
	GetNItems() (_swig_ret int)
	Verbose(arg1 bool)
	GetItem(arg1 int, arg2 *[]float32)
	SwigIsAnnoyIndex()
	SwigGetAnnoyIndex() AnnoyIndex
}


