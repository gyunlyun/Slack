package runner

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"slack-wails/core/webscan/proto"
	"slack-wails/lib/util"

	"strconv"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter/functions"
)

var (
	NewProgramOptions = []cel.ProgramOption{
		cel.Functions(
			// string
			&functions.Overload{
				Operator: "string_icontains_string",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					v1, ok := lhs.(types.String)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to contains", lhs.Type())
					}
					v2, ok := rhs.(types.String)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to contains", rhs.Type())
					}
					return types.Bool(strings.Contains(strings.ToLower(string(v1)), strings.ToLower(string(v2))))
				},
			},
			&functions.Overload{
				Operator: "substr_string_int_int",
				Function: func(values ...ref.Val) ref.Val {
					if len(values) == 3 {
						str, ok := values[0].(types.String)
						if !ok {
							return types.NewErr("invalid string to 'substr'")
						}
						start, ok := values[1].(types.Int)
						if !ok {
							return types.NewErr("invalid start to 'substr'")
						}
						length, ok := values[2].(types.Int)
						if !ok {
							return types.NewErr("invalid length to 'substr'")
						}
						runes := []rune(str)
						if start < 0 || length < 0 || int(start+length) > len(runes) {
							return types.NewErr("invalid start or length to 'substr'")
						}
						return types.String(runes[start : start+length])
					} else {
						return types.NewErr("too many arguments to 'substr'")
					}
				},
			},
			&functions.Overload{
				Operator: "replaceAll_string_string_string",
				Function: func(values ...ref.Val) ref.Val {
					s, ok := values[0].(types.String)
					if !ok {
						return types.ValOrErr(s, "unexpected type '%v' passed to replaceAll", s.Type())
					}
					old, ok := values[1].(types.String)
					if !ok {
						return types.ValOrErr(old, "unexpected type '%v' passed to replaceAll", old.Type())
					}
					new, ok := values[2].(types.String)
					if !ok {
						return types.ValOrErr(new, "unexpected type '%v' passed to replaceAll", new.Type())
					}

					return types.String(strings.ReplaceAll(string(s), string(old), string(new)))
				},
			},
			&functions.Overload{
				Operator: "toUpper_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to toUpper_string", value.Type())
					}

					return types.String(strings.ToUpper(string(v)))
				},
			},
			&functions.Overload{
				Operator: "toLower_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to toLower_string", value.Type())
					}

					return types.String(strings.ToLower(string(v)))
				},
			},

			// []byte
			&functions.Overload{
				Operator: "bytes_bcontains_bytes",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					v1, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to bcontains", lhs.Type())
					}
					v2, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to bcontains", rhs.Type())
					}
					return types.Bool(bytes.Contains(v1, v2))
				},
			},
			&functions.Overload{
				Operator: "bytes_ibcontains_bytes",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					v1, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to bcontains", lhs.Type())
					}
					v2, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to bcontains", rhs.Type())
					}
					return types.Bool(bytes.Contains(bytes.ToLower(v1), bytes.ToLower(v2)))
				},
			},
			&functions.Overload{
				Operator: "bytes_bstartsWith_bytes",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					v1, ok := lhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to bstartsWith", lhs.Type())
					}
					v2, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to bstartsWith", rhs.Type())
					}
					return types.Bool(bytes.HasPrefix(v1, v2))
				},
			},
			// encode
			&functions.Overload{
				Operator: "md5_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to md5_string", value.Type())
					}
					return types.String(fmt.Sprintf("%x", md5.Sum([]byte(v))))
				},
			},
			&functions.Overload{
				Operator: "base64_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to base64_string", value.Type())
					}
					return types.String(base64.StdEncoding.EncodeToString([]byte(v)))
				},
			},
			&functions.Overload{
				Operator: "base64_bytes",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.Bytes)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to base64_bytes", value.Type())
					}
					return types.String(base64.StdEncoding.EncodeToString(v))
				},
			},
			&functions.Overload{
				Operator: "base64Decode_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to base64Decode_string", value.Type())
					}
					decodeBytes, err := base64.StdEncoding.DecodeString(string(v))
					if err != nil {
						return types.NewErr("%v", err)
					}
					return types.String(decodeBytes)
				},
			},
			&functions.Overload{
				Operator: "base64Decode_bytes",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.Bytes)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to base64Decode_bytes", value.Type())
					}
					decodeBytes, err := base64.StdEncoding.DecodeString(string(v))
					if err != nil {
						return types.NewErr("%v", err)
					}
					return types.String(decodeBytes)
				},
			},
			&functions.Overload{
				Operator: "urlencode_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to urlencode_string", value.Type())
					}
					return types.String(url.QueryEscape(string(v)))
				},
			},
			&functions.Overload{
				Operator: "urlencode_bytes",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.Bytes)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to urlencode_bytes", value.Type())
					}
					return types.String(url.QueryEscape(string(v)))
				},
			},
			&functions.Overload{
				Operator: "urldecode_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to urldecode_string", value.Type())
					}
					decodeString, err := url.QueryUnescape(string(v))
					if err != nil {
						return types.NewErr("%v", err)
					}
					return types.String(decodeString)
				},
			},
			&functions.Overload{
				Operator: "urldecode_bytes",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.Bytes)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to urldecode_bytes", value.Type())
					}
					decodeString, err := url.QueryUnescape(string(v))
					if err != nil {
						return types.NewErr("%v", err)
					}
					return types.String(decodeString)
				},
			},
			&functions.Overload{
				Operator: "hexdecode_string",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.String)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to hexdecode_string", value.Type())
					}
					dst := make([]byte, hex.DecodedLen(len(v)))
					n, err := hex.Decode(dst, []byte(v))
					if err != nil {
						return types.ValOrErr(value, "unexpected type '%s' passed to hexdecode_string", err.Error())
					}
					return types.String(string(dst[:n]))
				},
			},
			// random
			&functions.Overload{
				Operator: "randomInt_int_int",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					from, ok := lhs.(types.Int)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to randomInt", lhs.Type())
					}
					to, ok := rhs.(types.Int)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to randomInt", rhs.Type())
					}
					min, max := int(from), int(to)
					return types.Int(rand.Intn(max-min) + min)
				},
			},
			&functions.Overload{
				Operator: "randomLowercase_int",
				Unary: func(value ref.Val) ref.Val {
					n, ok := value.(types.Int)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to randomLowercase", value.Type())
					}
					return types.String(util.RandLetters(int(n)))
				},
			},
			// regex
			&functions.Overload{
				Operator: "string_bmatches_bytes",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					var isMatch = false
					var err error

					v1, ok := lhs.(types.String)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to bmatches", lhs.Type())
					}
					v2, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to bmatches", rhs.Type())
					}
					re := regexp2.MustCompile(string(v1), 0)
					if isMatch, err = re.MatchString(string([]byte(v2))); err != nil {
						return types.NewErr("%v", err)
					}
					return types.Bool(isMatch)
				},
			},
			// reverse
			&functions.Overload{
				Operator: "reverse_wait_int",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					reverse, ok := lhs.Value().(*proto.Reverse)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to 'wait'", lhs.Type())
					}
					timeout, ok := rhs.Value().(int64)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to 'wait'", rhs.Type())
					}
					return types.Bool(reverseCheck(reverse, timeout))
				},
			},
			// other
			&functions.Overload{
				Operator: "sleep_int",
				Unary: func(value ref.Val) ref.Val {
					v, ok := value.(types.Int)
					if !ok {
						return types.ValOrErr(value, "unexpected type '%v' passed to sleep", value.Type())
					}
					time.Sleep(time.Duration(v) * time.Second)
					return nil
				},
			},
			// year
			&functions.Overload{
				Operator: "year_string",
				Unary: func(value ref.Val) ref.Val {
					year := time.Now().Format("2006")
					return types.String(year)
				},
			},
			&functions.Overload{
				Operator: "shortyear_string",
				Unary: func(value ref.Val) ref.Val {
					year := time.Now().Format("06")
					return types.String(year)
				},
			},
			&functions.Overload{
				Operator: "month_string",
				Unary: func(value ref.Val) ref.Val {
					month := time.Now().Format("01")
					return types.String(month)
				},
			},
			&functions.Overload{
				Operator: "day_string",
				Unary: func(value ref.Val) ref.Val {
					day := time.Now().Format("02")
					return types.String(day)
				},
			},
			&functions.Overload{
				Operator: "timestamp_second_string",
				Unary: func(value ref.Val) ref.Val {
					timestamp := strconv.FormatInt(time.Now().Unix(), 10)
					return types.String(timestamp)
				},
			},
		),
	}
)

func ReadProgramOptions(reg ref.TypeRegistry) []cel.ProgramOption {
	allProgramOpitons := []cel.ProgramOption{
		cel.Functions(
			&functions.Overload{
				Operator: "string_submatch_string",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					var (
						resultMap = make(map[string]string)
					)

					v1, ok := lhs.(types.String)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to submatch", lhs.Type())
					}
					v2, ok := rhs.(types.String)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to submatch", rhs.Type())
					}

					re := regexp2.MustCompile(string(v1), regexp2.RE2)
					if m, _ := re.FindStringMatch(string(v2)); m != nil {
						gps := m.Groups()
						for n, gp := range gps {
							if n == 0 {
								continue
							}
							resultMap[gp.Name] = gp.String()
						}
					}
					return types.NewStringStringMap(reg, resultMap)
				},
			},
			&functions.Overload{
				Operator: "string_bsubmatch_bytes",
				Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
					var (
						resultMap = make(map[string]string)
					)

					v1, ok := lhs.(types.String)
					if !ok {
						return types.ValOrErr(lhs, "unexpected type '%v' passed to bsubmatch", lhs.Type())
					}
					v2, ok := rhs.(types.Bytes)
					if !ok {
						return types.ValOrErr(rhs, "unexpected type '%v' passed to bsubmatch", rhs.Type())
					}

					re := regexp2.MustCompile(string(v1), regexp2.RE2)
					if m, _ := re.FindStringMatch(string([]byte(v2))); m != nil {
						gps := m.Groups()
						for n, gp := range gps {
							if n == 0 {
								continue
							}
							resultMap[gp.Name] = gp.String()
						}
					}

					return types.NewStringStringMap(reg, resultMap)
				},
			},
		),
	}
	allProgramOpitons = append(allProgramOpitons, NewProgramOptions...)
	return allProgramOpitons
}
