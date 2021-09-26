package composite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	fmt.Println(got)
	assert.Equal(t, 20, got)
}
