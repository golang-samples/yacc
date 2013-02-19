
//line calc.y:8

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var regs = make([]int, 26)
var base int


//line calc.y:25
type CalcSymType struct{
	yys int
	val int
}

const DIGIT = 57346
const LETTER = 57347
const UMINUS = 57348

var CalcToknames = []string{
	"DIGIT",
	"LETTER",
	" |",
	" &",
	" +",
	" -",
	" *",
	" /",
	" %",
	"UMINUS",
}
var CalcStatenames = []string{}

const CalcEofCode = 1
const CalcErrCode = 2
const CalcMaxDepth = 200

//line calc.y:94
      /*  start  of  programs  */

type CalcLex struct {
	s string
	pos int
}


func (l *CalcLex) Lex(lval *CalcSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c) - 'a'
		return LETTER
	}
	return int(c)
}

func (l *CalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func main() {
	fi := bufio.NewReader(os.NewFile(0, "stdin"))

	for {
		var eqn string
		var ok bool

		fmt.Printf("equation: ")
		if eqn, ok = readline(fi); ok {
			CalcParse(&CalcLex{s: eqn})
		} else {
			break
		}
	}
}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}

//line yacctab:1
var CalcExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const CalcNprod = 18
const CalcPrivate = 57344

var CalcTokenNames []string
var CalcStates []string

const CalcLast = 54

var CalcAct = []int{

	3, 10, 11, 12, 13, 14, 18, 20, 21, 17,
	9, 22, 23, 24, 25, 26, 27, 28, 29, 16,
	15, 10, 11, 12, 13, 14, 8, 19, 8, 4,
	30, 6, 2, 6, 1, 12, 13, 14, 5, 7,
	5, 16, 15, 10, 11, 12, 13, 14, 15, 10,
	11, 12, 13, 14,
}
var CalcPact = []int{

	-1000, 24, -4, 35, -6, 22, 22, 4, -1000, -1000,
	22, 22, 22, 22, 22, 22, 22, 22, 13, -1000,
	-1000, -1000, 25, 25, -1000, -1000, -1000, -7, 41, 35,
	-1000,
}
var CalcPgo = []int{

	0, 0, 39, 34, 32,
}
var CalcR1 = []int{

	0, 3, 3, 4, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2,
}
var CalcR2 = []int{

	0, 0, 3, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 1, 1, 1, 2,
}
var CalcChk = []int{

	-1000, -3, -4, -1, 5, 16, 9, -2, 4, 14,
	8, 9, 10, 11, 12, 7, 6, 15, -1, 5,
	-1, 4, -1, -1, -1, -1, -1, -1, -1, -1,
	17,
}
var CalcDef = []int{

	1, -2, 0, 3, 14, 0, 0, 15, 16, 2,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 14,
	13, 17, 6, 7, 8, 9, 10, 11, 12, 4,
	5,
}
var CalcTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 12, 7, 3,
	16, 17, 10, 8, 3, 9, 3, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 15, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 6,
}
var CalcTok2 = []int{

	2, 3, 4, 5, 13,
}
var CalcTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var CalcDebug = 0

type CalcLexer interface {
	Lex(lval *CalcSymType) int
	Error(s string)
}

const CalcFlag = -1000

func CalcTokname(c int) string {
	if c > 0 && c <= len(CalcToknames) {
		if CalcToknames[c-1] != "" {
			return CalcToknames[c-1]
		}
	}
	return fmt.Sprintf("tok-%v", c)
}

func CalcStatname(s int) string {
	if s >= 0 && s < len(CalcStatenames) {
		if CalcStatenames[s] != "" {
			return CalcStatenames[s]
		}
	}
	return fmt.Sprintf("state-%v", s)
}

func Calclex1(lex CalcLexer, lval *CalcSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = CalcTok1[0]
		goto out
	}
	if char < len(CalcTok1) {
		c = CalcTok1[char]
		goto out
	}
	if char >= CalcPrivate {
		if char < CalcPrivate+len(CalcTok2) {
			c = CalcTok2[char-CalcPrivate]
			goto out
		}
	}
	for i := 0; i < len(CalcTok3); i += 2 {
		c = CalcTok3[i+0]
		if c == char {
			c = CalcTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = CalcTok2[1] /* unknown char */
	}
	if CalcDebug >= 3 {
		fmt.Printf("lex %U %s\n", uint(char), CalcTokname(c))
	}
	return c
}

func CalcParse(Calclex CalcLexer) int {
	var Calcn int
	var Calclval CalcSymType
	var CalcVAL CalcSymType
	CalcS := make([]CalcSymType, CalcMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Calcstate := 0
	Calcchar := -1
	Calcp := -1
	goto Calcstack

ret0:
	return 0

ret1:
	return 1

Calcstack:
	/* put a state and value onto the stack */
	if CalcDebug >= 4 {
		fmt.Printf("char %v in %v\n", CalcTokname(Calcchar), CalcStatname(Calcstate))
	}

	Calcp++
	if Calcp >= len(CalcS) {
		nyys := make([]CalcSymType, len(CalcS)*2)
		copy(nyys, CalcS)
		CalcS = nyys
	}
	CalcS[Calcp] = CalcVAL
	CalcS[Calcp].yys = Calcstate

Calcnewstate:
	Calcn = CalcPact[Calcstate]
	if Calcn <= CalcFlag {
		goto Calcdefault /* simple state */
	}
	if Calcchar < 0 {
		Calcchar = Calclex1(Calclex, &Calclval)
	}
	Calcn += Calcchar
	if Calcn < 0 || Calcn >= CalcLast {
		goto Calcdefault
	}
	Calcn = CalcAct[Calcn]
	if CalcChk[Calcn] == Calcchar { /* valid shift */
		Calcchar = -1
		CalcVAL = Calclval
		Calcstate = Calcn
		if Errflag > 0 {
			Errflag--
		}
		goto Calcstack
	}

Calcdefault:
	/* default state action */
	Calcn = CalcDef[Calcstate]
	if Calcn == -2 {
		if Calcchar < 0 {
			Calcchar = Calclex1(Calclex, &Calclval)
		}

		/* look through exception table */
		xi := 0
		for {
			if CalcExca[xi+0] == -1 && CalcExca[xi+1] == Calcstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Calcn = CalcExca[xi+0]
			if Calcn < 0 || Calcn == Calcchar {
				break
			}
		}
		Calcn = CalcExca[xi+1]
		if Calcn < 0 {
			goto ret0
		}
	}
	if Calcn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Calclex.Error("syntax error")
			Nerrs++
			if CalcDebug >= 1 {
				fmt.Printf("%s", CalcStatname(Calcstate))
				fmt.Printf("saw %s\n", CalcTokname(Calcchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Calcp >= 0 {
				Calcn = CalcPact[CalcS[Calcp].yys] + CalcErrCode
				if Calcn >= 0 && Calcn < CalcLast {
					Calcstate = CalcAct[Calcn] /* simulate a shift of "error" */
					if CalcChk[Calcstate] == CalcErrCode {
						goto Calcstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if CalcDebug >= 2 {
					fmt.Printf("error recovery pops state %d\n", CalcS[Calcp].yys)
				}
				Calcp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if CalcDebug >= 2 {
				fmt.Printf("error recovery discards %s\n", CalcTokname(Calcchar))
			}
			if Calcchar == CalcEofCode {
				goto ret1
			}
			Calcchar = -1
			goto Calcnewstate /* try again in the same state */
		}
	}

	/* reduction by production Calcn */
	if CalcDebug >= 2 {
		fmt.Printf("reduce %v in:\n\t%v\n", Calcn, CalcStatname(Calcstate))
	}

	Calcnt := Calcn
	Calcpt := Calcp
	_ = Calcpt // guard against "declared and not used"

	Calcp -= CalcR2[Calcn]
	CalcVAL = CalcS[Calcp+1]

	/* consult goto table to find next state */
	Calcn = CalcR1[Calcn]
	Calcg := CalcPgo[Calcn]
	Calcj := Calcg + CalcS[Calcp].yys + 1

	if Calcj >= CalcLast {
		Calcstate = CalcAct[Calcg]
	} else {
		Calcstate = CalcAct[Calcj]
		if CalcChk[Calcstate] != -Calcn {
			Calcstate = CalcAct[Calcg]
		}
	}
	// dummy call; replaced with literal code
	switch Calcnt {

	case 3:
		//line calc.y:49
		{
				fmt.Printf( "%d\n", CalcS[Calcpt-0].val );
			}
	case 4:
		//line calc.y:53
		{
				regs[CalcS[Calcpt-2].val]  =  CalcS[Calcpt-0].val
			}
	case 5:
		//line calc.y:59
		{ CalcVAL.val  =  CalcS[Calcpt-1].val }
	case 6:
		//line calc.y:61
		{ CalcVAL.val  =  CalcS[Calcpt-2].val + CalcS[Calcpt-0].val }
	case 7:
		//line calc.y:63
		{ CalcVAL.val  =  CalcS[Calcpt-2].val - CalcS[Calcpt-0].val }
	case 8:
		//line calc.y:65
		{ CalcVAL.val  =  CalcS[Calcpt-2].val * CalcS[Calcpt-0].val }
	case 9:
		//line calc.y:67
		{ CalcVAL.val  =  CalcS[Calcpt-2].val / CalcS[Calcpt-0].val }
	case 10:
		//line calc.y:69
		{ CalcVAL.val  =  CalcS[Calcpt-2].val % CalcS[Calcpt-0].val }
	case 11:
		//line calc.y:71
		{ CalcVAL.val  =  CalcS[Calcpt-2].val & CalcS[Calcpt-0].val }
	case 12:
		//line calc.y:73
		{ CalcVAL.val  =  CalcS[Calcpt-2].val | CalcS[Calcpt-0].val }
	case 13:
		//line calc.y:75
		{ CalcVAL.val  = -CalcS[Calcpt-0].val  }
	case 14:
		//line calc.y:77
		{ CalcVAL.val  = regs[CalcS[Calcpt-0].val] }
	case 15:
		CalcVAL.val = CalcS[Calcpt-0].val
	case 16:
		//line calc.y:82
		{
				CalcVAL.val = CalcS[Calcpt-0].val;
				if CalcS[Calcpt-0].val==0 {
					base = 8
				} else {
					base = 10
				}
			}
	case 17:
		//line calc.y:91
		{ CalcVAL.val = base * CalcS[Calcpt-1].val + CalcS[Calcpt-0].val }
	}
	goto Calcstack /* stack new state and value */
}
