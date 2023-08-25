package buftermio

// key byte slices, used as constants (not immutable)
var DELETE []byte = []byte{uint8(127)}
var UPARROW []byte = []byte{uint8(27), uint8(91), uint8(65)}
var DOWNARROW []byte = []byte{uint8(27), uint8(91), uint8(66)}
var RIGHTARROW []byte = []byte{uint8(27), uint8(91), uint8(67)}
var LEFTARROW []byte = []byte{uint8(27), uint8(91), uint8(68)}
var TAB []byte = []byte{uint8(9)}
var FOURSPACES []byte = []byte{uint8(32), uint8(32), uint8(32), uint8(32)}
var ESCAPE []byte = []byte{uint8(27)}
var OPENBRACKET []byte = []byte{uint8(91)}
var CARRIAGERETURN []byte = []byte{uint8(10)}
