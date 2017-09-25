package msg_test

import (
	"testing"

	"github.com/peterj/chat/internal/msg"
)

const succeeded = "\u2713"
const failed = "\u2717"

// TestEconde test that the encoding of the message works
func TestEncode(t *testing.T) {
	m := msg.MSG{
		Name: "1234567890",
		Data: "hello",
	}

	t.Log("Given the need to test encoding.")
	{
		t.Logf("\tTest0: \tWhen checking for basic message")
		{
			data := msg.Encode(m)
			if len(data) != 17 {
				t.Fatalf("\t%s\tShould have the correct number of bytes : exp[17] got[%d]\n", failed, len(data))
			}
			t.Logf("\t%s\tShould have the correct number of bytes\n", succeeded)
		}
	}
}
