package LoggerStreamStorage

func (s *Storage) Send(message string) (err error) {
	err = s.sendToStream(message)
	return
}

func (s *Storage) sendToStream(message string) (err error) {
	_, err = s.stream.Write([]byte(message))
	return
}
