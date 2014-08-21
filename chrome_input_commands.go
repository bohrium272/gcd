// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the Input commands.
// API Version: 1.1

package gcd

import (
	"github.com/wirepair/gcd/gcdprotogen/types"
)

// add this API domain to ChromeTarget
func (c *ChromeTarget) Input() *ChromeInput {
	if c.input == nil {
		c.input = newChromeInput(c)
	}
	return c.input
}

type ChromeInput struct {
	target *ChromeTarget
}

func newChromeInput(target *ChromeTarget) *ChromeInput {
	c := &ChromeInput{target: target}
	return c
}

// dispatchKeyEvent - Dispatches a key event to the page.
// type - Type of the key event.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
// text - Text as generated by processing a virtual key code with a keyboard layout. Not needed for for <code>keyUp</code> and <code>rawKeyDown</code> events (default: "")
// unmodifiedText - Text that would have been generated by the keyboard if no modifiers were pressed (except for shift). Useful for shortcut (accelerator) key handling (default: "").
// keyIdentifier - Unique key identifier (e.g., 'U+0041') (default: "").
// windowsVirtualKeyCode - Windows virtual key code (default: 0).
// nativeVirtualKeyCode - Native virtual key code (default: 0).
// autoRepeat - Whether the event was generated from auto repeat (default: false).
// isKeypad - Whether the event was generated from the keypad (default: false).
// isSystemKey - Whether the event was a system key event (default: false).
func (c *ChromeInput) DispatchKeyEvent(theType string, modifiers int, timestamp float64, text string, unmodifiedText string, keyIdentifier string, windowsVirtualKeyCode int, nativeVirtualKeyCode int, autoRepeat bool, isKeypad bool, isSystemKey bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 11)
	paramRequest["type"] = theType
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	paramRequest["text"] = text
	paramRequest["unmodifiedText"] = unmodifiedText
	paramRequest["keyIdentifier"] = keyIdentifier
	paramRequest["windowsVirtualKeyCode"] = windowsVirtualKeyCode
	paramRequest["nativeVirtualKeyCode"] = nativeVirtualKeyCode
	paramRequest["autoRepeat"] = autoRepeat
	paramRequest["isKeypad"] = isKeypad
	paramRequest["isSystemKey"] = isSystemKey
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchKeyEvent", Params: paramRequest})
}

// dispatchMouseEvent - Dispatches a mouse event to the page.
// type - Type of the mouse event.
// x - X coordinate of the event relative to the main frame's viewport.
// y - Y coordinate of the event relative to the main frame's viewport. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
// button - Mouse button (default: "none").
// clickCount - Number of times the mouse button was clicked (default: 0).
// deviceSpace - If true, x and y are given in dip wrt current viewport.
func (c *ChromeInput) DispatchMouseEvent(theType string, x int, y int, modifiers int, timestamp float64, button string, clickCount int, deviceSpace bool) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 8)
	paramRequest["type"] = theType
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	paramRequest["button"] = button
	paramRequest["clickCount"] = clickCount
	paramRequest["deviceSpace"] = deviceSpace
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchMouseEvent", Params: paramRequest})
}

// dispatchTouchEvent - Dispatches a touch event to the page.
// type - Type of the touch event.
// touchPoints - Touch points.
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
func (c *ChromeInput) DispatchTouchEvent(theType string, touchPoints []*types.ChromeInputTouchPoint, modifiers int, timestamp float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 4)
	paramRequest["type"] = theType
	paramRequest["touchPoints"] = touchPoints
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchTouchEvent", Params: paramRequest})
}

// dispatchGestureEvent - Dispatches a gesture event to the page.
// type - Type of the gesture event.
// x - X coordinate relative to the screen's viewport.
// y - Y coordinate relative to the screen's viewport.
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
// deltaX - Delta X where apllies.
// deltaY - Delta Y where apllies.
// pinchScale - Pinch scale.
func (c *ChromeInput) DispatchGestureEvent(theType string, x int, y int, timestamp float64, deltaX int, deltaY int, pinchScale float64) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 7)
	paramRequest["type"] = theType
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["timestamp"] = timestamp
	paramRequest["deltaX"] = deltaX
	paramRequest["deltaY"] = deltaY
	paramRequest["pinchScale"] = pinchScale
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.dispatchGestureEvent", Params: paramRequest})
}

// emulateTouchFromMouseEvent - Emulates touch event from the mouse event parameters.
// type - Type of the mouse event.
// x - X coordinate of the mouse pointer in DIP.
// y - Y coordinate of the mouse pointer in DIP.
// deltaX - X delta in DIP for mouse wheel event (default: 0).
// deltaY - Y delta in DIP for mouse wheel event (default: 0).
// modifiers - Bit field representing pressed modifier keys. Alt=1, Ctrl=2, Meta/Command=4, Shift=8 (default: 0).
// timestamp - Time at which the event occurred. Measured in UTC time in seconds since January 1, 1970 (default: current time).
// button - Mouse button (default: "none").
// clickCount - Number of times the mouse button was clicked (default: 0).
func (c *ChromeInput) EmulateTouchFromMouseEvent(theType string, x int, y int, deltaX float64, deltaY float64, modifiers int, timestamp float64, button string, clickCount int) (*ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 9)
	paramRequest["type"] = theType
	paramRequest["x"] = x
	paramRequest["y"] = y
	paramRequest["deltaX"] = deltaX
	paramRequest["deltaY"] = deltaY
	paramRequest["modifiers"] = modifiers
	paramRequest["timestamp"] = timestamp
	paramRequest["button"] = button
	paramRequest["clickCount"] = clickCount
	return sendDefaultRequest(c.target.sendCh, &ParamRequest{Id: c.target.getId(), Method: "Input.emulateTouchFromMouseEvent", Params: paramRequest})
}
