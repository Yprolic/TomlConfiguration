package TomlConfiguration

import (
	"fmt"
	"testing"
)

func TestTagLoader_Load(t *testing.T) {
	tag := TagLoader{}
	s := Server{}
	tag.Load(s)
	fmt.Printf("%+v", s)
}
