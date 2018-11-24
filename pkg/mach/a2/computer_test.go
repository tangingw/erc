package a2

import (
	"testing"

	"github.com/pevans/erc/pkg/mach/a2/disk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type a2Suite struct {
	suite.Suite

	comp  *Computer
	drive *disk.Drive
}

func (s *a2Suite) SetupSuite() {
	s.comp = NewComputer()
	s.drive = disk.NewDrive()
}

func (s *a2Suite) SetupTest() {
	_ = s.comp.Boot()
}

func TestNewComputer(t *testing.T) {
	comp := NewComputer()

	assert.NotEqual(t, nil, comp.Main)
	assert.NotEqual(t, nil, comp.Aux)
	assert.NotEqual(t, nil, comp.ROM)
	assert.NotEqual(t, nil, comp.CPU)
	assert.NotEqual(t, nil, comp.CPU.RMem)
	assert.NotEqual(t, nil, comp.CPU.WMem)
	assert.NotEqual(t, nil, comp.RMap)
	assert.NotEqual(t, nil, comp.WMap)
}

func TestA2Suite(t *testing.T) {
	suite.Run(t, new(a2Suite))
}
