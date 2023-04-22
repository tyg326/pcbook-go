package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/techschool/pcbook/pb"
	"gitlab.com/techschool/pcbook/sample"
	"gitlab.com/techschool/pcbook/serializer"
	"google.golang.org/protobuf/proto"
)

// 单元测试 方法以Test开头,并将指向testing.T对象的指针作为输入参数
func TestFileSerializer(t *testing.T) {
	// 表示与其它测试 并行运行
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

	jsonFile := "../tmp/laptop.json"

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
