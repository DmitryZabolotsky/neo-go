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
	_ = x[OLDPUSH1-81]
	_ = x[RET-102]
	_ = x[APPCALL-103]
	_ = x[SYSCALL-104]
	_ = x[TAILCALL-105]
	_ = x[DUPFROMALTSTACK-106]
	_ = x[TOALTSTACK-107]
	_ = x[FROMALTSTACK-108]
	_ = x[XDROP-109]
	_ = x[XSWAP-114]
	_ = x[XTUCK-115]
	_ = x[DEPTH-116]
	_ = x[DROP-117]
	_ = x[DUP-118]
	_ = x[NIP-119]
	_ = x[OVER-120]
	_ = x[PICK-121]
	_ = x[ROLL-122]
	_ = x[ROT-123]
	_ = x[SWAP-124]
	_ = x[TUCK-125]
	_ = x[CAT-126]
	_ = x[SUBSTR-127]
	_ = x[LEFT-128]
	_ = x[RIGHT-129]
	_ = x[INVERT-144]
	_ = x[AND-145]
	_ = x[OR-146]
	_ = x[XOR-147]
	_ = x[EQUAL-151]
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
	_ = x[THROW-240]
	_ = x[THROWIFNOT-241]
}

const _Opcode_name = "PUSHINT8PUSHINT16PUSHINT32PUSHINT64PUSHINT128PUSHINT256PUSHNULLPUSHDATA1PUSHDATA2PUSHDATA4PUSHM1PUSH0PUSH1PUSH2PUSH3PUSH4PUSH5PUSH6PUSH7PUSH8PUSH9PUSH10PUSH11PUSH12PUSH13PUSH14PUSH15PUSH16NOPJMPJMPLJMPIFJMPIFLJMPIFNOTJMPIFNOTLJMPEQJMPEQLJMPNEJMPNELJMPGTJMPGTLJMPGEJMPGELJMPLTJMPLTLJMPLEJMPLELCALLCALLLOLDPUSH1RETAPPCALLSYSCALLTAILCALLDUPFROMALTSTACKTOALTSTACKFROMALTSTACKXDROPXSWAPXTUCKDEPTHDROPDUPNIPOVERPICKROLLROTSWAPTUCKCATSUBSTRLEFTRIGHTINVERTANDORXOREQUALSIGNABSNEGATEINCDECADDSUBMULDIVMODSHLSHRNOTBOOLANDBOOLORNZNUMEQUALNUMNOTEQUALLTLTEGTGTEMINMAXWITHINPACKUNPACKNEWARRAY0NEWARRAYNEWARRAYTNEWSTRUCT0NEWSTRUCTNEWMAPSIZEHASKEYKEYSVALUESPICKITEMAPPENDSETITEMREVERSEITEMSREMOVECLEARITEMSISNULLISTYPECONVERTTHROWTHROWIFNOT"

var _Opcode_map = map[Opcode]string{
	0:   _Opcode_name[0:8],
	1:   _Opcode_name[8:17],
	2:   _Opcode_name[17:26],
	3:   _Opcode_name[26:35],
	4:   _Opcode_name[35:45],
	5:   _Opcode_name[45:55],
	11:  _Opcode_name[55:63],
	12:  _Opcode_name[63:72],
	13:  _Opcode_name[72:81],
	14:  _Opcode_name[81:90],
	15:  _Opcode_name[90:96],
	16:  _Opcode_name[96:101],
	17:  _Opcode_name[101:106],
	18:  _Opcode_name[106:111],
	19:  _Opcode_name[111:116],
	20:  _Opcode_name[116:121],
	21:  _Opcode_name[121:126],
	22:  _Opcode_name[126:131],
	23:  _Opcode_name[131:136],
	24:  _Opcode_name[136:141],
	25:  _Opcode_name[141:146],
	26:  _Opcode_name[146:152],
	27:  _Opcode_name[152:158],
	28:  _Opcode_name[158:164],
	29:  _Opcode_name[164:170],
	30:  _Opcode_name[170:176],
	31:  _Opcode_name[176:182],
	32:  _Opcode_name[182:188],
	33:  _Opcode_name[188:191],
	34:  _Opcode_name[191:194],
	35:  _Opcode_name[194:198],
	36:  _Opcode_name[198:203],
	37:  _Opcode_name[203:209],
	38:  _Opcode_name[209:217],
	39:  _Opcode_name[217:226],
	40:  _Opcode_name[226:231],
	41:  _Opcode_name[231:237],
	42:  _Opcode_name[237:242],
	43:  _Opcode_name[242:248],
	44:  _Opcode_name[248:253],
	45:  _Opcode_name[253:259],
	46:  _Opcode_name[259:264],
	47:  _Opcode_name[264:270],
	48:  _Opcode_name[270:275],
	49:  _Opcode_name[275:281],
	50:  _Opcode_name[281:286],
	51:  _Opcode_name[286:292],
	52:  _Opcode_name[292:296],
	53:  _Opcode_name[296:301],
	81:  _Opcode_name[301:309],
	102: _Opcode_name[309:312],
	103: _Opcode_name[312:319],
	104: _Opcode_name[319:326],
	105: _Opcode_name[326:334],
	106: _Opcode_name[334:349],
	107: _Opcode_name[349:359],
	108: _Opcode_name[359:371],
	109: _Opcode_name[371:376],
	114: _Opcode_name[376:381],
	115: _Opcode_name[381:386],
	116: _Opcode_name[386:391],
	117: _Opcode_name[391:395],
	118: _Opcode_name[395:398],
	119: _Opcode_name[398:401],
	120: _Opcode_name[401:405],
	121: _Opcode_name[405:409],
	122: _Opcode_name[409:413],
	123: _Opcode_name[413:416],
	124: _Opcode_name[416:420],
	125: _Opcode_name[420:424],
	126: _Opcode_name[424:427],
	127: _Opcode_name[427:433],
	128: _Opcode_name[433:437],
	129: _Opcode_name[437:442],
	144: _Opcode_name[442:448],
	145: _Opcode_name[448:451],
	146: _Opcode_name[451:453],
	147: _Opcode_name[453:456],
	151: _Opcode_name[456:461],
	153: _Opcode_name[461:465],
	154: _Opcode_name[465:468],
	155: _Opcode_name[468:474],
	156: _Opcode_name[474:477],
	157: _Opcode_name[477:480],
	158: _Opcode_name[480:483],
	159: _Opcode_name[483:486],
	160: _Opcode_name[486:489],
	161: _Opcode_name[489:492],
	162: _Opcode_name[492:495],
	168: _Opcode_name[495:498],
	169: _Opcode_name[498:501],
	170: _Opcode_name[501:504],
	171: _Opcode_name[504:511],
	172: _Opcode_name[511:517],
	177: _Opcode_name[517:519],
	179: _Opcode_name[519:527],
	180: _Opcode_name[527:538],
	181: _Opcode_name[538:540],
	182: _Opcode_name[540:543],
	183: _Opcode_name[543:545],
	184: _Opcode_name[545:548],
	185: _Opcode_name[548:551],
	186: _Opcode_name[551:554],
	187: _Opcode_name[554:560],
	192: _Opcode_name[560:564],
	193: _Opcode_name[564:570],
	194: _Opcode_name[570:579],
	195: _Opcode_name[579:587],
	196: _Opcode_name[587:596],
	197: _Opcode_name[596:606],
	198: _Opcode_name[606:615],
	200: _Opcode_name[615:621],
	202: _Opcode_name[621:625],
	203: _Opcode_name[625:631],
	204: _Opcode_name[631:635],
	205: _Opcode_name[635:641],
	206: _Opcode_name[641:649],
	207: _Opcode_name[649:655],
	208: _Opcode_name[655:662],
	209: _Opcode_name[662:674],
	210: _Opcode_name[674:680],
	211: _Opcode_name[680:690],
	216: _Opcode_name[690:696],
	217: _Opcode_name[696:702],
	219: _Opcode_name[702:709],
	240: _Opcode_name[709:714],
	241: _Opcode_name[714:724],
}

func (i Opcode) String() string {
	if str, ok := _Opcode_map[i]; ok {
		return str
	}
	return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
}
