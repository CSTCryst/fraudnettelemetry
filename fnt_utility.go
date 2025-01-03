package fraudnettelemetry

import (
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/CSTCryst/fraudnettelemetry/internal"
	"github.com/CSTCryst/fraudnettelemetry/internal/device/device_screen"
	"github.com/CSTCryst/fraudnettelemetry/internal/fnt"
	"github.com/colduction/nocopy"
	"github.com/colduction/randomizer"
)

const (
	upperCaseAlphaNumeric string = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

	/*
		Golden ratio

		info: https://en.wikipedia.org/wiki/Golden_ratio

		ref: https://softwareengineering.stackexchange.com/a/63605
	*/
	goldenRatio int64 = 1640531527
	numRounds   int   = 32
)

var specialUnicodeValues map[rune]string = map[rune]string{0: "", 8: "", 9: "", 13: "", 16: "", 17: "", 18: "", 37: "", 38: "", 39: "", 40: "", 46: "", 91: "", 93: "", 219: "", 220: "", 224: ""}

func (m *fntBuilder) generateDC(userAgent string) (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	if userAgent == "" {
		return nil, internal.EmptyInputError("generateDC")
	}
	screen := device_screen.NewScreenBuilder()
	defer screen.Reset(true)
	screen.SetAppleSmartphone()
	dc := fnt.NewDCBuilder(userAgent)
	defer dc.Reset(true)
	dc.UA = userAgent
	dc.Screen.ColorDepth = screen.ColorDepth
	dc.Screen.PixelDepth = screen.ColorDepth
	dc.Screen.Height = screen.LogicalResolution[1]
	dc.Screen.Width = screen.LogicalResolution[0]
	dc.Screen.AvailHeight = screen.LogicalResolution[1] - randomizer.UintInterval[uint32](0, 250)
	dc.Screen.AvailWidth = screen.LogicalResolution[0] - randomizer.UintInterval[uint32](0, 250)
	data, err := dc.String()
	if err != nil {
		return nil, err
	}
	m.dc = data
	return m, nil
}

func (m *fntBuilder) generateChk() (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	if m.f == "" {
		return nil, internal.EmptyInputError("generateChk")
	}
	ctime := time.Now().UnixMilli()
	m.chk.TS = ctime - randomizer.IntInterval[int64](9000, 20000)
	var (
		output         [8]interface{}
		c              [4]int64
		u              [2]int64
		fnsiRuneArr    []rune = []rune(m.f)
		startTime      string = strconv.FormatInt(m.chk.TS, 10)
		fnsiRuneArrLen int    = len(fnsiRuneArr)
		i              int
		j              int
	)
	for i = 0; i < 4; i++ {
		c[i] = f(internal.SliceString(startTime, 4*i, 4*(i+1)))
	}
	for i = 0; i < fnsiRuneArrLen; i += 8 {
		u[0] = int64(((255 & internal.IndexNum(fnsiRuneArr, i)) << 24) | ((255 & internal.IndexNum(fnsiRuneArr, i+1)) << 16) | ((255 & internal.IndexNum(fnsiRuneArr, i+2)) << 8) | (255 & internal.IndexNum(fnsiRuneArr, i+3)))
		u[1] = int64(((255 & internal.IndexNum(fnsiRuneArr, i+4)) << 24) | ((255 & internal.IndexNum(fnsiRuneArr, i+5)) << 16) | ((255 & internal.IndexNum(fnsiRuneArr, i+6)) << 8) | (255 & internal.IndexNum(fnsiRuneArr, i+7)))
		xteaEncrypt(&u, &c)
		output[j] = u[0]
		output[j+1] = u[1]
		j += 2
	}
	m.chk.ETEID = output[:]
	m.chk.TTS = ctime - m.chk.TS
	return m, nil
}

// Generates "ts" values and both simulated keyboard pressed keys and simulated mouse movements values with some changes
func (m *fntBuilder) generateD(inputArr []string) (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	inputArrLen := len(inputArr)
	if inputArr == nil || inputArrLen <= 0 || inputArrLen > 5 {
		return nil, internal.EmptyOrInvalidInputError("generateD")
	}
	var (
		tmp         [][2]string     = make([][2]string, inputArrLen)
		ts, tsIndex strings.Builder = strings.Builder{}, strings.Builder{}
	)
	tsIndex.Grow(2 + internal.CountTotalDigits(inputArrLen))
	for i := 0; i < inputArrLen; i++ {
		for j, inputArrIndexedLen := 0, len(inputArr[i]); j < inputArrIndexedLen; j++ {
			if v, err := simulateKeyPress(inputArr[i][j], j); err != nil {
				return nil, err
			} else {
				ts.WriteString(v)
				tsIndex.WriteString("ts")
				tsIndex.WriteString(strconv.Itoa(i + 1))
				tmp[i][0] = tsIndex.String()
				tmp[i][1] = ts.String()
			}
			tsIndex.Reset()
		}
		ts.Reset()
	}
	m.d = make(map[string]string, inputArrLen)
	for i, tmpLen := 0, len(tmp); i < tmpLen; i++ {
		ts.WriteString(tmp[i][1])
		ts.WriteString("Uh:")
		ts.WriteString(strconv.FormatInt(internal.SumUnicodeValue[int64](tmp[i][1]), 10))
		m.d[tmp[i][0]] = ts.String()
		ts.Reset()
	}
	ts.Reset()
	tsIndex.Reset()
	return m, nil
}

func GenerateTLTSID() string {
	var (
		sb     strings.Builder = strings.Builder{}
		output string          = ""
	)
	sb.Grow(32)
	sb.WriteString(generateRandomString(1, "123456789"))
	sb.WriteString(generateRandomString(31, "1234567890"))
	output = sb.String()
	sb.Reset()
	return output
}

func simulateKeyPress(b byte, index int) (string, error) {
	if b == 0 {
		return "", internal.EmptyOrInvalidInputError("simulateKeyPress")
	}
	var (
		timeDiff int64           = 0
		output   string          = ""
		sb       strings.Builder = strings.Builder{}
	)
	if index == 0 {
		timeDiff = randomizer.IntInterval[int64](2900, 4900)
	} else {
		timeDiff = randomizer.IntInterval[int64](65, 125)
	}
	keyCode, _ := utf8.DecodeRune(nocopy.ByteToByteSlice(b))
	timeDiffStr := strconv.FormatInt(timeDiff, 10)
	if _, ok := specialUnicodeValues[keyCode]; ok {
		sb.WriteString("Dk")
		sb.WriteString(strconv.FormatInt(int64(keyCode), 10))
		sb.WriteString(":")
		sb.WriteString(timeDiffStr)
		sb.WriteString("Uk")
		sb.WriteString(strconv.FormatInt(int64(keyCode), 10))
		sb.WriteString(":")
		if index == 0 {
			sb.WriteString(strconv.FormatInt(randomizer.IntInterval[int64](80, 150), 10))
		} else {
			sb.WriteString(strconv.FormatInt(randomizer.IntInterval[int64](35, 900), 10))
		}
		output = sb.String()
		sb.Reset()
		return output, nil
	}
	indexStr := strconv.FormatInt(int64(index), 10)
	sb.WriteString("Di")
	sb.WriteString(indexStr)
	sb.WriteString(":")
	sb.WriteString(timeDiffStr)
	sb.WriteString("Ui")
	sb.WriteString(indexStr)
	sb.WriteString(":")
	if index == 0 {
		sb.WriteString(strconv.FormatInt(randomizer.IntInterval[int64](80, 150), 10))
	} else {
		sb.WriteString(strconv.FormatInt(randomizer.IntInterval[int64](35, 900), 10))
	}
	output = sb.String()
	sb.Reset()
	return output, nil
}

// XTea Encryption - JavaScript Implementation: https://stackoverflow.com/questions/76060523/converting-customized-xtea-algorithm-from-javascript-to-golang
func xteaEncrypt(v *[2]int64, k *[4]int64) {
	if v == nil || k == nil {
		return
	}
	for i, sum, temp := 0, int64(0), int32(0); i < numRounds; i++ {
		temp = int32(v[1])
		v[0] += int64((((temp << 4) ^ (temp >> 5)) + temp) ^ int32(sum+k[int32(sum)&3]))
		sum -= goldenRatio
		temp = int32(v[0])
		v[1] += int64((((temp << 4) ^ (temp >> 5)) + temp) ^ int32(sum+k[(int32(sum)>>11)&3]))
	}
}

func f(s string) (t int64) {
	if s != "" {
		for i, sLen := 0, len(s); i < sLen; i++ {
			t |= int64(int32(s[i]) << (8 * i)) // Bitwise inclusive OR operation
		}
	}
	return t
}

func generateRandomString(length int, sMap string) string {
	if length <= 0 {
		return ""
	}
	if sMap == "" {
		sMap = upperCaseAlphaNumeric
	}
	var (
		output string          = ""
		sb     strings.Builder = strings.Builder{}
	)
	sb.Grow(length)
	for i, r := 0, len(sMap); i < length; i++ {
		sb.WriteByte(sMap[randomizer.IntInterval(0, r)])
	}
	output = sb.String()
	sb.Reset()
	return output
}
