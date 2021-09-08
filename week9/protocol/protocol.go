package protocol

import (
	"bytes"
	"encoding/binary"
)

const (
	DEFAULE_HEADER           = "kiri"
	DEFAULT_HEADER_LENGTH    = 4
	DEFAULT_SAVE_DATA_LENGTH = 4
)

type Packet struct {
	Header         string
	HeaderLength   int32
	SaveDataLength int32
	Data           []byte
}

func (this *Packet) SetHeader(header string) *Packet {
	this.Header = header
	this.HeaderLength = int32(len([]byte(header)))
	return this
}

func NewDefaultPacket(data []byte) *Packet {
	return &Packet{
		Header:         DEFAULE_HEADER,
		HeaderLength:   DEFAULT_HEADER_LENGTH,
		SaveDataLength: DEFAULT_SAVE_DATA_LENGTH,
		Data:           data,
	}
}

func (this *Packet) Packet() []byte {
	return append(append([]byte(this.Header), this.IntToBytes(int32(len(this.Data)))...), this.Data...)
}

func (this *Packet) UnPacket() ([]byte, []byte) {
	dataLen := int32(len(this.Data))
	var whole []byte
	var i int32
	for i = 0; i < dataLen; i++ {
		//够了
		if dataLen < i+this.HeaderLength+this.SaveDataLength {
			break
		}
		//找到header
		if string(this.Data[i:i+this.HeaderLength]) == this.Header {
			saveDataLenBeginIndex := i + this.HeaderLength
			actualDataLen := this.BytesToInt(this.Data[saveDataLenBeginIndex : saveDataLenBeginIndex+this.SaveDataLength])
			//小于一个包的情况
			if dataLen < i+this.HeaderLength+this.SaveDataLength+actualDataLen {
				break
			}
			//得到包
			packageData := this.Data[saveDataLenBeginIndex+this.SaveDataLength : saveDataLenBeginIndex+this.SaveDataLength+actualDataLen]
			//发送
			whole = packageData
			//处理下一个包
			i += this.HeaderLength + this.SaveDataLength + actualDataLen - 1
			break
		}
	}
	if i >= dataLen {
		return []byte{}, whole
	}

	return this.Data[i:], whole
}

func (this *Packet) IntToBytes(i int32) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, i)
	return byteBuffer.Bytes()
}

func (this *Packet) BytesToInt(data []byte) int32 {
	var val int32
	byteBuffer := bytes.NewBuffer(data)
	binary.Read(byteBuffer, binary.BigEndian, &val)
	return val
}
