package sm3

//
//import (
//	"testing"
//	"github.com/stretchr/testify/assert"
//	"github.com/tendermint/tendermint/libs/log"
//	"encoding/hex"
//)
//
//func mockLogger() log.Logger {
//	return log.NewNopLogger()
//}
//
//// test create sm3 instance
//func TestNew(t *testing.T) {
//	assert := assert.New(t)
//	sm3 := New()
//	assert.NotNil(sm3)
//}
//
//// test write data to sm3
//func TestSm3_Write(t *testing.T) {
//	assert := assert.New(t)
//	sm3 := New()
//	assert.NotNil(sm3)
//
//	data := []byte{'a', 'b', 'c'}
//	n, err := sm3.Write(data)
//	assert.Nil(err)
//	assert.Equal(len(data), n)
//}
//
//// test get data hash
//func TestSm3_Sum(t *testing.T) {
//	assert := assert.New(t)
//	sm3 := New()
//	assert.NotNil(sm3)
//
//	data := []byte{'a', 'b', 'c'}
//	n, err := sm3.Write(data)
//	assert.Nil(err)
//	assert.Equal(len(data), n)
//
//	sig := sm3.Sum(nil)
//	expected, _ := hex.DecodeString("66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0")
//	assert.Equal(expected, sig)
//}
//
////test get sm3's size
//func TestSm3_Size(t *testing.T) {
//	assert := assert.New(t)
//	sm3 := New()
//	assert.NotNil(sm3)
//
//	assert.Equal(Size, sm3.Size())
//}
//
//// test get sm3's block size
//func TestSm3_BlockSize(t *testing.T) {
//	assert := assert.New(t)
//	sm3 := New()
//	assert.NotNil(sm3)
//
//	assert.Equal(BlockSize, sm3.BlockSize())
//}
//
//// test reset sm3
//func TestSm3_Reset(t *testing.T) {
//	assert := assert.New(t)
//	sm3 := New()
//	assert.NotNil(sm3)
//
//	data := []byte{'h', 'e', 'l', 'l', 'o'}
//	n, err := sm3.Write(data)
//	assert.Nil(err)
//	assert.Equal(len(data), n)
//
//	sm3.Reset()
//	data = []byte{'a', 'b', 'c'}
//	n, err = sm3.Write(data)
//	assert.Nil(err)
//	assert.Equal(len(data), n)
//	sig := sm3.Sum(nil)
//	expected, _ := hex.DecodeString("66c7f0f462eeedd9d1f2d46bdc10e4e24167c4875cf2f7a2297da02b8f4ba8e0")
//	assert.Equal(expected, sig)
//}
