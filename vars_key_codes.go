package buftermio

// key code byte slices
var cr []byte = []byte{uint8(10)}
var del []byte = []byte{uint8(127)}
var enq []byte = []byte{uint8(5)}
var esc []byte = []byte{uint8(27)}
var etb []byte = []byte{uint8(23)}
var fourSpaces []byte = []byte{uint8(32), uint8(32), uint8(32), uint8(32)}
var nak []byte = []byte{uint8(21)}
var openBracket []byte = []byte{uint8(91)}
var soh []byte = []byte{uint8(1)}
var tab []byte = []byte{uint8(9)}
var vt []byte = []byte{uint8(11)}

var up []byte = []byte{uint8(27), uint8(91), uint8(65)}
var down []byte = []byte{uint8(27), uint8(91), uint8(66)}
var right []byte = []byte{uint8(27), uint8(91), uint8(67)}
var left []byte = []byte{uint8(27), uint8(91), uint8(68)}

var space byte = 32
