package TomlConfiguration

import (
	"fmt"
	"testing"
)

func TestTagLoader_Load(t *testing.T) {
	tag := TagLoader{}
	s := Server{}
	tag.load(&s)
	fmt.Printf("%+v", s)
}
