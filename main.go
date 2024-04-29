package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var defaultCount = 10

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	count := defaultCount
	if len(os.Args) > 1 {
		countStr := os.Args[1]
		var err error
		count, err = strconv.Atoi(countStr)
		if err != nil {
			fmt.Println("Must provide a valid count number")
			os.Exit(1)
		}
	}

	characterMap := hiraganaToRomaji
	characterList := hiragana
	if len(os.Args) > 2 {
		val := os.Args[2]
		if val == "katakana" {
			characterMap = katakanaToRomaji
			characterList = katakana
		}
	}

	stats := map[string]*Performance{}
	fmt.Printf("Practicing %d kana\n", count)
	for i := 0; i < count; i++ {
		curr := randomItem(r, characterList)

		currStat, found := stats[curr]
		if !found {
			currStat = &Performance{
				Correct: 0,
				Wrong:   0,
			}
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s: ", curr)
		input, err := reader.ReadString('\n')
		if err != nil {
			writeStats(stats)
			os.Exit(1)
			return
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			writeStats(stats)
			os.Exit(0)
			return
		}
		ans, _ := characterMap[curr]
		if input != ans {
			currStat.Wrong++
			fmt.Printf("x (%s)\n", ans)
		} else {
			currStat.Correct++
			fmt.Printf("✓\n")
		}
		stats[curr] = currStat
	}
	writeStats(stats)
}

type Performance struct {
	Correct int
	Wrong   int
}

func writeStats(m map[string]*Performance) {
	totalCorrect := 0
	totalWrong := 0
	for k, v := range m {
		totalCorrect += v.Correct
		totalWrong += v.Wrong
		fmt.Printf("%s\t %d ✓ - %d x\n", k, v.Correct, v.Wrong)
	}
	score := (float32(totalCorrect) / float32(totalCorrect+totalWrong)) * 100
	fmt.Printf("Score: %.2f%%\n", score)
}

var (
	hiraganaToRomaji = map[string]string{
		"あ": "a", "い": "i", "う": "u", "え": "e", "お": "o",
		"か": "ka", "き": "ki", "く": "ku", "け": "ke", "こ": "ko",
		"さ": "sa", "し": "shi", "す": "su", "せ": "se", "そ": "so",
		"た": "ta", "ち": "chi", "つ": "tsu", "て": "te", "と": "to",
		"な": "na", "に": "ni", "ぬ": "nu", "ね": "ne", "の": "no",
		"は": "ha", "ひ": "hi", "ふ": "fu", "へ": "he", "ほ": "ho",
		"ま": "ma", "み": "mi", "む": "mu", "め": "me", "も": "mo",
		"や": "ya", "ゆ": "yu", "よ": "yo",
		"ら": "ra", "り": "ri", "る": "ru", "れ": "re", "ろ": "ro",
		"わ": "wa", "を": "wo", "ん": "n",
		"が": "ga", "ぎ": "gi", "ぐ": "gu", "げ": "ge", "ご": "go",
		"ざ": "za", "じ": "ji", "ず": "zu", "ぜ": "ze", "ぞ": "zo",
		"だ": "da", "ぢ": "ji", "づ": "zu", "で": "de", "ど": "do",
		"ば": "ba", "び": "bi", "ぶ": "bu", "べ": "be", "ぼ": "bo",
		"ぱ": "pa", "ぴ": "pi", "ぷ": "pu", "ぺ": "pe", "ぽ": "po",
		"きゃ": "kya", "きゅ": "kyu", "きょ": "kyo",
		"しゃ": "sha", "しゅ": "shu", "しょ": "sho",
		"ちゃ": "cha", "ちゅ": "chu", "ちょ": "cho",
		"にゃ": "nya", "にゅ": "nyu", "にょ": "nyo",
		"ひゃ": "hya", "ひゅ": "hyu", "ひょ": "hyo",
		"みゃ": "mya", "みゅ": "myu", "みょ": "myo",
		"りゃ": "rya", "りゅ": "ryu", "りょ": "ryo",
		"ぎゃ": "gya", "ぎゅ": "gyu", "ぎょ": "gyo",
		"じゃ": "ja", "じゅ": "ju", "じょ": "jo",
		"びゃ": "bya", "びゅ": "byu", "びょ": "byo",
		"ぴゃ": "pya", "ぴゅ": "pyu", "ぴょ": "pyo",
	}
	hiragana = keysToSlice(hiraganaToRomaji)

	katakanaToRomaji = map[string]string{
		"ア": "a", "イ": "i", "ウ": "u", "エ": "e", "オ": "o",
		"カ": "ka", "キ": "ki", "ク": "ku", "ケ": "ke", "コ": "ko",
		"サ": "sa", "シ": "shi", "ス": "su", "セ": "se", "ソ": "so",
		"タ": "ta", "チ": "chi", "ツ": "tsu", "テ": "te", "ト": "to",
		"ナ": "na", "ニ": "ni", "ヌ": "nu", "ネ": "ne", "ノ": "no",
		"ハ": "ha", "ヒ": "hi", "フ": "fu", "ヘ": "he", "ホ": "ho",
		"マ": "ma", "ミ": "mi", "ム": "mu", "メ": "me", "モ": "mo",
		"ヤ": "ya", "ユ": "yu", "ヨ": "yo",
		"ラ": "ra", "リ": "ri", "ル": "ru", "レ": "re", "ロ": "ro",
		"ワ": "wa", "ヲ": "wo", "ン": "n",
		"ガ": "ga", "ギ": "gi", "グ": "gu", "ゲ": "ge", "ゴ": "go",
		"ザ": "za", "ジ": "ji", "ズ": "zu", "ゼ": "ze", "ゾ": "zo",
		"ダ": "da", "ヂ": "ji", "ヅ": "zu", "デ": "de", "ド": "do",
		"バ": "ba", "ビ": "bi", "ブ": "bu", "ベ": "be", "ボ": "bo",
		"パ": "pa", "ピ": "pi", "プ": "pu", "ペ": "pe", "ポ": "po",
		"キャ": "kya", "キュ": "kyu", "キョ": "kyo",
		"シャ": "sha", "シュ": "shu", "ショ": "sho",
		"チャ": "cha", "チュ": "chu", "チョ": "cho",
		"ニャ": "nya", "ニュ": "nyu", "ニョ": "nyo",
		"ヒャ": "hya", "ヒュ": "hyu", "ヒョ": "hyo",
		"ミャ": "mya", "ミュ": "myu", "ミョ": "myo",
		"リャ": "rya", "リュ": "ryu", "リョ": "ryo",
		"ギャ": "gya", "ギュ": "gyu", "ギョ": "gyo",
		"ジャ": "ja", "ジュ": "ju", "ジョ": "jo",
		"ビャ": "bya", "ビュ": "byu", "ビョ": "byo",
		"ピャ": "pya", "ピュ": "pyu", "ピョ": "pyo",
	}
	katakana = keysToSlice(katakanaToRomaji)
)

func keysToSlice(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys

}

func randomItem(r *rand.Rand, list []string) string {
	randomIndex := r.Intn(len(list))
	return list[randomIndex]
}
