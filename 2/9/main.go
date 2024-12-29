package main

func pkcs7(originalBuffer []byte, blockSize uint32) (paddedBuffer []byte) {
	paddedBuffer = originalBuffer[:]
	for i := 0; i < int(blockSize)%len(originalBuffer); i++ {
		paddedBuffer = append(paddedBuffer, 4)
	}
	return
}
