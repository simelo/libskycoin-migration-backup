package main

/*
#include <string.h>
#include <stdlib.h>

#include "skytypes.h"

*/
import "C"

import (
	"reflect"
	"unsafe"

	"github.com/skycoin/skycoin/src/cipher"
)

/**
 * Functions in github.com/skycoin/skycoin/src/cipher/address.go
 */

//export SKY_cipher_DecodeBase58Address
func SKY_cipher_DecodeBase58Address(_addr string, _arg1 *C.cipher__Address) uint32 {
	addr, err := cipher.DecodeBase58Address(_addr)
	errcode := libErrorCode(err)
	if err == nil {
		*_arg1 = *(*C.cipher__Address)(unsafe.Pointer(&addr))
	}
	return errcode
}

//export SKY_cipher_AddressFromBytes
func SKY_cipher_AddressFromBytes(_b []byte, _arg1 *C.cipher__Address) uint32 {
	addr, err := cipher.AddressFromBytes(_b)
	errcode := libErrorCode(err)
	if err == nil {
		*_arg1 = *(*C.cipher__Address)(unsafe.Pointer(&addr))
	}
	return errcode
}

//export SKY_cipher_AddressFromPubKey
func SKY_cipher_AddressFromPubKey(_pubKey *C.cipher__PubKey, _arg1 *C.cipher__Address) {
	pubKey := (*cipher.PubKey)(unsafe.Pointer(_pubKey))

	addr := cipher.AddressFromPubKey(*pubKey)
	*_arg1 = *(*C.cipher__Address)(unsafe.Pointer(&addr))
}

//export SKY_cipher_AddressFromSecKey
func SKY_cipher_AddressFromSecKey(_secKey *C.cipher__SecKey, _arg1 *C.cipher__Address) uint32 {
	var secKey cipher.SecKey
	secKey = *(*cipher.SecKey)(unsafe.Pointer(_secKey))
	addr, err := cipher.AddressFromSecKey(secKey)
	errcode := libErrorCode(err)
	if err == nil {
		*_arg1 = *(*C.cipher__Address)(unsafe.Pointer(&addr))
	}
	return errcode
}

//export SKY_cipher_Address_Bytes
func SKY_cipher_Address_Bytes(_addr *C.cipher__Address, _arg0 *C.GoSlice_) {
	addr := (*cipher.Address)(unsafe.Pointer(_addr))
	bytes := addr.Bytes()
	copyToGoSlice(reflect.ValueOf(bytes), _arg0)
}

//export SKY_cipher_Address_Verify
func SKY_cipher_Address_Verify(_addr *C.cipher__Address, _key *C.cipher__PubKey) uint32 {
	addr := (*cipher.Address)(unsafe.Pointer(_addr))
	key := (*cipher.PubKey)(unsafe.Pointer(_key))
	err := addr.Verify(*key)
	return libErrorCode(err)
}

//export SKY_cipher_Address_String
func SKY_cipher_Address_String(_addr *C.cipher__Address, _arg1 *C.GoString_) {
	addr := (*cipher.Address)(unsafe.Pointer(_addr))
	s := addr.String()
	copyString(s, _arg1)
}

//export SKY_cipher_Address_Checksum
func SKY_cipher_Address_Checksum(_addr *C.cipher__Address, _arg0 *C.cipher__Checksum) {
	addr := (*cipher.Address)(unsafe.Pointer(_addr))
	cs := addr.Checksum()
	C.memcpy(unsafe.Pointer(_arg0), unsafe.Pointer(&cs[0]), C.size_t(len(cs)))
}
