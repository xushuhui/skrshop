package utils

import (
	cr "crypto/rand"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

/*
字符串截取函数

参数：
	str:带截取字符串
	begin:开始截取位置
	length:截取长度
*/
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

// 生成一个随机int64

func MinInt64(a, b int64) (r int64) {
	if a > b {
		r = b
	} else {
		r = a
	}
	return
}

func MaxInt64(a, b int64) (r int64) {
	if a > b {
		r = a
	} else {
		r = b
	}
	return
}

const (
	snum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func FilterEmoji(content string) string {
	added := false
	newContent := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		} else {
			if !added {
				newContent += "{emoji表情}"
				added = true
			}
		}
	}
	return newContent
}

func SubAtNameString(str string) (nameList []string) {
	if ok := strings.Contains(str, "@all "); ok {
		nameList = append(nameList, "all")
		return
	}
	atIdx := strings.Index(str, "@")
	atStrs := strings.Split(str[atIdx+1:], "@")
	for _, v := range atStrs {
		placeIdx := strings.Index(v, " ")
		if placeIdx == -1 {
			continue
		}
		nameList = append(nameList, v[:placeIdx])
	}
	return
}

func IntFenToYuanStr(fen int64) (yuan string) {
	yuan = strconv.FormatFloat(float64(fen)/100, 'f', -1, 64)
	return
}

func Int01FenToYuanStr(fen int64) (yuan string) {
	yuan = strconv.FormatFloat(float64(fen)/10000, 'f', -1, 64)
	return
}
func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(cr.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
