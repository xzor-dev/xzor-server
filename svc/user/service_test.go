package user_test

import "testing"

import "github.com/xzor-dev/xzor-server/svc/user"

func TestUserService(t *testing.T) {
	u, err := user.New()
	if err != nil {
		t.Fatalf("%v", err)
	}

	s := user.NewService()
	err = s.SaveUser(u)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
