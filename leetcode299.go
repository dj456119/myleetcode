/*
 * @Descripttion:你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：

写出一个秘密数字，并请朋友猜这个数字是多少。朋友每猜测一次，你就会给他一个包含下述信息的提示：

猜测数字中有多少位属于数字和确切位置都猜对了（称为 "Bulls", 公牛），
有多少位属于数字猜对了但是位置不对（称为 "Cows", 奶牛）。也就是说，这次猜测中有多少位非公牛数字可以通过重新排列转换成公牛数字。
给你一个秘密数字 secret 和朋友猜测的数字 guess ，请你返回对朋友这次猜测的提示。

提示的格式为 "xAyB" ，x 是公牛个数， y 是奶牛个数，A 表示公牛，B 表示奶牛。

请注意秘密数字和朋友猜测的数字都可能含有重复数字。
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-08 00:26:57
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-08 00:26:58
*/

package myleetcode

import "fmt"

func getHint(secret string, guess string) string {
	Acount := 0
	Bcount := 0
	serMap := make(map[byte]int)
	guesMap := make(map[byte]int)

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			Acount++
		} else {
			if _, ok := serMap[secret[i]]; ok {
				serMap[secret[i]]++
			} else {
				serMap[secret[i]]++
			}
			if _, ok := guesMap[guess[i]]; ok {
				guesMap[guess[i]]++
			} else {
				guesMap[guess[i]]++
			}
		}
	}
	for k, v := range guesMap {
		if v1, ok := serMap[k]; ok {
			if v >= v1 {
				Bcount = Bcount + v1
			} else {
				Bcount = Bcount + v
			}
		}
	}
	return fmt.Sprintf("%d%c%d%c", Acount, 'A', Bcount, 'B')
}

func GetHint(secret string, guess string) string {
	return getHint(secret, guess)
}
