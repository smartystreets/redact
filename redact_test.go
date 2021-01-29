package redact

import "testing"

func assertRedaction(t *testing.T, redaction *Redaction, input, expected string) {
	actual := redaction.All(input)
	if actual == expected {
		return
	}
	t.Helper()
	t.Errorf("\n"+
		"Expected: %s\n"+
		"Actual:   %s",
		expected,
		actual,
	)
}

func TestRedactCreditCard(t *testing.T) {
	t.Parallel()

	redaction := New()

	assertRedaction(t, redaction,
		"Blank 5500-0000-0000-0004.",
		"Blank *******************.",
	)
	assertRedaction(t, redaction,
		"36551639043330",
		"**************",
	)
	assertRedaction(t, redaction,
		"4111 1111 1111 1101 111 4556-7375-8689-9855. taco ",
		"*********************** *******************. taco ",
	)
	assertRedaction(t, redaction,
		"10415  80932-1415",
		"10415  80932-1415",
	)
	assertRedaction(t, redaction,
		"3 STE 100 12205-1621     ",
		"3 STE 100 12205-1621     ",
	)
}
func TestRedactEmail(t *testing.T) {
	t.Parallel()

	redaction := New()

	assertRedaction(t, redaction,
		"Blah test@gmail.com, our employee's email is test@gmail. and we have one more which may or not be an email test@test taco",
		"Blah ****@gmail.com, our employee's email is ****@gmail. and we have one more which may or not be an email ****@test taco",
	)
}
func TestRedactPhone(t *testing.T) {
	t.Parallel()

	redaction := New()

	assertRedaction(t, redaction,
		"[1401 WEST MAIN ST  Receiving Dock PO-63308-25   ]",
		"[1401 WEST MAIN ST  Receiving Dock PO-63308-25   ]",
	)
	assertRedaction(t, redaction,
		"Blah 801-111-1111 and  801 111 1111 and (801) 111-1111 +1(801)111-1111 taco",
		"Blah ************ and  ************ and ************** +1************* taco",
	)
	assertRedaction(t, redaction,
		"40512 4618",
		"40512 4618",
	)
}
func TestRedactSSN(t *testing.T) {
	t.Parallel()

	redaction := New()

	assertRedaction(t, redaction,
		"Blah 123-12-1234.",
		"Blah ***********.",
	)
	assertRedaction(t, redaction,
		"123 12 1234 taco",
		"*********** taco",
	)
	assertRedaction(t, redaction,
		" 123-121234 taco",
		" 123-121234 taco",
	)
	assertRedaction(t, redaction,
		"450 900 100",
		"450 900 100",
	)
}

func TestRedactDOB(t *testing.T) {
	t.Parallel()

	redaction := New()
	//assertRedaction(t, redaction,
	//	"Jan 1, 2021",
	//	"*****, 2021",
	//)
	assertRedaction(t, redaction,
		" February 1, 2020",
		" **********, 2020",
	)
	assertRedaction(t, redaction,
		"17-15-13/1",
		"17-15-13/1",
	)
	assertRedaction(t, redaction,
		" 12-15-13/1",
		" ********/1",
	)
	assertRedaction(t, redaction,
		"Blah 12-01-1998 and 12/01/1998 ",
		"Blah ********** and ********** ",
	)
	assertRedaction(t, redaction,
		"29-25-15-12",
		"29-25-15-12",
	)
	assertRedaction(t, redaction,
		"30-12-12",
		"********",
	)
	assertRedaction(t, redaction,
		"30-30-12",
		"30-30-12",
	)
	assertRedaction(t, redaction,
		"8/11/1982",
		"8/11/1982",
	)
	assertRedaction(t, redaction,
		"1/12/21",
		"*******",
	)
	assertRedaction(t, redaction,
		"26-30-62",
		"26-30-62",
	)
	assertRedaction(t, redaction,
		"[5-4-212/80]",
		"[5-4-212/80]",
	)
	assertRedaction(t, redaction,
		"[1230 88-89-92-93     ]",
		"[1230 88-89-92-93     ]",
	)
}
