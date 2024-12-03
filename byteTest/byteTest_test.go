package byteTest

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
)

func GetRandomStr() {

}

func BenchmarkBytes(b *testing.B) {

	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString("Sakura")
		_ = []byte(builder.String())
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("Sakura")
		_ = buffer.Bytes()
	}
}

func TestCompare(t *testing.T) {
	b := []byte("Sakura")
	b2 := []byte("Sakura")
	compare := bytes.Compare(b, b2)
	fmt.Println(compare)
}

func TestPrefix(t *testing.T) {
	b1 := []byte("Sakura-Key")
	//b2 := []byte("Saks")

	prefix := bytes.HasPrefix(b1, []byte("Sakuras"))
	fmt.Println(prefix)
}

func logRecordKeyWithSeq1(key []byte, seqNo uint64) []byte {
	seq := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(seq, seqNo)

	resultKey := make([]byte, n+len(key))
	// 先拷贝序列号, 再拷贝 key
	// key: 序列化+key
	copy(resultKey[:n], seq[:n])
	copy(resultKey[n:], key)

	return resultKey
}
func logRecordKeyWithSeq2(key []byte, seqNo uint64) []byte {
	resultKey := make([]byte, binary.MaxVarintLen64+len(key))
	n := binary.PutUvarint(resultKey, seqNo)
	copy(resultKey[n:], key)
	return resultKey[:n+len(key)]
}

func TestLogRecordKeyWithSeq(t *testing.T) {
	seq1 := logRecordKeyWithSeq1([]byte("Sakura"), 1)
	fmt.Println(string(seq1))
	seq2 := logRecordKeyWithSeq2([]byte("Sakura"), 1)
	fmt.Println(string(seq2))
}

func BenchmarkBytesConcat1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logRecordKeyWithSeq1([]byte("Sakura"), 1)
	}
}
func BenchmarkBytesConcat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logRecordKeyWithSeq2([]byte("Sakura"), 1)
	}
}

func TestByte1(t *testing.T) {
	//tesS := make([]int, 4, 50)
	//tesS[0] = 1
	//
	//for _, value := range tesS {
	//	fmt.Println(value)
	//}

	//test := []byte(" ")
	//
	//if test != nil {
	//	fmt.Println(len(test))
	//}
	//src := "../tmp"
	//filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
	//	if path == src {
	//		return nil
	//	}
	//	fmt.Println("Path:", path)
	//	//fileName := strings.Replace(path, src, "", 1)
	//	fmt.Println("base:", filepath.Base(path))
	//	//fmt.Println("fileName:", fileName)
	//
	//	return nil
	//})
	//ch := make(chan int, 4)
	//ch <- 1
	//fmt.Println(len(ch))

	//testSlice := []byte('A', 'B', 'C', 'D')
	//t.Log(testSlice)
}
