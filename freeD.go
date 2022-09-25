package freeD

import (
	"encoding/binary"
	"errors"
)

type FreeD struct {
	Pitch float32
	Yaw   float32
	Roll  float32
	PosZ  float32
	PosX  float32
	PosY  float32
	Zoom  int
	Focus int
}

func Decode(data []byte) (FreeD, error) {
	if checksum(data) == int(data[28]) {
		var TrackingData FreeD
		TrackingData.Pitch = getRotation(data[2:5])
		TrackingData.Yaw = getRotation(data[5:8])
		TrackingData.Roll = getRotation(data[8:11])
		TrackingData.PosZ = getPosition(data[11:14])
		TrackingData.PosX = getPosition(data[14:17])
		TrackingData.PosY = getPosition(data[17:20])
		TrackingData.Zoom = getEncoder(data[20:23])
		TrackingData.Focus = getEncoder(data[23:26])

		return TrackingData, nil
	}
	return FreeD{}, errors.New("calculated checksum does not match provided data. probalby not freeD")
}

func Encode(data FreeD) []byte {
	output := []byte{0xD1}                                     //Identifier
	output = append(output, []byte{0xFF}...)                   //ID
	output = append(output, setRotation(data.Pitch)...)        //Pitch
	output = append(output, setRotation(data.Yaw)...)          //Yaw
	output = append(output, setRotation(data.Roll)...)         //Roll
	output = append(output, setPosition(data.PosZ)...)         //X
	output = append(output, setPosition(data.PosX)...)         //Y
	output = append(output, setPosition(data.PosY)...)         //Z
	output = append(output, setEncoder(data.Zoom)...)          //Zoom
	output = append(output, setEncoder(data.Focus)...)         //Focus
	output = append(output, []byte{0x00, 0x00}...)             //Reserved
	output = append(output, []byte{byte(checksum(output))}...) //Checksum

	return output
}

func setPosition(pos float32) []byte {
	position := int64(pos * 64 * 256)
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(position))
	return data[4:7]
}

func setRotation(rot float32) []byte {
	rotation := int64(rot * 32768 * 256)
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(rotation))
	return data[4:7]
}

func setEncoder(enc int) []byte {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(enc))
	return append([]byte{0x00}, data[6:]...)
}

func getPosition(data []byte) float32 {
	return float32(int32(data[0])<<24|int32(data[1])<<16|int32(data[2])<<8) / 64 / 256
}

func getRotation(data []byte) float32 {
	return float32(int32(data[0])<<24|int32(data[1])<<16|int32(data[2])<<8) / 32768 / 256
}

func getEncoder(data []byte) int {
	value := []byte{0x00}
	value = append(value, data...)
	return int(binary.BigEndian.Uint32(value))
}

func checksum(data []byte) int {
	sum := int(64)
	for _, element := range data[:28] {
		sum = sum - int(element)
	}
	return modulo(sum, 256)
}

func modulo(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
