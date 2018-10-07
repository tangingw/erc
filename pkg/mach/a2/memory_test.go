package a2

import (
	"github.com/pevans/erc/pkg/mach"
	"github.com/stretchr/testify/assert"
)

func (s *a2Suite) TestZeroPageRead() {
	addr := mach.DByte(0x123)
	cases := []struct {
		mode int
		main mach.Byte
		aux  mach.Byte
		want mach.Byte
	}{
		{BankAuxiliary, 0x1, 0x2, 0x2},
		{0, 0x3, 0x2, 0x3},
	}

	for _, c := range cases {
		s.comp.Main.Set(addr, c.main)
		s.comp.Aux.Set(addr, c.aux)
		s.comp.BankMode = c.mode

		assert.Equal(s.T(), c.want, s.comp.Get(addr))
	}
}

func (s *a2Suite) TestZeroPageWrite() {
	addr := mach.DByte(0x123)
	cases := []struct {
		mode int
		main mach.Byte
		aux  mach.Byte
		want mach.Byte
	}{
		{BankAuxiliary, 0x0, 0x2, 0x2},
		{0, 0x3, 0x0, 0x3},
	}

	for _, c := range cases {
		s.comp.Main.Set(addr, 0x0)
		s.comp.Aux.Set(addr, 0x0)

		s.comp.BankMode = c.mode
		s.comp.Set(addr, c.want)

		assert.Equal(s.T(), c.main, s.comp.Main.Get(addr))
		assert.Equal(s.T(), c.aux, s.comp.Aux.Get(addr))
	}
}

func (s *a2Suite) TestMemorySwitchIsSetR() {
	var addr mach.DByte

	cases := []struct {
		memMode int
		flag    int
		want    mach.Byte
	}{
		{0x0, MemPage2, 0x0},
		{0x0, 0, 0x0},
		{MemPage2 | MemHires, MemPage2, 0x80},
		{MemHires, MemPage2, 0x0},
	}

	for _, c := range cases {
		s.comp.MemMode = c.memMode
		fn := s.comp.memorySwitchIsSetR(c.flag)

		assert.Equal(s.T(), c.want, fn(s.comp, addr))
	}
}

func (s *a2Suite) TestMemorySwitchSetR() {
	var addr mach.DByte

	cases := []struct {
		startMode int
		flag      int
		endMode   int
		want      mach.Byte
	}{
		{0, MemPage2, MemPage2, 0x80},
		{0, 0, 0, 0x80},
		{MemPage2, MemPage2, MemPage2, 0x80},
		{MemHires, MemPage2, MemHires | MemPage2, 0x80},
	}

	for _, c := range cases {
		s.comp.MemMode = c.startMode
		fn := s.comp.memorySwitchSetR(c.flag)

		assert.Equal(s.T(), c.want, fn(s.comp, addr))
		assert.Equal(s.T(), c.endMode, s.comp.MemMode)
	}
}

func (s *a2Suite) TestMemorySwitchUnsetR() {
	var addr mach.DByte

	cases := []struct {
		startMode int
		flag      int
		endMode   int
		want      mach.Byte
	}{
		{0, MemPage2, 0, 0x0},
		{0, 0, 0, 0x0},
		{MemPage2, MemPage2, 0, 0x0},
		{MemHires | MemPage2, MemPage2, MemHires, 0x0},
	}

	for _, c := range cases {
		s.comp.MemMode = c.startMode
		fn := s.comp.memorySwitchUnsetR(c.flag)

		assert.Equal(s.T(), c.want, fn(s.comp, addr))
		assert.Equal(s.T(), c.endMode, s.comp.MemMode)
	}
}

func (s *a2Suite) TestMemorySwitchSetW() {
	var addr mach.DByte

	cases := []struct {
		startMode int
		flag      int
		endMode   int
	}{
		{0, MemPage2, MemPage2},
		{0, 0, 0},
		{MemPage2, MemPage2, MemPage2},
		{MemPage2, MemHires, MemHires | MemPage2},
	}

	for _, c := range cases {
		s.comp.MemMode = c.startMode
		fn := s.comp.memorySwitchSetW(c.flag)

		fn(s.comp, addr, 0)
		assert.Equal(s.T(), c.endMode, s.comp.MemMode)
	}
}

func (s *a2Suite) TestMemorySwitchUnsetW() {
	var addr mach.DByte

	cases := []struct {
		startMode int
		flag      int
		endMode   int
	}{
		{0, MemPage2, 0},
		{0, 0, 0},
		{MemPage2, MemPage2, 0},
		{MemHires | MemPage2, MemPage2, MemHires},
	}

	for _, c := range cases {
		s.comp.MemMode = c.startMode
		fn := s.comp.memorySwitchUnsetW(c.flag)

		fn(s.comp, addr, 0)
		assert.Equal(s.T(), c.endMode, s.comp.MemMode)
	}
}
