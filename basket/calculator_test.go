package basket

import (
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x, y := 3, 3
	c := Calculate{}
	actual := c.Add(x, y)
	expected := 6
	if actual != expected {
		t.Error("Eşit değil")
	}

}

//Test table yöntemi

func TestSubstract(t *testing.T) {
	c := Calculate{}
	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 0}, {5, 2, 3}, {5, 2, 1},
	}
	//Tablomuzdaki 3.değerde hata olduğu için konsola bu bilgi girilecektir.
	for _, v := range tables {
		actual := c.Substract(v.x, v.y)
		if actual != v.expected {
			t.Errorf("Calculate.Substract(%d,%d) failed. Expected: %d, Actual %d", v.x, v.y, v.expected, actual)
		}

	}
}

func TestMultiply(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}
	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 4}, {5, 2, 10}, {5, 2, 1},
	}
	// *! Go'da bir değişkenin tipi belli değilse interfce{} verilebilir.
	//Üsteki örneğin hazır kütüphane ile bir başka versionu çıktıyı daha temiz elde edeceğiz.
	for _, v := range tables {
		actual := c.Multiply(v.x, v.y)

		assert.Equal(v.expected, actual)

	}
}

func TestDivide(t *testing.T) {
	assert := asrt.New(t)
	c := Calculate{}
	tables := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{2, 2, 1},
		{4, 2, 2},
		{8, 4, 1},
	}
	for _, v := range tables {
		actual := c.Divide(v.x, v.y)
		assert.Equal(v.expected, actual)
	}
}
