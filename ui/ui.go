package ui

// Ui is the component of Otto responsible for reading/writing to the
// console.
//
// All Ui implementations MUST expect colorstring[1] style inputs. If
// the output interface doesn't support colors, these must be stripped.
// The StripColors helper in this package can be used to do this. For
// terminals, the Colorize helper in this package can be used.
//
// [1]: github.com/mitchellh/colorstring
type Ui interface {
	// Header, Message, and Raw are all methods for outputting messages
	// to the Ui, all with different styles. Header and Message should
	// be used liberally, and Raw should be scarcely used if possible.
	//
	// Header outputs the message with a style denoting it is a sectional
	// message. An example: "==> TEXT" might be header text.
	//
	// Message outputs the message with a style, but one that is less
	// important looking than Header. Example: "    TEXT" might be
	// the message text, prefixed with spaces so that it lines up with
	// header items.
	//
	// Raw outputs a message with no styling.
	Header(string)
	Message(string)
	Raw(string)
}
