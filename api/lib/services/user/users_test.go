package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestCreateUser(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req1, err := inst.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := appengine.NewContext(req1)

	output, err := Create(c1, &User{
		Name:  `Pedro Costa`,
		Email: `pedrocelsonunes@gmail.com`,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "Pedro Costa", output.Name)
	assert.Equal(t, "pedrocelsonunes@gmail.com", output.Email)
}