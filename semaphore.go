package semaphore

func Run(bufferedLength int64, f func() error, quit chan error) {
	bufferedChan := make(chan struct{}, bufferedLength)

	for {
		bufferedChan <- struct{}{}
		go func() {
			if err := f(); err != nil {
				quit <- err
			}
			<-bufferedChan
		}()
	}
}
