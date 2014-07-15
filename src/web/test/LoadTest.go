package test

import "time"

// address - 127.0.0.1:80
// loads - 100 clients
// expectedTime - 300 * time.Second
// return whether pass the test
func LoadTest(address string, loads int, expectedTime time.Duration) (bool, error) {

	return true, nil
}
