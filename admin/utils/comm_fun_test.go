package utils

import (
	"fmt"
	"testing"
)


func TestComm(t *testing.T) {
	t.Run("token create", getPageLimit)

}

// 生成token
func getPageLimit(t *testing.T) {

	//num := GetPageLimit(12, 5)
	c := 10 / 5
	fmt.Println(c)

	m := 10 % 5
	fmt.Println(m)
}
