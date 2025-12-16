package sm3

//
///*
//#cgo CFLAGS: -I../include/
//#cgo LDFLAGS: -L${SRCDIR}/../lib/ -lcrypto
//#include <openssl/evp.h>
//*/
//import "C"
//import (
//	"hash"
//	"unsafe"
//	"runtime"
//	"errors"
//	"github.com/DSiSc/craft/log"
//)
//
//// The size of a SM3 checksum in bytes.
//const Size = 32
//
//// The blocksize of SM3 in bytes.
//const BlockSize = 64
//// openssl sm3 wrapper.
//type sm3 struct {
//	ctx *C.EVP_MD_CTX
//	md *C.EVP_MD
//}
//
//// New create sm3 hash instance.
//func New() hash.Hash {
//	cname := C.CString("SM3")
//	defer C.free(unsafe.Pointer(cname))
//	md := C.EVP_get_digestbyname(cname)
//	if md == nil {
//		log.Error("failed to get digest with name 'SM3'")
//		return nil
//	}
//
//	ctx := C.EVP_MD_CTX_new()
//	if ctx == nil {
//		return nil
//	}
//
//	self := &sm3{
//		ctx:ctx,
//		md:md,
//	}
//	runtime.SetFinalizer(self, func(sm3Instance *sm3) {
//		C.EVP_MD_CTX_free(sm3Instance.ctx)
//	})
//
//	//if 1 != C.EVP_DigestInit_ex(ctx, md, nil) {
//	if 1 != C.EVP_DigestInit(ctx, md) {
//		log.Error("failed to init digest context")
//		return nil
//	}
//	return self
//}
//
//func (self *sm3) Write(data []byte) (n int, err error)  {
//	if len(data) == 0 {
//		return 0, nil
//	}
//	if 1 != C.EVP_DigestUpdate(self.ctx, unsafe.Pointer(&data[0]), C.size_t(len(data))) {
//		return 0, errors.New("hello")
//	}
//	return len(data), nil
//}
//
//func (self *sm3) Sum(b []byte) []byte  {
//	outbuf := make([]byte, 64)
//	outlen := C.uint(len(outbuf))
//	if 1 != C.EVP_DigestFinal(self.ctx, (*C.uchar)(unsafe.Pointer(&outbuf[0])), &outlen) {
//		return nil
//	}
//	if b != nil && len(b) > 0 {
//		return append(b, outbuf[:outlen]...)
//	}
//	return outbuf[:outlen]
//}
//
//func (self *sm3) Reset()  {
//	if 1 != C.EVP_MD_CTX_reset(self.ctx) || 1 != C.EVP_DigestInit(self.ctx, self.md) {
//		log.Error("failed to reset digest context")
//	}
//}
//
//func (self *sm3) Size() int {
//	return Size
//}
//
//func (self *sm3) BlockSize() int {
//	return BlockSize
//}
