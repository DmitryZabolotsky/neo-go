// Code generated by "stringer -type Opcode"; DO NOT EDIT.

package opcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PUSHINT8-0]
	_ = x[PUSHINT16-1]
	_ = x[PUSHINT32-2]
	_ = x[PUSHINT64-3]
	_ = x[PUSHINT128-4]
	_ = x[PUSHINT256-5]
	_ = x[PUSHA-10]
	_ = x[PUSHNULL-11]
	_ = x[PUSHDATA1-12]
	_ = x[PUSHDATA2-13]
	_ = x[PUSHDATA4-14]
	_ = x[PUSHM1-15]
	_ = x[PUSH0-16]
	_ = x[PUSHF-16]
	_ = x[PUSH1-17]
	_ = x[PUSHT-17]
	_ = x[PUSH2-18]
	_ = x[PUSH3-19]
	_ = x[PUSH4-20]
	_ = x[PUSH5-21]
	_ = x[PUSH6-22]
	_ = x[PUSH7-23]
	_ = x[PUSH8-24]
	_ = x[PUSH9-25]
	_ = x[PUSH10-26]
	_ = x[PUSH11-27]
	_ = x[PUSH12-28]
	_ = x[PUSH13-29]
	_ = x[PUSH14-30]
	_ = x[PUSH15-31]
	_ = x[PUSH16-32]
	_ = x[NOP-33]
	_ = x[JMP-34]
	_ = x[JMPL-35]
	_ = x[JMPIF-36]
	_ = x[JMPIFL-37]
	_ = x[JMPIFNOT-38]
	_ = x[JMPIFNOTL-39]
	_ = x[JMPEQ-40]
	_ = x[JMPEQL-41]
	_ = x[JMPNE-42]
	_ = x[JMPNEL-43]
	_ = x[JMPGT-44]
	_ = x[JMPGTL-45]
	_ = x[JMPGE-46]
	_ = x[JMPGEL-47]
	_ = x[JMPLT-48]
	_ = x[JMPLTL-49]
	_ = x[JMPLE-50]
	_ = x[JMPLEL-51]
	_ = x[CALL-52]
	_ = x[CALLL-53]
	_ = x[CALLA-54]
	_ = x[ABORT-55]
	_ = x[ASSERT-56]
	_ = x[THROW-58]
	_ = x[RET-64]
	_ = x[SYSCALL-65]
	_ = x[DEPTH-67]
	_ = x[DROP-69]
	_ = x[NIP-70]
	_ = x[XDROP-72]
	_ = x[CLEAR-73]
	_ = x[DUP-74]
	_ = x[OVER-75]
	_ = x[PICK-77]
	_ = x[TUCK-78]
	_ = x[SWAP-80]
	_ = x[OLDPUSH1-81]
	_ = x[ROT-81]
	_ = x[ROLL-82]
	_ = x[REVERSE3-83]
	_ = x[REVERSE4-84]
	_ = x[REVERSEN-85]
	_ = x[DUPFROMALTSTACK-106]
	_ = x[TOALTSTACK-107]
	_ = x[FROMALTSTACK-108]
	_ = x[CAT-126]
	_ = x[SUBSTR-127]
	_ = x[LEFT-128]
	_ = x[RIGHT-129]
	_ = x[INVERT-144]
	_ = x[AND-145]
	_ = x[OR-146]
	_ = x[XOR-147]
	_ = x[EQUAL-151]
	_ = x[NOTEQUAL-152]
	_ = x[SIGN-153]
	_ = x[ABS-154]
	_ = x[NEGATE-155]
	_ = x[INC-156]
	_ = x[DEC-157]
	_ = x[ADD-158]
	_ = x[SUB-159]
	_ = x[MUL-160]
	_ = x[DIV-161]
	_ = x[MOD-162]
	_ = x[SHL-168]
	_ = x[SHR-169]
	_ = x[NOT-170]
	_ = x[BOOLAND-171]
	_ = x[BOOLOR-172]
	_ = x[NZ-177]
	_ = x[NUMEQUAL-179]
	_ = x[NUMNOTEQUAL-180]
	_ = x[LT-181]
	_ = x[LTE-182]
	_ = x[GT-183]
	_ = x[GTE-184]
	_ = x[MIN-185]
	_ = x[MAX-186]
	_ = x[WITHIN-187]
	_ = x[PACK-192]
	_ = x[UNPACK-193]
	_ = x[NEWARRAY0-194]
	_ = x[NEWARRAY-195]
	_ = x[NEWARRAYT-196]
	_ = x[NEWSTRUCT0-197]
	_ = x[NEWSTRUCT-198]
	_ = x[NEWMAP-200]
	_ = x[SIZE-202]
	_ = x[HASKEY-203]
	_ = x[KEYS-204]
	_ = x[VALUES-205]
	_ = x[PICKITEM-206]
	_ = x[APPEND-207]
	_ = x[SETITEM-208]
	_ = x[REVERSEITEMS-209]
	_ = x[REMOVE-210]
	_ = x[CLEARITEMS-211]
	_ = x[ISNULL-216]
	_ = x[ISTYPE-217]
	_ = x[CONVERT-219]
}

const _Opcode_name = "PUSHINT8PUSHINT16PUSHINT32PUSHINT64PUSHINT128PUSHINT256PUSHAPUSHNULLPUSHDATA1PUSHDATA2PUSHDATA4PUSHM1PUSH0PUSH1PUSH2PUSH3PUSH4PUSH5PUSH6PUSH7PUSH8PUSH9PUSH10PUSH11PUSH12PUSH13PUSH14PUSH15PUSH16NOPJMPJMPLJMPIFJMPIFLJMPIFNOTJMPIFNOTLJMPEQJMPEQLJMPNEJMPNELJMPGTJMPGTLJMPGEJMPGELJMPLTJMPLTLJMPLEJMPLELCALLCALLLCALLAABORTASSERTTHROWRETSYSCALLDEPTHDROPNIPXDROPCLEARDUPOVERPICKTUCKSWAPOLDPUSH1ROLLREVERSE3REVERSE4REVERSENDUPFROMALTSTACKTOALTSTACKFROMALTSTACKCATSUBSTRLEFTRIGHTINVERTANDORXOREQUALNOTEQUALSIGNABSNEGATEINCDECADDSUBMULDIVMODSHLSHRNOTBOOLANDBOOLORNZNUMEQUALNUMNOTEQUALLTLTEGTGTEMINMAXWITHINPACKUNPACKNEWARRAY0NEWARRAYNEWARRAYTNEWSTRUCT0NEWSTRUCTNEWMAPSIZEHASKEYKEYSVALUESPICKITEMAPPENDSETITEMREVERSEITEMSREMOVECLEARITEMSISNULLISTYPECONVERT"

var _Opcode_map = map[Opcode]string{
	0:   _Opcode_name[0:8],
	1:   _Opcode_name[8:17],
	2:   _Opcode_name[17:26],
	3:   _Opcode_name[26:35],
	4:   _Opcode_name[35:45],
	5:   _Opcode_name[45:55],
	10:  _Opcode_name[55:60],
	11:  _Opcode_name[60:68],
	12:  _Opcode_name[68:77],
	13:  _Opcode_name[77:86],
	14:  _Opcode_name[86:95],
	15:  _Opcode_name[95:101],
	16:  _Opcode_name[101:106],
	17:  _Opcode_name[106:111],
	18:  _Opcode_name[111:116],
	19:  _Opcode_name[116:121],
	20:  _Opcode_name[121:126],
	21:  _Opcode_name[126:131],
	22:  _Opcode_name[131:136],
	23:  _Opcode_name[136:141],
	24:  _Opcode_name[141:146],
	25:  _Opcode_name[146:151],
	26:  _Opcode_name[151:157],
	27:  _Opcode_name[157:163],
	28:  _Opcode_name[163:169],
	29:  _Opcode_name[169:175],
	30:  _Opcode_name[175:181],
	31:  _Opcode_name[181:187],
	32:  _Opcode_name[187:193],
	33:  _Opcode_name[193:196],
	34:  _Opcode_name[196:199],
	35:  _Opcode_name[199:203],
	36:  _Opcode_name[203:208],
	37:  _Opcode_name[208:214],
	38:  _Opcode_name[214:222],
	39:  _Opcode_name[222:231],
	40:  _Opcode_name[231:236],
	41:  _Opcode_name[236:242],
	42:  _Opcode_name[242:247],
	43:  _Opcode_name[247:253],
	44:  _Opcode_name[253:258],
	45:  _Opcode_name[258:264],
	46:  _Opcode_name[264:269],
	47:  _Opcode_name[269:275],
	48:  _Opcode_name[275:280],
	49:  _Opcode_name[280:286],
	50:  _Opcode_name[286:291],
	51:  _Opcode_name[291:297],
	52:  _Opcode_name[297:301],
	53:  _Opcode_name[301:306],
	54:  _Opcode_name[306:311],
	55:  _Opcode_name[311:316],
	56:  _Opcode_name[316:322],
	58:  _Opcode_name[322:327],
	64:  _Opcode_name[327:330],
	65:  _Opcode_name[330:337],
	67:  _Opcode_name[337:342],
	69:  _Opcode_name[342:346],
	70:  _Opcode_name[346:349],
	72:  _Opcode_name[349:354],
	73:  _Opcode_name[354:359],
	74:  _Opcode_name[359:362],
	75:  _Opcode_name[362:366],
	77:  _Opcode_name[366:370],
	78:  _Opcode_name[370:374],
	80:  _Opcode_name[374:378],
	81:  _Opcode_name[378:386],
	82:  _Opcode_name[386:390],
	83:  _Opcode_name[390:398],
	84:  _Opcode_name[398:406],
	85:  _Opcode_name[406:414],
	106: _Opcode_name[414:429],
	107: _Opcode_name[429:439],
	108: _Opcode_name[439:451],
	126: _Opcode_name[451:454],
	127: _Opcode_name[454:460],
	128: _Opcode_name[460:464],
	129: _Opcode_name[464:469],
	144: _Opcode_name[469:475],
	145: _Opcode_name[475:478],
	146: _Opcode_name[478:480],
	147: _Opcode_name[480:483],
	151: _Opcode_name[483:488],
	152: _Opcode_name[488:496],
	153: _Opcode_name[496:500],
	154: _Opcode_name[500:503],
	155: _Opcode_name[503:509],
	156: _Opcode_name[509:512],
	157: _Opcode_name[512:515],
	158: _Opcode_name[515:518],
	159: _Opcode_name[518:521],
	160: _Opcode_name[521:524],
	161: _Opcode_name[524:527],
	162: _Opcode_name[527:530],
	168: _Opcode_name[530:533],
	169: _Opcode_name[533:536],
	170: _Opcode_name[536:539],
	171: _Opcode_name[539:546],
	172: _Opcode_name[546:552],
	177: _Opcode_name[552:554],
	179: _Opcode_name[554:562],
	180: _Opcode_name[562:573],
	181: _Opcode_name[573:575],
	182: _Opcode_name[575:578],
	183: _Opcode_name[578:580],
	184: _Opcode_name[580:583],
	185: _Opcode_name[583:586],
	186: _Opcode_name[586:589],
	187: _Opcode_name[589:595],
	192: _Opcode_name[595:599],
	193: _Opcode_name[599:605],
	194: _Opcode_name[605:614],
	195: _Opcode_name[614:622],
	196: _Opcode_name[622:631],
	197: _Opcode_name[631:641],
	198: _Opcode_name[641:650],
	200: _Opcode_name[650:656],
	202: _Opcode_name[656:660],
	203: _Opcode_name[660:666],
	204: _Opcode_name[666:670],
	205: _Opcode_name[670:676],
	206: _Opcode_name[676:684],
	207: _Opcode_name[684:690],
	208: _Opcode_name[690:697],
	209: _Opcode_name[697:709],
	210: _Opcode_name[709:715],
	211: _Opcode_name[715:725],
	216: _Opcode_name[725:731],
	217: _Opcode_name[731:737],
	219: _Opcode_name[737:744],
}

func (i Opcode) String() string {
	if str, ok := _Opcode_map[i]; ok {
		return str
	}
	return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
}
