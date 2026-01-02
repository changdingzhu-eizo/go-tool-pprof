package data

import "testing"

const url = "https://github.com/changdingzhu-eizo"

func TestAdd(t *testing.T) {
	s := Add(url)
	if s == "" {
		t.Error("Test.Add error!")
	}
}

func BechmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}