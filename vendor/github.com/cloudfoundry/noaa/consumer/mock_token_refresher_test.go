// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package consumer_test

type mockTokenRefresher struct {
	RefreshAuthTokenCalled chan bool
	RefreshAuthTokenOutput struct {
		Token     chan string
		AuthError chan error
	}
}

func newMockTokenRefresher() *mockTokenRefresher {
	m := &mockTokenRefresher{}
	m.RefreshAuthTokenCalled = make(chan bool, 100)
	m.RefreshAuthTokenOutput.Token = make(chan string, 100)
	m.RefreshAuthTokenOutput.AuthError = make(chan error, 100)
	return m
}
func (m *mockTokenRefresher) RefreshAuthToken() (token string, authError error) {
	m.RefreshAuthTokenCalled <- true
	return <-m.RefreshAuthTokenOutput.Token, <-m.RefreshAuthTokenOutput.AuthError
}
