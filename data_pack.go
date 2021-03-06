package tower

import (
	"bytes"
	"encoding/binary"
)

type DataPack struct{}

// NewDataPack data pack instance initialization
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen get pack head's length
func (dp *DataPack) GetHeadLen() uint32 {
	//Id uint32(4字节) +  DataLen uint32(4字节)
	return 8
}

// Pack pack data, compress data
func (dp *DataPack) Pack(msg *Message) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// write data's length
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// write message id
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// write data's content
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack unpack data, uncompress data
func (dp *DataPack) Unpack(binaryData []byte) (*Message, error) {
	// create ioReader from binary input
	dataBuff := bytes.NewReader(binaryData)

	// just uncompress head data,get data's length and id
	var msg = new(Message)

	// read data len
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// read messag id
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
