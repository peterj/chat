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
				Sender:    "BillKenned",
				Recipient: "JillKenned",
				Data:      "hello",
			},
			length: 27,
		},
		{
			name: "shortname",
			msg: msg.MSG{
				Sender:    "Bill",
				Recipient: "Cory",
				Data:      "helloworld",
			},
			length: 32,
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
				if msg.Sender != tst.msg.Sender {
					t.Fatalf("\t%s\tShould have the correct sender : exp[%v] got[%v]\n", failed, tst.msg.Sender, msg.Sender)
				}
				t.Logf("\t%s\tShould have the correct sender\n", succeeded)

				if msg.Recipient != tst.msg.Recipient {
					t.Fatalf("\t%s\tShould have the correct recipient : exp[%v] got[%v]\n", failed, tst.msg.Recipient, msg.Recipient)
				}
				t.Logf("\t%s\tShould have the correct recipient\n", succeeded)

				if msg.Data != tst.msg.Data {
					t.Fatalf("\t%s\tShould have the correct data : exp[%s] got[%s]\n", failed, tst.msg.Data, msg.Data)
				}
				t.Logf("\t%s\tShould have the correct data\n", succeeded)

			}
		}
	}
}
