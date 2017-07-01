package m

import "reflect"

func panics(errs []error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}

type mbool struct {
	b    bool
	errs []error
}
type mstring struct {
	s    string
	errs []error
}
type mint struct {
	i    int
	errs []error
}
type mint8 struct {
	i    int8
	errs []error
}
type mint16 struct {
	i    int16
	errs []error
}
type mint32 struct {
	i    int32
	errs []error
}
type mint64 struct {
	i    int64
	errs []error
}
type muint struct {
	u    uint
	errs []error
}
type muint8 struct {
	u    uint8
	errs []error
}
type muint16 struct {
	u    uint16
	errs []error
}
type muint32 struct {
	u    uint32
	errs []error
}
type muint64 struct {
	u    uint64
	errs []error
}
type muintptr struct {
	u    uintptr
	errs []error
}
type mbyte struct {
	b    byte
	errs []error
}
type mrune struct {
	r    rune
	errs []error
}
type mfloat32 struct {
	f    float32
	errs []error
}
type mfloat64 struct {
	f    float64
	errs []error
}
type mcomplex64 struct {
	c    complex64
	errs []error
}
type mcomplex128 struct {
	c    complex128
	errs []error
}
type mtype struct {
	v    reflect.Value
	errs []error
}

func Bool(b bool, err ...error) mbool                   { return mbool{b, err} }
func String(s string, err ...error) mstring             { return mstring{s, err} }
func Int(i int, err ...error) mint                      { return mint{i, err} }
func Int8(i int8, err ...error) mint8                   { return mint8{i, err} }
func Int16(i int16, err ...error) mint16                { return mint16{i, err} }
func Int32(i int32, err ...error) mint32                { return mint32{i, err} }
func Int64(i int64, err ...error) mint64                { return mint64{i, err} }
func UInt(u uint, err ...error) muint                   { return muint{u, err} }
func UInt8(u uint8, err ...error) muint8                { return muint8{u, err} }
func UInt16(u uint16, err ...error) muint16             { return muint16{u, err} }
func UInt32(u uint32, err ...error) muint32             { return muint32{u, err} }
func UInt64(u uint64, err ...error) muint64             { return muint64{u, err} }
func UIntPtr(u uintptr, err ...error) muintptr          { return muintptr{u, err} }
func Byte(b byte, err ...error) mbyte                   { return mbyte{b, err} }
func Rune(r rune, err ...error) mrune                   { return mrune{r, err} }
func Float32(f float32, err ...error) mfloat32          { return mfloat32{f, err} }
func Float64(f float64, err ...error) mfloat64          { return mfloat64{f, err} }
func Complex64(c complex64, err ...error) mcomplex64    { return mcomplex64{c, err} }
func Complex128(c complex128, err ...error) mcomplex128 { return mcomplex128{c, err} }
func Type(v interface{}, err ...error) mtype            { return mtype{reflect.ValueOf(v), err} }

func (m mbool) Get() bool {
	panics(m.errs)
	return m.b
}
func (m mstring) Get() string {
	panics(m.errs)
	return m.s
}
func (m mint) Get() int {
	panics(m.errs)
	return m.i
}
func (m mint8) Get() int8 {
	panics(m.errs)
	return m.i
}
func (m mint16) Get() int16 {
	panics(m.errs)
	return m.i
}
func (m mint32) Get() int32 {
	panics(m.errs)
	return m.i
}
func (m mint64) Get() int64 {
	panics(m.errs)
	return m.i
}
func (m muint) Get() uint {
	panics(m.errs)
	return m.u
}
func (m muint8) Get() uint8 {
	panics(m.errs)
	return m.u
}
func (m muint16) Get() uint16 {
	panics(m.errs)
	return m.u
}
func (m muint32) Get() uint32 {
	panics(m.errs)
	return m.u
}
func (m muint64) Get() uint64 {
	panics(m.errs)
	return m.u
}
func (m muintptr) Get() uintptr {
	panics(m.errs)
	return m.u
}
func (m mbyte) Get() byte {
	panics(m.errs)
	return m.b
}
func (m mrune) Get() rune {
	panics(m.errs)
	return m.r
}
func (m mfloat32) Get() float32 {
	panics(m.errs)
	return m.f
}
func (m mfloat64) Get() float64 {
	panics(m.errs)
	return m.f
}
func (m mcomplex64) Get() complex64 {
	panics(m.errs)
	return m.c
}
func (m mcomplex128) Get() complex128 {
	panics(m.errs)
	return m.c
}
func (m mtype) Get(v interface{}) {
	panics(m.errs)
	reflect.Indirect(reflect.ValueOf(v)).Set(m.v)
}

func (m mbool) Set(fn func(bool) bool) mbool {
	return mbool{fn(m.b), m.errs}
}
func (m mstring) Set(fn func(string) string) mstring {
	return mstring{fn(m.s), m.errs}
}
func (m mint) Set(fn func(int) int) mint {
	return mint{fn(m.i), m.errs}
}
func (m mint8) Set(fn func(int8) int8) mint8 {
	return mint8{fn(m.i), m.errs}
}
func (m mint16) Set(fn func(int16) int16) mint16 {
	return mint16{fn(m.i), m.errs}
}
func (m mint32) Set(fn func(int32) int32) mint32 {
	return mint32{fn(m.i), m.errs}
}
func (m mint64) Set(fn func(int64) int64) mint64 {
	return mint64{fn(m.i), m.errs}
}
func (m muint) Set(fn func(uint) uint) muint {
	return muint{fn(m.u), m.errs}
}
func (m muint8) Set(fn func(uint8) uint8) muint8 {
	return muint8{fn(m.u), m.errs}
}
func (m muint16) Set(fn func(uint16) uint16) muint16 {
	return muint16{fn(m.u), m.errs}
}
func (m muint32) Set(fn func(uint32) uint32) muint32 {
	return muint32{fn(m.u), m.errs}
}
func (m muint64) Set(fn func(uint64) uint64) muint64 {
	return muint64{fn(m.u), m.errs}
}
func (m muintptr) Set(fn func(uintptr) uintptr) muintptr {
	return muintptr{fn(m.u), m.errs}
}
func (m mbyte) Set(fn func(byte) byte) mbyte {
	return mbyte{fn(m.b), m.errs}
}
func (m mrune) Set(fn func(rune) rune) mrune {
	return mrune{fn(m.r), m.errs}
}
func (m mfloat32) Set(fn func(float32) float32) mfloat32 {
	return mfloat32{fn(m.f), m.errs}
}
func (m mfloat64) Set(fn func(float64) float64) mfloat64 {
	return mfloat64{fn(m.f), m.errs}
}
func (m mcomplex64) Set(fn func(complex64) complex64) mcomplex64 {
	return mcomplex64{fn(m.c), m.errs}
}
func (m mcomplex128) Set(fn func(complex128) complex128) mcomplex128 {
	return mcomplex128{fn(m.c), m.errs}
}
func (m mtype) Set(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0],
		m.errs,
	}
}

func (m mbool) ToString(fn func(bool) string) mstring {
	return mstring{fn(m.b), m.errs}
}
func (m mbool) ToInt(fn func(bool) int) mint {
	return mint{fn(m.b), m.errs}
}
func (m mbool) ToInt8(fn func(bool) int8) mint8 {
	return mint8{fn(m.b), m.errs}
}
func (m mbool) ToInt16(fn func(bool) int16) mint16 {
	return mint16{fn(m.b), m.errs}
}
func (m mbool) ToInt32(fn func(bool) int32) mint32 {
	return mint32{fn(m.b), m.errs}
}
func (m mbool) ToInt64(fn func(bool) int64) mint64 {
	return mint64{fn(m.b), m.errs}
}
func (m mbool) ToUInt(fn func(bool) uint) muint {
	return muint{fn(m.b), m.errs}
}
func (m mbool) ToUInt8(fn func(bool) uint8) muint8 {
	return muint8{fn(m.b), m.errs}
}
func (m mbool) ToUInt16(fn func(bool) uint16) muint16 {
	return muint16{fn(m.b), m.errs}
}
func (m mbool) ToUInt32(fn func(bool) uint32) muint32 {
	return muint32{fn(m.b), m.errs}
}
func (m mbool) ToUInt64(fn func(bool) uint64) muint64 {
	return muint64{fn(m.b), m.errs}
}
func (m mbool) ToUIntPtr(fn func(bool) uintptr) muintptr {
	return muintptr{fn(m.b), m.errs}
}
func (m mbool) ToByte(fn func(bool) byte) mbyte {
	return mbyte{fn(m.b), m.errs}
}
func (m mbool) ToRune(fn func(bool) rune) mrune {
	return mrune{fn(m.b), m.errs}
}
func (m mbool) ToFloat32(fn func(bool) float32) mfloat32 {
	return mfloat32{fn(m.b), m.errs}
}
func (m mbool) ToFloat64(fn func(bool) float64) mfloat64 {
	return mfloat64{fn(m.b), m.errs}
}
func (m mbool) ToComplex64(fn func(bool) complex64) mcomplex64 {
	return mcomplex64{fn(m.b), m.errs}
}
func (m mbool) ToComplex128(fn func(bool) complex128) mcomplex128 {
	return mcomplex128{fn(m.b), m.errs}
}
func (m mbool) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.b)})[0],
		m.errs,
	}
}

func (m mstring) ToBool(fn func(string) bool) mbool {
	return mbool{fn(m.s), m.errs}
}
func (m mstring) ToInt(fn func(string) int) mint {
	return mint{fn(m.s), m.errs}
}
func (m mstring) ToInt8(fn func(string) int8) mint8 {
	return mint8{fn(m.s), m.errs}
}
func (m mstring) ToInt16(fn func(string) int16) mint16 {
	return mint16{fn(m.s), m.errs}
}
func (m mstring) ToInt32(fn func(string) int32) mint32 {
	return mint32{fn(m.s), m.errs}
}
func (m mstring) ToInt64(fn func(string) int64) mint64 {
	return mint64{fn(m.s), m.errs}
}
func (m mstring) ToUInt(fn func(string) uint) muint {
	return muint{fn(m.s), m.errs}
}
func (m mstring) ToUInt8(fn func(string) uint8) muint8 {
	return muint8{fn(m.s), m.errs}
}
func (m mstring) ToUInt16(fn func(string) uint16) muint16 {
	return muint16{fn(m.s), m.errs}
}
func (m mstring) ToUInt32(fn func(string) uint32) muint32 {
	return muint32{fn(m.s), m.errs}
}
func (m mstring) ToUInt64(fn func(string) uint64) muint64 {
	return muint64{fn(m.s), m.errs}
}
func (m mstring) ToUIntPtr(fn func(string) uintptr) muintptr {
	return muintptr{fn(m.s), m.errs}
}
func (m mstring) ToByte(fn func(string) byte) mbyte {
	return mbyte{fn(m.s), m.errs}
}
func (m mstring) ToRune(fn func(string) rune) mrune {
	return mrune{fn(m.s), m.errs}
}
func (m mstring) ToFloat32(fn func(string) float32) mfloat32 {
	return mfloat32{fn(m.s), m.errs}
}
func (m mstring) ToFloat64(fn func(string) float64) mfloat64 {
	return mfloat64{fn(m.s), m.errs}
}
func (m mstring) ToComplex64(fn func(string) complex64) mcomplex64 {
	return mcomplex64{fn(m.s), m.errs}
}
func (m mstring) ToComplex128(fn func(string) complex128) mcomplex128 {
	return mcomplex128{fn(m.s), m.errs}
}
func (m mstring) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.s)})[0],
		m.errs,
	}
}

func (m mint) ToBool(fn func(int) bool) mbool {
	return mbool{fn(m.i), m.errs}
}
func (m mint) ToString(fn func(int) string) mstring {
	return mstring{fn(m.i), m.errs}
}
func (m mint) ToInt8(fn func(int) int8) mint8 {
	return mint8{fn(m.i), m.errs}
}
func (m mint) ToInt16(fn func(int) int16) mint16 {
	return mint16{fn(m.i), m.errs}
}
func (m mint) ToInt32(fn func(int) int32) mint32 {
	return mint32{fn(m.i), m.errs}
}
func (m mint) ToInt64(fn func(int) int64) mint64 {
	return mint64{fn(m.i), m.errs}
}
func (m mint) ToUInt(fn func(int) uint) muint {
	return muint{fn(m.i), m.errs}
}
func (m mint) ToUInt8(fn func(int) uint8) muint8 {
	return muint8{fn(m.i), m.errs}
}
func (m mint) ToUInt16(fn func(int) uint16) muint16 {
	return muint16{fn(m.i), m.errs}
}
func (m mint) ToUInt32(fn func(int) uint32) muint32 {
	return muint32{fn(m.i), m.errs}
}
func (m mint) ToUInt64(fn func(int) uint64) muint64 {
	return muint64{fn(m.i), m.errs}
}
func (m mint) ToUIntPtr(fn func(int) uintptr) muintptr {
	return muintptr{fn(m.i), m.errs}
}
func (m mint) ToByte(fn func(int) byte) mbyte {
	return mbyte{fn(m.i), m.errs}
}
func (m mint) ToRune(fn func(int) rune) mrune {
	return mrune{fn(m.i), m.errs}
}
func (m mint) ToFloat32(fn func(int) float32) mfloat32 {
	return mfloat32{fn(m.i), m.errs}
}
func (m mint) ToFloat64(fn func(int) float64) mfloat64 {
	return mfloat64{fn(m.i), m.errs}
}
func (m mint) ToComplex64(fn func(int) complex64) mcomplex64 {
	return mcomplex64{fn(m.i), m.errs}
}
func (m mint) ToComplex128(fn func(int) complex128) mcomplex128 {
	return mcomplex128{fn(m.i), m.errs}
}
func (m mint) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.i)})[0],
		m.errs,
	}
}

func (m mint8) ToBool(fn func(int8) bool) mbool {
	return mbool{fn(m.i), m.errs}
}
func (m mint8) ToString(fn func(int8) string) mstring {
	return mstring{fn(m.i), m.errs}
}
func (m mint8) ToInt(fn func(int8) int) mint {
	return mint{fn(m.i), m.errs}
}
func (m mint8) ToInt16(fn func(int8) int16) mint16 {
	return mint16{fn(m.i), m.errs}
}
func (m mint8) ToInt32(fn func(int8) int32) mint32 {
	return mint32{fn(m.i), m.errs}
}
func (m mint8) ToInt64(fn func(int8) int64) mint64 {
	return mint64{fn(m.i), m.errs}
}
func (m mint8) ToUInt(fn func(int8) uint) muint {
	return muint{fn(m.i), m.errs}
}
func (m mint8) ToUInt8(fn func(int8) uint8) muint8 {
	return muint8{fn(m.i), m.errs}
}
func (m mint8) ToUInt16(fn func(int8) uint16) muint16 {
	return muint16{fn(m.i), m.errs}
}
func (m mint8) ToUInt32(fn func(int8) uint32) muint32 {
	return muint32{fn(m.i), m.errs}
}
func (m mint8) ToUInt64(fn func(int8) uint64) muint64 {
	return muint64{fn(m.i), m.errs}
}
func (m mint8) ToUIntPtr(fn func(int8) uintptr) muintptr {
	return muintptr{fn(m.i), m.errs}
}
func (m mint8) ToByte(fn func(int8) byte) mbyte {
	return mbyte{fn(m.i), m.errs}
}
func (m mint8) ToRune(fn func(int8) rune) mrune {
	return mrune{fn(m.i), m.errs}
}
func (m mint8) ToFloat32(fn func(int8) float32) mfloat32 {
	return mfloat32{fn(m.i), m.errs}
}
func (m mint8) ToFloat64(fn func(int8) float64) mfloat64 {
	return mfloat64{fn(m.i), m.errs}
}
func (m mint8) ToComplex64(fn func(int8) complex64) mcomplex64 {
	return mcomplex64{fn(m.i), m.errs}
}
func (m mint8) ToComplex128(fn func(int8) complex128) mcomplex128 {
	return mcomplex128{fn(m.i), m.errs}
}
func (m mint8) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.i)})[0],
		m.errs,
	}
}

func (m mint16) ToBool(fn func(int16) bool) mbool {
	return mbool{fn(m.i), m.errs}
}
func (m mint16) ToString(fn func(int16) string) mstring {
	return mstring{fn(m.i), m.errs}
}
func (m mint16) ToInt(fn func(int16) int) mint {
	return mint{fn(m.i), m.errs}
}
func (m mint16) ToInt8(fn func(int16) int8) mint8 {
	return mint8{fn(m.i), m.errs}
}
func (m mint16) ToInt32(fn func(int16) int32) mint32 {
	return mint32{fn(m.i), m.errs}
}
func (m mint16) ToInt64(fn func(int16) int64) mint64 {
	return mint64{fn(m.i), m.errs}
}
func (m mint16) ToUInt(fn func(int16) uint) muint {
	return muint{fn(m.i), m.errs}
}
func (m mint16) ToUInt8(fn func(int16) uint8) muint8 {
	return muint8{fn(m.i), m.errs}
}
func (m mint16) ToUInt16(fn func(int16) uint16) muint16 {
	return muint16{fn(m.i), m.errs}
}
func (m mint16) ToUInt32(fn func(int16) uint32) muint32 {
	return muint32{fn(m.i), m.errs}
}
func (m mint16) ToUInt64(fn func(int16) uint64) muint64 {
	return muint64{fn(m.i), m.errs}
}
func (m mint16) ToUIntPtr(fn func(int16) uintptr) muintptr {
	return muintptr{fn(m.i), m.errs}
}
func (m mint16) ToByte(fn func(int16) byte) mbyte {
	return mbyte{fn(m.i), m.errs}
}
func (m mint16) ToRune(fn func(int16) rune) mrune {
	return mrune{fn(m.i), m.errs}
}
func (m mint16) ToFloat32(fn func(int16) float32) mfloat32 {
	return mfloat32{fn(m.i), m.errs}
}
func (m mint16) ToFloat64(fn func(int16) float64) mfloat64 {
	return mfloat64{fn(m.i), m.errs}
}
func (m mint16) ToComplex64(fn func(int16) complex64) mcomplex64 {
	return mcomplex64{fn(m.i), m.errs}
}
func (m mint16) ToComplex128(fn func(int16) complex128) mcomplex128 {
	return mcomplex128{fn(m.i), m.errs}
}
func (m mint16) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.i)})[0],
		m.errs,
	}
}

func (m mint32) ToBool(fn func(int32) bool) mbool {
	return mbool{fn(m.i), m.errs}
}
func (m mint32) ToString(fn func(int32) string) mstring {
	return mstring{fn(m.i), m.errs}
}
func (m mint32) ToInt(fn func(int32) int) mint {
	return mint{fn(m.i), m.errs}
}
func (m mint32) ToInt8(fn func(int32) int8) mint8 {
	return mint8{fn(m.i), m.errs}
}
func (m mint32) ToInt16(fn func(int32) int16) mint16 {
	return mint16{fn(m.i), m.errs}
}
func (m mint32) ToInt64(fn func(int32) int64) mint64 {
	return mint64{fn(m.i), m.errs}
}
func (m mint32) ToUInt(fn func(int32) uint) muint {
	return muint{fn(m.i), m.errs}
}
func (m mint32) ToUInt8(fn func(int32) uint8) muint8 {
	return muint8{fn(m.i), m.errs}
}
func (m mint32) ToUInt16(fn func(int32) uint16) muint16 {
	return muint16{fn(m.i), m.errs}
}
func (m mint32) ToUInt32(fn func(int32) uint32) muint32 {
	return muint32{fn(m.i), m.errs}
}
func (m mint32) ToUInt64(fn func(int32) uint64) muint64 {
	return muint64{fn(m.i), m.errs}
}
func (m mint32) ToUIntPtr(fn func(int32) uintptr) muintptr {
	return muintptr{fn(m.i), m.errs}
}
func (m mint32) ToByte(fn func(int32) byte) mbyte {
	return mbyte{fn(m.i), m.errs}
}
func (m mint32) ToRune(fn func(int32) rune) mrune {
	return mrune{fn(m.i), m.errs}
}
func (m mint32) ToFloat32(fn func(int32) float32) mfloat32 {
	return mfloat32{fn(m.i), m.errs}
}
func (m mint32) ToFloat64(fn func(int32) float64) mfloat64 {
	return mfloat64{fn(m.i), m.errs}
}
func (m mint32) ToComplex64(fn func(int32) complex64) mcomplex64 {
	return mcomplex64{fn(m.i), m.errs}
}
func (m mint32) ToComplex128(fn func(int32) complex128) mcomplex128 {
	return mcomplex128{fn(m.i), m.errs}
}
func (m mint32) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.i)})[0],
		m.errs,
	}
}

func (m mint64) ToBool(fn func(int64) bool) mbool {
	return mbool{fn(m.i), m.errs}
}
func (m mint64) ToString(fn func(int64) string) mstring {
	return mstring{fn(m.i), m.errs}
}
func (m mint64) ToInt(fn func(int64) int) mint {
	return mint{fn(m.i), m.errs}
}
func (m mint64) ToInt8(fn func(int64) int8) mint8 {
	return mint8{fn(m.i), m.errs}
}
func (m mint64) ToInt16(fn func(int64) int16) mint16 {
	return mint16{fn(m.i), m.errs}
}
func (m mint64) ToInt32(fn func(int64) int32) mint32 {
	return mint32{fn(m.i), m.errs}
}
func (m mint64) ToUInt(fn func(int64) uint) muint {
	return muint{fn(m.i), m.errs}
}
func (m mint64) ToUInt8(fn func(int64) uint8) muint8 {
	return muint8{fn(m.i), m.errs}
}
func (m mint64) ToUInt16(fn func(int64) uint16) muint16 {
	return muint16{fn(m.i), m.errs}
}
func (m mint64) ToUInt32(fn func(int64) uint32) muint32 {
	return muint32{fn(m.i), m.errs}
}
func (m mint64) ToUInt64(fn func(int64) uint64) muint64 {
	return muint64{fn(m.i), m.errs}
}
func (m mint64) ToUIntPtr(fn func(int64) uintptr) muintptr {
	return muintptr{fn(m.i), m.errs}
}
func (m mint64) ToByte(fn func(int64) byte) mbyte {
	return mbyte{fn(m.i), m.errs}
}
func (m mint64) ToRune(fn func(int64) rune) mrune {
	return mrune{fn(m.i), m.errs}
}
func (m mint64) ToFloat32(fn func(int64) float32) mfloat32 {
	return mfloat32{fn(m.i), m.errs}
}
func (m mint64) ToFloat64(fn func(int64) float64) mfloat64 {
	return mfloat64{fn(m.i), m.errs}
}
func (m mint64) ToComplex64(fn func(int64) complex64) mcomplex64 {
	return mcomplex64{fn(m.i), m.errs}
}
func (m mint64) ToComplex128(fn func(int64) complex128) mcomplex128 {
	return mcomplex128{fn(m.i), m.errs}
}
func (m mint64) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.i)})[0],
		m.errs,
	}
}

func (m muint) ToBool(fn func(uint) bool) mbool {
	return mbool{fn(m.u), m.errs}
}
func (m muint) ToString(fn func(uint) string) mstring {
	return mstring{fn(m.u), m.errs}
}
func (m muint) ToInt(fn func(uint) int) mint {
	return mint{fn(m.u), m.errs}
}
func (m muint) ToInt8(fn func(uint) int8) mint8 {
	return mint8{fn(m.u), m.errs}
}
func (m muint) ToInt16(fn func(uint) int16) mint16 {
	return mint16{fn(m.u), m.errs}
}
func (m muint) ToInt32(fn func(uint) int32) mint32 {
	return mint32{fn(m.u), m.errs}
}
func (m muint) ToInt64(fn func(uint) int64) mint64 {
	return mint64{fn(m.u), m.errs}
}
func (m muint) ToUInt8(fn func(uint) uint8) muint8 {
	return muint8{fn(m.u), m.errs}
}
func (m muint) ToUInt16(fn func(uint) uint16) muint16 {
	return muint16{fn(m.u), m.errs}
}
func (m muint) ToUInt32(fn func(uint) uint32) muint32 {
	return muint32{fn(m.u), m.errs}
}
func (m muint) ToUInt64(fn func(uint) uint64) muint64 {
	return muint64{fn(m.u), m.errs}
}
func (m muint) ToUIntPtr(fn func(uint) uintptr) muintptr {
	return muintptr{fn(m.u), m.errs}
}
func (m muint) ToByte(fn func(uint) byte) mbyte {
	return mbyte{fn(m.u), m.errs}
}
func (m muint) ToRune(fn func(uint) rune) mrune {
	return mrune{fn(m.u), m.errs}
}
func (m muint) ToFloat32(fn func(uint) float32) mfloat32 {
	return mfloat32{fn(m.u), m.errs}
}
func (m muint) ToFloat64(fn func(uint) float64) mfloat64 {
	return mfloat64{fn(m.u), m.errs}
}
func (m muint) ToComplex64(fn func(uint) complex64) mcomplex64 {
	return mcomplex64{fn(m.u), m.errs}
}
func (m muint) ToComplex128(fn func(uint) complex128) mcomplex128 {
	return mcomplex128{fn(m.u), m.errs}
}
func (m muint) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.u)})[0],
		m.errs,
	}
}

func (m muint8) ToBool(fn func(uint8) bool) mbool {
	return mbool{fn(m.u), m.errs}
}
func (m muint8) ToString(fn func(uint8) string) mstring {
	return mstring{fn(m.u), m.errs}
}
func (m muint8) ToInt(fn func(uint8) int) mint {
	return mint{fn(m.u), m.errs}
}
func (m muint8) ToInt8(fn func(uint8) int8) mint8 {
	return mint8{fn(m.u), m.errs}
}
func (m muint8) ToInt16(fn func(uint8) int16) mint16 {
	return mint16{fn(m.u), m.errs}
}
func (m muint8) ToInt32(fn func(uint8) int32) mint32 {
	return mint32{fn(m.u), m.errs}
}
func (m muint8) ToInt64(fn func(uint8) int64) mint64 {
	return mint64{fn(m.u), m.errs}
}
func (m muint8) ToUInt(fn func(uint8) uint) muint {
	return muint{fn(m.u), m.errs}
}
func (m muint8) ToUInt16(fn func(uint8) uint16) muint16 {
	return muint16{fn(m.u), m.errs}
}
func (m muint8) ToUInt32(fn func(uint8) uint32) muint32 {
	return muint32{fn(m.u), m.errs}
}
func (m muint8) ToUInt64(fn func(uint8) uint64) muint64 {
	return muint64{fn(m.u), m.errs}
}
func (m muint8) ToUIntPtr(fn func(uint8) uintptr) muintptr {
	return muintptr{fn(m.u), m.errs}
}
func (m muint8) ToByte(fn func(uint8) byte) mbyte {
	return mbyte{fn(m.u), m.errs}
}
func (m muint8) ToRune(fn func(uint8) rune) mrune {
	return mrune{fn(m.u), m.errs}
}
func (m muint8) ToFloat32(fn func(uint8) float32) mfloat32 {
	return mfloat32{fn(m.u), m.errs}
}
func (m muint8) ToFloat64(fn func(uint8) float64) mfloat64 {
	return mfloat64{fn(m.u), m.errs}
}
func (m muint8) ToComplex64(fn func(uint8) complex64) mcomplex64 {
	return mcomplex64{fn(m.u), m.errs}
}
func (m muint8) ToComplex128(fn func(uint8) complex128) mcomplex128 {
	return mcomplex128{fn(m.u), m.errs}
}
func (m muint8) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.u)})[0],
		m.errs,
	}
}

func (m muint16) ToBool(fn func(uint16) bool) mbool {
	return mbool{fn(m.u), m.errs}
}
func (m muint16) ToString(fn func(uint16) string) mstring {
	return mstring{fn(m.u), m.errs}
}
func (m muint16) ToInt(fn func(uint16) int) mint {
	return mint{fn(m.u), m.errs}
}
func (m muint16) ToInt8(fn func(uint16) int8) mint8 {
	return mint8{fn(m.u), m.errs}
}
func (m muint16) ToInt16(fn func(uint16) int16) mint16 {
	return mint16{fn(m.u), m.errs}
}
func (m muint16) ToInt32(fn func(uint16) int32) mint32 {
	return mint32{fn(m.u), m.errs}
}
func (m muint16) ToInt64(fn func(uint16) int64) mint64 {
	return mint64{fn(m.u), m.errs}
}
func (m muint16) ToUInt(fn func(uint16) uint) muint {
	return muint{fn(m.u), m.errs}
}
func (m muint16) ToUInt8(fn func(uint16) uint8) muint8 {
	return muint8{fn(m.u), m.errs}
}
func (m muint16) ToUInt32(fn func(uint16) uint32) muint32 {
	return muint32{fn(m.u), m.errs}
}
func (m muint16) ToUInt64(fn func(uint16) uint64) muint64 {
	return muint64{fn(m.u), m.errs}
}
func (m muint16) ToUIntPtr(fn func(uint16) uintptr) muintptr {
	return muintptr{fn(m.u), m.errs}
}
func (m muint16) ToByte(fn func(uint16) byte) mbyte {
	return mbyte{fn(m.u), m.errs}
}
func (m muint16) ToRune(fn func(uint16) rune) mrune {
	return mrune{fn(m.u), m.errs}
}
func (m muint16) ToFloat32(fn func(uint16) float32) mfloat32 {
	return mfloat32{fn(m.u), m.errs}
}
func (m muint16) ToFloat64(fn func(uint16) float64) mfloat64 {
	return mfloat64{fn(m.u), m.errs}
}
func (m muint16) ToComplex64(fn func(uint16) complex64) mcomplex64 {
	return mcomplex64{fn(m.u), m.errs}
}
func (m muint16) ToComplex128(fn func(uint16) complex128) mcomplex128 {
	return mcomplex128{fn(m.u), m.errs}
}
func (m muint16) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.u)})[0],
		m.errs,
	}
}

func (m muint32) ToBool(fn func(uint32) bool) mbool {
	return mbool{fn(m.u), m.errs}
}
func (m muint32) ToString(fn func(uint32) string) mstring {
	return mstring{fn(m.u), m.errs}
}
func (m muint32) ToInt(fn func(uint32) int) mint {
	return mint{fn(m.u), m.errs}
}
func (m muint32) ToInt8(fn func(uint32) int8) mint8 {
	return mint8{fn(m.u), m.errs}
}
func (m muint32) ToInt16(fn func(uint32) int16) mint16 {
	return mint16{fn(m.u), m.errs}
}
func (m muint32) ToInt32(fn func(uint32) int32) mint32 {
	return mint32{fn(m.u), m.errs}
}
func (m muint32) ToInt64(fn func(uint32) int64) mint64 {
	return mint64{fn(m.u), m.errs}
}
func (m muint32) ToUInt(fn func(uint32) uint) muint {
	return muint{fn(m.u), m.errs}
}
func (m muint32) ToUInt8(fn func(uint32) uint8) muint8 {
	return muint8{fn(m.u), m.errs}
}
func (m muint32) ToUInt16(fn func(uint32) uint16) muint16 {
	return muint16{fn(m.u), m.errs}
}
func (m muint32) ToUInt64(fn func(uint32) uint64) muint64 {
	return muint64{fn(m.u), m.errs}
}
func (m muint32) ToUIntPtr(fn func(uint32) uintptr) muintptr {
	return muintptr{fn(m.u), m.errs}
}
func (m muint32) ToByte(fn func(uint32) byte) mbyte {
	return mbyte{fn(m.u), m.errs}
}
func (m muint32) ToRune(fn func(uint32) rune) mrune {
	return mrune{fn(m.u), m.errs}
}
func (m muint32) ToFloat32(fn func(uint32) float32) mfloat32 {
	return mfloat32{fn(m.u), m.errs}
}
func (m muint32) ToFloat64(fn func(uint32) float64) mfloat64 {
	return mfloat64{fn(m.u), m.errs}
}
func (m muint32) ToComplex64(fn func(uint32) complex64) mcomplex64 {
	return mcomplex64{fn(m.u), m.errs}
}
func (m muint32) ToComplex128(fn func(uint32) complex128) mcomplex128 {
	return mcomplex128{fn(m.u), m.errs}
}
func (m muint32) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.u)})[0],
		m.errs,
	}
}

func (m muint64) ToBool(fn func(uint64) bool) mbool {
	return mbool{fn(m.u), m.errs}
}
func (m muint64) ToString(fn func(uint64) string) mstring {
	return mstring{fn(m.u), m.errs}
}
func (m muint64) ToInt(fn func(uint64) int) mint {
	return mint{fn(m.u), m.errs}
}
func (m muint64) ToInt8(fn func(uint64) int8) mint8 {
	return mint8{fn(m.u), m.errs}
}
func (m muint64) ToInt16(fn func(uint64) int16) mint16 {
	return mint16{fn(m.u), m.errs}
}
func (m muint64) ToInt32(fn func(uint64) int32) mint32 {
	return mint32{fn(m.u), m.errs}
}
func (m muint64) ToInt64(fn func(uint64) int64) mint64 {
	return mint64{fn(m.u), m.errs}
}
func (m muint64) ToUInt(fn func(uint64) uint) muint {
	return muint{fn(m.u), m.errs}
}
func (m muint64) ToUInt8(fn func(uint64) uint8) muint8 {
	return muint8{fn(m.u), m.errs}
}
func (m muint64) ToUInt16(fn func(uint64) uint16) muint16 {
	return muint16{fn(m.u), m.errs}
}
func (m muint64) ToUInt32(fn func(uint64) uint32) muint32 {
	return muint32{fn(m.u), m.errs}
}
func (m muint64) ToUIntPtr(fn func(uint64) uintptr) muintptr {
	return muintptr{fn(m.u), m.errs}
}
func (m muint64) ToByte(fn func(uint64) byte) mbyte {
	return mbyte{fn(m.u), m.errs}
}
func (m muint64) ToRune(fn func(uint64) rune) mrune {
	return mrune{fn(m.u), m.errs}
}
func (m muint64) ToFloat32(fn func(uint64) float32) mfloat32 {
	return mfloat32{fn(m.u), m.errs}
}
func (m muint64) ToFloat64(fn func(uint64) float64) mfloat64 {
	return mfloat64{fn(m.u), m.errs}
}
func (m muint64) ToComplex64(fn func(uint64) complex64) mcomplex64 {
	return mcomplex64{fn(m.u), m.errs}
}
func (m muint64) ToComplex128(fn func(uint64) complex128) mcomplex128 {
	return mcomplex128{fn(m.u), m.errs}
}
func (m muint64) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.u)})[0],
		m.errs,
	}
}

func (m muintptr) ToBool(fn func(uintptr) bool) mbool {
	return mbool{fn(m.u), m.errs}
}
func (m muintptr) ToString(fn func(uintptr) string) mstring {
	return mstring{fn(m.u), m.errs}
}
func (m muintptr) ToInt(fn func(uintptr) int) mint {
	return mint{fn(m.u), m.errs}
}
func (m muintptr) ToInt8(fn func(uintptr) int8) mint8 {
	return mint8{fn(m.u), m.errs}
}
func (m muintptr) ToInt16(fn func(uintptr) int16) mint16 {
	return mint16{fn(m.u), m.errs}
}
func (m muintptr) ToInt32(fn func(uintptr) int32) mint32 {
	return mint32{fn(m.u), m.errs}
}
func (m muintptr) ToInt64(fn func(uintptr) int64) mint64 {
	return mint64{fn(m.u), m.errs}
}
func (m muintptr) ToUInt(fn func(uintptr) uint) muint {
	return muint{fn(m.u), m.errs}
}
func (m muintptr) ToUInt8(fn func(uintptr) uint8) muint8 {
	return muint8{fn(m.u), m.errs}
}
func (m muintptr) ToUInt16(fn func(uintptr) uint16) muint16 {
	return muint16{fn(m.u), m.errs}
}
func (m muintptr) ToUInt32(fn func(uintptr) uint32) muint32 {
	return muint32{fn(m.u), m.errs}
}
func (m muintptr) ToUInt64(fn func(uintptr) uint64) muint64 {
	return muint64{fn(m.u), m.errs}
}
func (m muintptr) ToByte(fn func(uintptr) byte) mbyte {
	return mbyte{fn(m.u), m.errs}
}
func (m muintptr) ToRune(fn func(uintptr) rune) mrune {
	return mrune{fn(m.u), m.errs}
}
func (m muintptr) ToFloat32(fn func(uintptr) float32) mfloat32 {
	return mfloat32{fn(m.u), m.errs}
}
func (m muintptr) ToFloat64(fn func(uintptr) float64) mfloat64 {
	return mfloat64{fn(m.u), m.errs}
}
func (m muintptr) ToComplex64(fn func(uintptr) complex64) mcomplex64 {
	return mcomplex64{fn(m.u), m.errs}
}
func (m muintptr) ToComplex128(fn func(uintptr) complex128) mcomplex128 {
	return mcomplex128{fn(m.u), m.errs}
}
func (m muintptr) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.u)})[0],
		m.errs,
	}
}

func (m mbyte) ToBool(fn func(byte) bool) mbool {
	return mbool{fn(m.b), m.errs}
}
func (m mbyte) ToString(fn func(byte) string) mstring {
	return mstring{fn(m.b), m.errs}
}
func (m mbyte) ToInt(fn func(byte) int) mint {
	return mint{fn(m.b), m.errs}
}
func (m mbyte) ToInt8(fn func(byte) int8) mint8 {
	return mint8{fn(m.b), m.errs}
}
func (m mbyte) ToInt16(fn func(byte) int16) mint16 {
	return mint16{fn(m.b), m.errs}
}
func (m mbyte) ToInt32(fn func(byte) int32) mint32 {
	return mint32{fn(m.b), m.errs}
}
func (m mbyte) ToInt64(fn func(byte) int64) mint64 {
	return mint64{fn(m.b), m.errs}
}
func (m mbyte) ToUInt(fn func(byte) uint) muint {
	return muint{fn(m.b), m.errs}
}
func (m mbyte) ToUInt8(fn func(byte) uint8) muint8 {
	return muint8{fn(m.b), m.errs}
}
func (m mbyte) ToUInt16(fn func(byte) uint16) muint16 {
	return muint16{fn(m.b), m.errs}
}
func (m mbyte) ToUInt32(fn func(byte) uint32) muint32 {
	return muint32{fn(m.b), m.errs}
}
func (m mbyte) ToUInt64(fn func(byte) uint64) muint64 {
	return muint64{fn(m.b), m.errs}
}
func (m mbyte) ToUIntPtr(fn func(byte) uintptr) muintptr {
	return muintptr{fn(m.b), m.errs}
}
func (m mbyte) ToRune(fn func(byte) rune) mrune {
	return mrune{fn(m.b), m.errs}
}
func (m mbyte) ToFloat32(fn func(byte) float32) mfloat32 {
	return mfloat32{fn(m.b), m.errs}
}
func (m mbyte) ToFloat64(fn func(byte) float64) mfloat64 {
	return mfloat64{fn(m.b), m.errs}
}
func (m mbyte) ToComplex64(fn func(byte) complex64) mcomplex64 {
	return mcomplex64{fn(m.b), m.errs}
}
func (m mbyte) ToComplex128(fn func(byte) complex128) mcomplex128 {
	return mcomplex128{fn(m.b), m.errs}
}
func (m mbyte) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.b)})[0],
		m.errs,
	}
}

func (m mrune) ToBool(fn func(rune) bool) mbool {
	return mbool{fn(m.r), m.errs}
}
func (m mrune) ToString(fn func(rune) string) mstring {
	return mstring{fn(m.r), m.errs}
}
func (m mrune) ToInt(fn func(rune) int) mint {
	return mint{fn(m.r), m.errs}
}
func (m mrune) ToInt8(fn func(rune) int8) mint8 {
	return mint8{fn(m.r), m.errs}
}
func (m mrune) ToInt16(fn func(rune) int16) mint16 {
	return mint16{fn(m.r), m.errs}
}
func (m mrune) ToInt32(fn func(rune) int32) mint32 {
	return mint32{fn(m.r), m.errs}
}
func (m mrune) ToInt64(fn func(rune) int64) mint64 {
	return mint64{fn(m.r), m.errs}
}
func (m mrune) ToUInt(fn func(rune) uint) muint {
	return muint{fn(m.r), m.errs}
}
func (m mrune) ToUInt8(fn func(rune) uint8) muint8 {
	return muint8{fn(m.r), m.errs}
}
func (m mrune) ToUInt16(fn func(rune) uint16) muint16 {
	return muint16{fn(m.r), m.errs}
}
func (m mrune) ToUInt32(fn func(rune) uint32) muint32 {
	return muint32{fn(m.r), m.errs}
}
func (m mrune) ToUInt64(fn func(rune) uint64) muint64 {
	return muint64{fn(m.r), m.errs}
}
func (m mrune) ToUIntPtr(fn func(rune) uintptr) muintptr {
	return muintptr{fn(m.r), m.errs}
}
func (m mrune) ToByte(fn func(rune) byte) mbyte {
	return mbyte{fn(m.r), m.errs}
}
func (m mrune) ToFloat32(fn func(rune) float32) mfloat32 {
	return mfloat32{fn(m.r), m.errs}
}
func (m mrune) ToFloat64(fn func(rune) float64) mfloat64 {
	return mfloat64{fn(m.r), m.errs}
}
func (m mrune) ToComplex64(fn func(rune) complex64) mcomplex64 {
	return mcomplex64{fn(m.r), m.errs}
}
func (m mrune) ToComplex128(fn func(rune) complex128) mcomplex128 {
	return mcomplex128{fn(m.r), m.errs}
}
func (m mrune) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.r)})[0],
		m.errs,
	}
}

func (m mfloat32) ToBool(fn func(float32) bool) mbool {
	return mbool{fn(m.f), m.errs}
}
func (m mfloat32) ToString(fn func(float32) string) mstring {
	return mstring{fn(m.f), m.errs}
}
func (m mfloat32) ToInt(fn func(float32) int) mint {
	return mint{fn(m.f), m.errs}
}
func (m mfloat32) ToInt8(fn func(float32) int8) mint8 {
	return mint8{fn(m.f), m.errs}
}
func (m mfloat32) ToInt16(fn func(float32) int16) mint16 {
	return mint16{fn(m.f), m.errs}
}
func (m mfloat32) ToInt32(fn func(float32) int32) mint32 {
	return mint32{fn(m.f), m.errs}
}
func (m mfloat32) ToInt64(fn func(float32) int64) mint64 {
	return mint64{fn(m.f), m.errs}
}
func (m mfloat32) ToUInt(fn func(float32) uint) muint {
	return muint{fn(m.f), m.errs}
}
func (m mfloat32) ToUInt8(fn func(float32) uint8) muint8 {
	return muint8{fn(m.f), m.errs}
}
func (m mfloat32) ToUInt16(fn func(float32) uint16) muint16 {
	return muint16{fn(m.f), m.errs}
}
func (m mfloat32) ToUInt32(fn func(float32) uint32) muint32 {
	return muint32{fn(m.f), m.errs}
}
func (m mfloat32) ToUInt64(fn func(float32) uint64) muint64 {
	return muint64{fn(m.f), m.errs}
}
func (m mfloat32) ToUIntPtr(fn func(float32) uintptr) muintptr {
	return muintptr{fn(m.f), m.errs}
}
func (m mfloat32) ToByte(fn func(float32) byte) mbyte {
	return mbyte{fn(m.f), m.errs}
}
func (m mfloat32) ToRune(fn func(float32) rune) mrune {
	return mrune{fn(m.f), m.errs}
}
func (m mfloat32) ToFloat64(fn func(float32) float64) mfloat64 {
	return mfloat64{fn(m.f), m.errs}
}
func (m mfloat32) ToComplex64(fn func(float32) complex64) mcomplex64 {
	return mcomplex64{fn(m.f), m.errs}
}
func (m mfloat32) ToComplex128(fn func(float32) complex128) mcomplex128 {
	return mcomplex128{fn(m.f), m.errs}
}
func (m mfloat32) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.f)})[0],
		m.errs,
	}
}

func (m mfloat64) ToBool(fn func(float64) bool) mbool {
	return mbool{fn(m.f), m.errs}
}
func (m mfloat64) ToString(fn func(float64) string) mstring {
	return mstring{fn(m.f), m.errs}
}
func (m mfloat64) ToInt(fn func(float64) int) mint {
	return mint{fn(m.f), m.errs}
}
func (m mfloat64) ToInt8(fn func(float64) int8) mint8 {
	return mint8{fn(m.f), m.errs}
}
func (m mfloat64) ToInt16(fn func(float64) int16) mint16 {
	return mint16{fn(m.f), m.errs}
}
func (m mfloat64) ToInt32(fn func(float64) int32) mint32 {
	return mint32{fn(m.f), m.errs}
}
func (m mfloat64) ToInt64(fn func(float64) int64) mint64 {
	return mint64{fn(m.f), m.errs}
}
func (m mfloat64) ToUInt(fn func(float64) uint) muint {
	return muint{fn(m.f), m.errs}
}
func (m mfloat64) ToUInt8(fn func(float64) uint8) muint8 {
	return muint8{fn(m.f), m.errs}
}
func (m mfloat64) ToUInt16(fn func(float64) uint16) muint16 {
	return muint16{fn(m.f), m.errs}
}
func (m mfloat64) ToUInt32(fn func(float64) uint32) muint32 {
	return muint32{fn(m.f), m.errs}
}
func (m mfloat64) ToUInt64(fn func(float64) uint64) muint64 {
	return muint64{fn(m.f), m.errs}
}
func (m mfloat64) ToUIntPtr(fn func(float64) uintptr) muintptr {
	return muintptr{fn(m.f), m.errs}
}
func (m mfloat64) ToByte(fn func(float64) byte) mbyte {
	return mbyte{fn(m.f), m.errs}
}
func (m mfloat64) ToRune(fn func(float64) rune) mrune {
	return mrune{fn(m.f), m.errs}
}
func (m mfloat64) ToFloat32(fn func(float64) float32) mfloat32 {
	return mfloat32{fn(m.f), m.errs}
}
func (m mfloat64) ToComplex64(fn func(float64) complex64) mcomplex64 {
	return mcomplex64{fn(m.f), m.errs}
}
func (m mfloat64) ToComplex128(fn func(float64) complex128) mcomplex128 {
	return mcomplex128{fn(m.f), m.errs}
}
func (m mfloat64) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.f)})[0],
		m.errs,
	}
}

func (m mcomplex64) ToBool(fn func(complex64) bool) mbool {
	return mbool{fn(m.c), m.errs}
}
func (m mcomplex64) ToString(fn func(complex64) string) mstring {
	return mstring{fn(m.c), m.errs}
}
func (m mcomplex64) ToInt(fn func(complex64) int) mint {
	return mint{fn(m.c), m.errs}
}
func (m mcomplex64) ToInt8(fn func(complex64) int8) mint8 {
	return mint8{fn(m.c), m.errs}
}
func (m mcomplex64) ToInt16(fn func(complex64) int16) mint16 {
	return mint16{fn(m.c), m.errs}
}
func (m mcomplex64) ToInt32(fn func(complex64) int32) mint32 {
	return mint32{fn(m.c), m.errs}
}
func (m mcomplex64) ToInt64(fn func(complex64) int64) mint64 {
	return mint64{fn(m.c), m.errs}
}
func (m mcomplex64) ToUInt(fn func(complex64) uint) muint {
	return muint{fn(m.c), m.errs}
}
func (m mcomplex64) ToUInt8(fn func(complex64) uint8) muint8 {
	return muint8{fn(m.c), m.errs}
}
func (m mcomplex64) ToUInt16(fn func(complex64) uint16) muint16 {
	return muint16{fn(m.c), m.errs}
}
func (m mcomplex64) ToUInt32(fn func(complex64) uint32) muint32 {
	return muint32{fn(m.c), m.errs}
}
func (m mcomplex64) ToUInt64(fn func(complex64) uint64) muint64 {
	return muint64{fn(m.c), m.errs}
}
func (m mcomplex64) ToUIntPtr(fn func(complex64) uintptr) muintptr {
	return muintptr{fn(m.c), m.errs}
}
func (m mcomplex64) ToByte(fn func(complex64) byte) mbyte {
	return mbyte{fn(m.c), m.errs}
}
func (m mcomplex64) ToRune(fn func(complex64) rune) mrune {
	return mrune{fn(m.c), m.errs}
}
func (m mcomplex64) ToFloat32(fn func(complex64) float32) mfloat32 {
	return mfloat32{fn(m.c), m.errs}
}
func (m mcomplex64) ToFloat64(fn func(complex64) float64) mfloat64 {
	return mfloat64{fn(m.c), m.errs}
}
func (m mcomplex64) ToComplex128(fn func(complex64) complex128) mcomplex128 {
	return mcomplex128{fn(m.c), m.errs}
}
func (m mcomplex64) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.c)})[0],
		m.errs,
	}
}

func (m mcomplex128) ToBool(fn func(complex128) bool) mbool {
	return mbool{fn(m.c), m.errs}
}
func (m mcomplex128) ToString(fn func(complex128) string) mstring {
	return mstring{fn(m.c), m.errs}
}
func (m mcomplex128) ToInt(fn func(complex128) int) mint {
	return mint{fn(m.c), m.errs}
}
func (m mcomplex128) ToInt8(fn func(complex128) int8) mint8 {
	return mint8{fn(m.c), m.errs}
}
func (m mcomplex128) ToInt16(fn func(complex128) int16) mint16 {
	return mint16{fn(m.c), m.errs}
}
func (m mcomplex128) ToInt32(fn func(complex128) int32) mint32 {
	return mint32{fn(m.c), m.errs}
}
func (m mcomplex128) ToInt64(fn func(complex128) int64) mint64 {
	return mint64{fn(m.c), m.errs}
}
func (m mcomplex128) ToUInt(fn func(complex128) uint) muint {
	return muint{fn(m.c), m.errs}
}
func (m mcomplex128) ToUInt8(fn func(complex128) uint8) muint8 {
	return muint8{fn(m.c), m.errs}
}
func (m mcomplex128) ToUInt16(fn func(complex128) uint16) muint16 {
	return muint16{fn(m.c), m.errs}
}
func (m mcomplex128) ToUInt32(fn func(complex128) uint32) muint32 {
	return muint32{fn(m.c), m.errs}
}
func (m mcomplex128) ToUInt64(fn func(complex128) uint64) muint64 {
	return muint64{fn(m.c), m.errs}
}
func (m mcomplex128) ToUIntPtr(fn func(complex128) uintptr) muintptr {
	return muintptr{fn(m.c), m.errs}
}
func (m mcomplex128) ToByte(fn func(complex128) byte) mbyte {
	return mbyte{fn(m.c), m.errs}
}
func (m mcomplex128) ToRune(fn func(complex128) rune) mrune {
	return mrune{fn(m.c), m.errs}
}
func (m mcomplex128) ToFloat32(fn func(complex128) float32) mfloat32 {
	return mfloat32{fn(m.c), m.errs}
}
func (m mcomplex128) ToFloat64(fn func(complex128) float64) mfloat64 {
	return mfloat64{fn(m.c), m.errs}
}
func (m mcomplex128) ToComplex64(fn func(complex128) complex64) mcomplex64 {
	return mcomplex64{fn(m.c), m.errs}
}
func (m mcomplex128) ToType(fn interface{}) mtype {
	return mtype{
		reflect.ValueOf(fn).
			Call([]reflect.Value{reflect.ValueOf(m.c)})[0],
		m.errs,
	}
}

func (m mtype) ToBool(fn interface{}) mbool {
	return mbool{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(bool),
		m.errs,
	}
}
func (m mtype) ToString(fn interface{}) mstring {
	return mstring{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(string),
		m.errs,
	}
}
func (m mtype) ToInt(fn interface{}) mint {
	return mint{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(int),
		m.errs,
	}
}
func (m mtype) ToInt8(fn interface{}) mint8 {
	return mint8{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(int8),
		m.errs,
	}
}
func (m mtype) ToInt16(fn interface{}) mint16 {
	return mint16{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(int16),
		m.errs,
	}
}
func (m mtype) ToInt32(fn interface{}) mint32 {
	return mint32{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(int32),
		m.errs,
	}
}
func (m mtype) ToInt64(fn interface{}) mint64 {
	return mint64{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(int64),
		m.errs,
	}
}
func (m mtype) ToUint(fn interface{}) muint {
	return muint{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(uint),
		m.errs,
	}
}
func (m mtype) ToUint8(fn interface{}) muint8 {
	return muint8{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(uint8),
		m.errs,
	}
}
func (m mtype) ToUint16(fn interface{}) muint16 {
	return muint16{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(uint16),
		m.errs,
	}
}
func (m mtype) ToUint32(fn interface{}) muint32 {
	return muint32{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(uint32),
		m.errs,
	}
}
func (m mtype) ToUint64(fn interface{}) muint64 {
	return muint64{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(uint64),
		m.errs,
	}
}
func (m mtype) ToUintptr(fn interface{}) muintptr {
	return muintptr{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(uintptr),
		m.errs,
	}
}
func (m mtype) ToByte(fn interface{}) mbyte {
	return mbyte{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(byte),
		m.errs,
	}
}
func (m mtype) ToRune(fn interface{}) mrune {
	return mrune{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(rune),
		m.errs,
	}
}
func (m mtype) ToFloat32(fn interface{}) mfloat32 {
	return mfloat32{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(float32),
		m.errs,
	}
}
func (m mtype) ToFloat64(fn interface{}) mfloat64 {
	return mfloat64{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(float64),
		m.errs,
	}
}
func (m mtype) ToComplex64(fn interface{}) mcomplex64 {
	return mcomplex64{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(complex64),
		m.errs,
	}
}
func (m mtype) ToComplex128(fn interface{}) mcomplex128 {
	return mcomplex128{
		reflect.ValueOf(fn).
			Call([]reflect.Value{m.v})[0].
			Interface().(complex128),
		m.errs,
	}
}
