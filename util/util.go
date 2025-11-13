package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"exclusive_base_qz/kitex_gen/base"
	"fmt"
	"html"
	"math"
	"math/rand"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ProcessBaseResp(StatusCode int, StatusMessage string) *base.BaseResp {
	BaseResp := &base.BaseResp{
		StatusCode:    int32(StatusCode),
		StatusMessage: StatusMessage,
	}
	return BaseResp
}

func ArrayToString(slice []interface{}, separator string) string {
	array := TransInterfaceToSlice(slice)
	res := ""
	for i, val := range array {
		if i != 0 {
			res += separator
		}
		res += ToString(val)
	}
	return res
}

func ArrayStrToString(strArray []*string, separator string) string {
	res := ""
	for i, val := range strArray {
		if i != 0 {
			res += separator
		}
		res += *val
	}
	return res
}

func BoolToInt(b bool) int8 {
	if b {
		return int8(1)
	}
	return int8(0)
}

func CopyInterface(src, dst interface{}) error {
	dataByte, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = Unmarshal(string(dataByte), &dst)
	if err != nil {
		return err
	}
	return nil
}

func EqualsTo(a interface{}) func(interface{}) bool {
	return func(b interface{}) bool {
		return a == b
	}
}

func NotEqualsTo(a interface{}) func(interface{}) bool {
	return func(b interface{}) bool {
		return a != b
	}
}

func FilterDupInt64ByArrayB(arrA []int64, arrB []int64) []int64 {
	if len(arrA) == 0 || len(arrB) == 0 {
		return arrA
	}

	res := []int64{}
	sort.Slice(arrA, func(i, j int) bool { return arrA[i] < arrA[j] })
	sort.Slice(arrB, func(i, j int) bool { return arrB[i] < arrB[j] })
	j := 0
	for i := 0; i < len(arrA); i++ {
		for ; j < len(arrB); j++ {
			if arrB[j] >= arrA[i] {
				break
			}
		}
		if j == len(arrB) || arrB[j] > arrA[i] {
			res = append(res, arrA[i])
		}
	}

	return res
}

func FilterDuplicates(elems interface{}) []interface{} {
	elemIntrs := TransInterfaceToSlice(elems)
	encountered := make(map[interface{}]bool)

	j := 0
	for i := 0; i < len(elemIntrs); i++ {
		if _, found := encountered[elemIntrs[i]]; !found {
			if i != j {
				elemIntrs[j] = elemIntrs[i]
			}
			j++
			encountered[elemIntrs[i]] = true
		}
	}
	return elemIntrs[:j]
}

func FilterInt32(a []int32, lambda func(interface{}) bool) (ret []int32) {
	for _, ele := range a {
		if lambda(ele) {
			ret = append(ret, ele)
		}
	}
	return
}

func FilterInt64(a []int64, lambda func(interface{}) bool) (ret []int64) {
	for _, ele := range a {
		if lambda(ele) {
			ret = append(ret, ele)
		}
	}
	return
}

func FilterString(a []string, lambda func(interface{}) bool) (ret []string) {
	for _, ele := range a {
		if lambda(ele) {
			ret = append(ret, ele)
		}
	}
	return
}

func StringEquals(a string, b string) bool {
	return strings.Compare(a, b) == 0
}

func GetFromNil(i interface{}, t reflect.Type) reflect.Value {
	if i != nil {
		return reflect.ValueOf(i)
	}
	return reflect.Zero(t)
}

func GenMD5(s string, salt string) string {
	if salt != "" {
		s = GetStrFromTpl("%v_%v", s, salt)
	}
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetStrFromTpl(tpl string, args ...interface{}) string {
	return fmt.Sprintf(tpl, args...)
}

func InterfaceConv(i interface{}, defaultValue interface{}, valueType string) interface{} {
	stringValue := ToString(i)
	return StringConv(stringValue, defaultValue, valueType)
}

func MapInt32Keys(imap interface{}) []int32 {
	values := reflect.ValueOf(imap).MapKeys()
	keys := make([]int32, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.Interface().(int32))
	}
	return keys
}

func MapInt64Keys(imap interface{}) []int64 {
	values := reflect.ValueOf(imap).MapKeys()
	keys := make([]int64, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.Interface().(int64))
	}
	return keys
}

func MapStringKeys(imap interface{}) []string {
	values := reflect.ValueOf(imap).MapKeys()
	keys := make([]string, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.Interface().(string))
	}
	return keys
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinInt16(x, y int16) int16 {
	if x < y {
		return x
	}
	return y
}

func MaxInt16(x, y int16) int16 {
	if x > y {
		return x
	}
	return y
}

func MinInt32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func MaxInt32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MergeJson(src, dst string) (string, error) {
	if len(src) < 3 {
		return dst, nil
	}

	from, to := make(map[string]interface{}), make(map[string]interface{})
	Unmarshal(src, &from)
	Unmarshal(dst, &to)
	for k, v := range from {
		to[k] = v
	}
	data, err := json.Marshal(&to)
	return string(data), err
}

func MergeJsonStrFromMap(srcStr string, data map[string]interface{}) (string, error) {
	src := make(map[string]interface{})
	Unmarshal(srcStr, &src)

	for k, v := range data {
		src[k] = v
	}

	rawData, err := json.Marshal(&src)
	return string(rawData), err
}

func StringConv(s string, defaultValue interface{}, valueType string) interface{} {
	var err error
	var value interface{}
	switch valueType {
	case "bool":
		value, err = strconv.ParseBool(s)
	case "int":
		value, err = strconv.ParseInt(s, 10, 0)
		value = int(value.(int64))
	case "int8":
		value, err = strconv.ParseInt(s, 10, 8)
		value = int8(value.(int64))
	case "int16":
		value, err = strconv.ParseInt(s, 10, 16)
		value = int16(value.(int64))
	case "int32":
		value, err = strconv.ParseInt(s, 10, 32)
		value = int32(value.(int64))
	case "int64":
		value, err = strconv.ParseInt(s, 10, 64)
	case "float64":
		value, err = strconv.ParseFloat(s, 64)
	case "string":
		value = s
		if value == "" {
			value = defaultValue.(string)
		}
	case "[]string":
		value = strings.Fields(s)
	default:
		panic(fmt.Sprintf("nonexistent value type: %s", valueType))
	}
	if err != nil {
		value = defaultValue
	}
	return value
}

func StringPtrIfNotEmpty(input string) *string {
	if input == "" {
		return nil
	}

	return &input
}

func StringToArray(str string, separator string) []interface{} {
	array := strings.Split(str, separator)
	res := make([]interface{}, len(array))
	for i, val := range array {
		res[i] = val
	}
	return res
}

func StringToMap(str string) map[string]interface{} {
	m := map[string]interface{}{}
	if err := Unmarshal(str, &m); err != nil {
		return nil
	}
	return m
}

func ToError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

func ToJsonString(i interface{}) string {
	if i == nil {
		return ""
	}

	data, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(data)
}

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprint(i)
}

func TransInterfaceToSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("TransInterfaceToSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func UnescapedString(unescaped string) string {
	return html.UnescapeString(unescaped)
}

func Unmarshal(data string, v interface{}) error {
	d := json.NewDecoder(strings.NewReader(data))
	d.UseNumber()
	return d.Decode(v)
}

func UrlEncoded(str string) string {
	u, err := url.Parse(str)
	if err != nil {
		return ""
	}
	return u.String()
}

func SliceInt8Contains(slice []int8, element int8) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceInt16Contains(slice []int16, element int16) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceInt32Contains(slice []int32, element int32) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceInt64Contains(slice []int64, element int64) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func SliceIntrContains(slice interface{}, ele interface{}) bool {
	for _, e := range TransInterfaceToSlice(slice) {
		if e == ele {
			return true
		}
	}
	return false
}

func SliceStrContains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func StringIndexBy(i int, s []string) string {
	if i < 0 || i >= len(s) {
		return ""
	}
	return s[i]
}

// Returns an int >= min, < max
func RandomIntRange(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func BoolToIntStr(v bool) string {
	if v {
		return "1"
	} else {
		return "0"
	}
}

func If(cond bool, t interface{}, f interface{}) interface{} {
	if cond {
		return t
	} else {
		return f
	}
}

// if typeOf(fn) == func, run it
func callFn(fn interface{}) interface{} {
	if fn != nil {
		tp := reflect.TypeOf(fn)
		if tp.Kind() == reflect.Func && tp.NumIn() == 0 {
			function := reflect.ValueOf(fn)
			in := make([]reflect.Value, 0)
			out := function.Call(in)
			if num := len(out); num > 0 {
				list := make([]interface{}, num)
				for i, value := range out {
					list[i] = value.Interface()
				}
				if num == 1 {
					return list[0]
				}
				return list
			}
			return nil
		}
	}
	return fn
}

// example
//
//	a := 0
//	b := 4
//	c := IF(a == 0, 0, func() int { return b / a })
func IF(cond bool, t interface{}, f interface{}) interface{} {
	if cond {
		return callFn(t)
	} else {
		return callFn(f)
	}
}

//func Retry(count int, sleep int, f func() (success bool)) bool {
//	for retry := 0; retry < count; retry++ {
//		success := f()
//		if success {
//			return true
//		} else {
//			left := count - retry - 1
//			if left == 0 {
//				return false
//			} else {
//				time.Sleep(time.Duration(sleep) * time.Second)
//			}
//		}
//	}
//	return false
//}

// 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

func ToJson(obj interface{}) string {
	byt, err := json.Marshal(obj)
	if err != nil {

	}
	return string(byt)
}

func StrArray2Int64Array(strArray []string) ([]int64, error) {
	result := make([]int64, 0)
	for _, str := range strArray {
		var data int64
		var err error
		if data, err = strconv.ParseInt(str, 10, 64); err != nil {
			return nil, err
		}
		result = append(result, data)
	}
	return result, nil
}

func GetNoExistsId(existsMap map[int64]bool, totalIds []int64) (noExists []int64) {
	noExists = make([]int64, 0)
	for _, id := range totalIds {
		if _, ok := existsMap[id]; !ok {
			noExists = append(noExists, id)
		}
	}
	return noExists
}

func Int64Array2ToStrArray(intList []int64) []string {
	result := make([]string, len(intList))
	for i, v := range intList {
		result[i] = fmt.Sprintf("%d", v)
	}
	return result
}

func CheckSameIDs(ids []string) bool {
	if ids == nil || len(ids) <= 0 {
		return false
	}
	idMap := make(map[string]bool)
	for _, id := range ids {
		_, ok := idMap[id]
		if ok {
			return true
		}
		idMap[id] = true
	}
	return false
}

func StringToStringArrayWithComma(data string) []string {
	if len(data) == 0 {
		return make([]string, 0)
	}
	return strings.Split(data, ",")
}

func StringArray2ToInt64Array(stringList []string) ([]int64, error) {
	result := make([]int64, len(stringList))
	var err error
	for i, v := range stringList {
		result[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func StrToInt64(str string) int64 {
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func Int64ToFloat64(i int64) float64 {
	i64str := strconv.FormatInt(i, 10)
	fData, _ := strconv.ParseFloat(i64str, 64)
	return fData
}

// IsNum 是否是数字
func IsNum(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func ArrayIsContainString(items []interface{}, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
