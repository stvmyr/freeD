# freeD
## A simple freeD tracking protocol implementation written in golang 

### What is freeD?
freeD is a very simple protocol used to exchange camera tracking data. It was originally developed by Vinten and is now supported by a wide range of hard- and software including Unreal Engine, disguise, stYpe, Mo-Sys and Panasonic. 

The original protocol documentation can be found here:
https://www.manualsdir.com/manuals/641433/vinten-radamec-free-d.html

Note that the original system is designed to transmit the Data via RS232 or RS422. 
See manual section A.3 to get a detailed look of what's going on.

If you need support or have a idea to make this library better, don't hesitate to contact me... :)

## Install

```shell
go get github.com/stvmyr/freeD
```
## Functions

### freeD.Decode()
Decode takes a byte array (typically received via UDP nowadays), parses the data and returns a freeD struct and error. Only if the Internal checksum calculation failes, an error is returned. 

### freeD.Encode()
Encode takes a freeD struct as described below, and generates a byte array in the freeD format. This array can then transmitted via UDP.


## FreeD struct

```go
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
```

## FreeD Protocol

A typical freeD package contains 29 Bytes: 


| Offset    | Function          | Description                     |
| ----------- | ----------------- |-------------------------------- |
| 0           | Identifier        | Message Type. The Encode function always uses 0xD1. (see freeD manual section A.3.1 for further information) |
| 1           | ID                | Camera ID. This is a relict when using multiple Systems via RS232 or RS422. |
| 2:5         | Pitch             | Camera Pitch described in degree.|
| 5:8         | Yaw               | Camera Yaw described in degree.|
| 8:11        | Roll              | Camera Roll described in degree.|
| 11:14       | Position Z        | Camera Z Offset from origin. Typically described in millimeter. |
| 14:17       | Position Y        | Camera Y Offset from origin. Typically described in millimeter. |
| 17:20       | Position X        | Camera X Offset from origin. Typically described in millimeter. |
| 20:23       | Zoom              | Lens Zoom Position. Typically measured with an external encoder attached to the Lens. In the most cases this is a value between 0-4095. |
| 23:26       | Focus             | Lens Focus Position. Typically measured with an external encoder attached to the Lens. In the most cases this is a value between 0-4095. |
| 26:28       | Reserved          | Currently not used in freeD. |
| 28          | Checksum          | Checksum of the first 28 bytes. The Decode function uses the checksum to verify if the incoming data is a valid freeD package.|
