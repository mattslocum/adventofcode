package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroReturnsErr(t *testing.T) {
	_, err := Roman(0)
	assert.NotNil(t, err, "0 should return an error")
}

func TestNegativeReturnsErr(t *testing.T) {
	_, err := Roman(-1)
	assert.NotNil(t, err, "-1 should return an error")
}

func TestPositiveDoesNotErr(t *testing.T) {
	_, err := Roman(1)
	assert.Nil(t, err, "1 should not error")
}

func TestOver3999DoesError(t *testing.T) {
	_, err := Roman(4123)
	assert.NotNil(t, err, "4123 should return an error")
}

func Test1ReturnsI(t *testing.T) {
	r, _ := Roman(1)
	assert.Equal(t, "I", r, "1 should return I")
}

func Test2ReturnsII(t *testing.T) {
	r, _ := Roman(2)
	assert.Equal(t, "II", r, "2 should return II")
}

func Test3ReturnsIII(t *testing.T) {
	r, _ := Roman(3)
	assert.Equal(t, "III", r, "3 should return III")
}

func Test4ReturnsIV(t *testing.T) {
	r, _ := Roman(4)
	assert.Equal(t, "IV", r, "4 should return IV")
}

func Test5ReturnsV(t *testing.T) {
	r, _ := Roman(5)
	assert.Equal(t, "V", r, "5 should return V")
}

func Test6ReturnsVI(t *testing.T) {
	r, _ := Roman(6)
	assert.Equal(t, "VI", r, "6 should return VI")
}

func Test7ReturnsVII(t *testing.T) {
	r, _ := Roman(7)
	assert.Equal(t, "VII", r, "7 should return VII")
}

func Test8ReturnsVIII(t *testing.T) {
	r, _ := Roman(8)
	assert.Equal(t, "VIII", r, "8 should return VIII")
}

func Test9ReturnsIX(t *testing.T) {
	r, _ := Roman(9)
	assert.Equal(t, "IX", r, "9 should return IX")
}

func Test10ReturnsX(t *testing.T) {
	r, _ := Roman(10)
	assert.Equal(t, "X", r, "10 should return X")
}

func Test20ReturnsXX(t *testing.T) {
	r, _ := Roman(20)
	assert.Equal(t, "XX", r, "20 should return XX")
}

func Test30ReturnsXXX(t *testing.T) {
	r, _ := Roman(30)
	assert.Equal(t, "XXX", r, "30 should return XXX")
}

func Test40ReturnsXL(t *testing.T) {
	r, _ := Roman(40)
	assert.Equal(t, "XL", r, "40 should return XL")
}

func Test100ReturnsC(t *testing.T) {
	r, _ := Roman(100)
	assert.Equal(t, "C", r, "100 should return C")
}

func Test111ReturnsCXI(t *testing.T) {
	r, _ := Roman(111)
	assert.Equal(t, "CXI", r, "111 should return CXI")
}

func Test1111ReturnsMCXI(t *testing.T) {
	r, _ := Roman(1111)
	assert.Equal(t, "MCXI", r, "1111 should return MCXI")
}

func Test3999ReturnsMMMCMXCIX(t *testing.T) {
	r, _ := Roman(3999)
	assert.Equal(t, "MMMCMXCIX", r, "3999 should return MMMCMXCIX")
}
