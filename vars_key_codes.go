package buftermio

// key code byte slices
var carriageReturn []byte = []byte{uint8(10)}
var ctrlA []byte = []byte{uint8(17), uint8(97)}
var ctrlE []byte = []byte{uint8(17), uint8(101)}
var ctrlW []byte = []byte{uint8(17), uint8(119)}
var del []byte = []byte{uint8(127)}
var esc []byte = []byte{uint8(27)}
var fourSpaces []byte = []byte{uint8(32), uint8(32), uint8(32), uint8(32)}
var openBracket []byte = []byte{uint8(91)}
var tab []byte = []byte{uint8(9)}

var upArrow []byte = []byte{uint8(27), uint8(91), uint8(65)}
var downArrow []byte = []byte{uint8(27), uint8(91), uint8(66)}
var rightArrow []byte = []byte{uint8(27), uint8(91), uint8(67)}
var leftArrow []byte = []byte{uint8(27), uint8(91), uint8(68)}
