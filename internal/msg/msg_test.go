package msg_test

import (
	"testing"

	"github.com/peterj/chat/internal/msg"
)

const succeeded = "\u2713"
const failed = "\u2717"

// TestEconde test that the encoding of the message works
func TestEncode(t *testing.T) {
	tt := []struct {
		name   string
		msg    msg.MSG
		length int
	}{
		{
			name: "length",
			msg: msg.MSG{
				Name: "1234567890",
				Data: "hello",
			},
			length: 17,
		},
		{
			name: "shortname",
			msg: msg.MSG{
				Name: "012345",
				Data: "helloworld",
			},
			length: 22,
		},
	}

	t.Log("Given the need to test encoding.")
	{
		for i, tst := range tt {
			t.Logf("\tTest %d: \t%s", i, tst.name)
			{
				data := msg.Encode(tst.msg)
				if len(data) != tst.length {
					t.Fatalf("\t%s\tShould have the correct number of bytes : exp[%d] got[%d]\n", failed, tst.length, len(data))
				}
				t.Logf("\t%s\tShould have the correct number of bytes\n", succeeded)

				msg := msg.Decode(data)
				if msg.Name != tst.msg.Name {
					t.Fatalf("\t%s\tShould have the correct name : exp[%v] got[%v]\n", failed, []byte(tst.msg.Name), []byte(msg.Name))
				}
				t.Logf("\t%s\tShould have the correct name\n", succeeded)

				if msg.Data != tst.msg.Data {
					t.Fatalf("\t%s\tShould have the correct data : exp[%s] got[%s]\n", failed, tst.msg.Data, msg.Data)
				}
				t.Logf("\t%s\tShould have the correct data\n", succeeded)

			}
		}
	}
}
