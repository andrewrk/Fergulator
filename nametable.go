package main

const (
	MirroringVertical = iota
	MirroringHorizontal
	MirroringSingleUpper
	MirroringSingleLower
)

type Nametable struct {
	Mirroring     int
	LogicalTables [4]*[0x400]uint8
	Nametable0    [0x400]uint8
	Nametable1    [0x400]uint8
}

func (n *Nametable) SetMirroring(m int) {
	n.Mirroring = m

	switch n.Mirroring {
	case MirroringHorizontal:
		n.LogicalTables[0] = &n.Nametable0
		n.LogicalTables[1] = &n.Nametable0
		n.LogicalTables[2] = &n.Nametable1
		n.LogicalTables[3] = &n.Nametable1
	case MirroringVertical:
		n.LogicalTables[0] = &n.Nametable0
		n.LogicalTables[1] = &n.Nametable1
		n.LogicalTables[2] = &n.Nametable0
		n.LogicalTables[3] = &n.Nametable1
	case MirroringSingleUpper:
		n.LogicalTables[0] = &n.Nametable0
		n.LogicalTables[1] = &n.Nametable0
		n.LogicalTables[2] = &n.Nametable0
		n.LogicalTables[3] = &n.Nametable0
	case MirroringSingleLower:
		n.LogicalTables[0] = &n.Nametable1
		n.LogicalTables[1] = &n.Nametable1
		n.LogicalTables[2] = &n.Nametable1
		n.LogicalTables[3] = &n.Nametable1
	}
}

func (n *Nametable) writeNametableData(a int, v uint8) {
	n.LogicalTables[(a&0xC00)>>10][a&0x3FF] = v
}

func (n *Nametable) readNametableData(a int) uint8 {
	return n.LogicalTables[(a&0xC00)>>10][a&0x3FF]
}
