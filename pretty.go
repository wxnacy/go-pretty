package pretty

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/mattn/go-runewidth"
)

type Field struct {
	Name       string // 输出标题
	Value      string //输出内容
	IsFillLeft bool   // 是否填充左边
}

type Pretty interface {
	BuildPretty() []Field
}

type enumerate interface {
	GetWriter() io.Writer
	List() []Pretty
}

type List []Pretty

func (l List) GetWriter() io.Writer {
	return os.Stdout
}

func (l List) List() []Pretty {
	return l
}

func (l *List) Add(pretty Pretty) *List {
	*l = append(*l, pretty)
	return l
}

func (l List) Print() {
	PrintList(l)
}

func PrintList(p enumerate) {
	items := p.List()
	if len(items) == 0 {
		return
	}
	maxLengthMap := getMaxLengthMap(items)
	prettyMap := getPrettyMap(items[0])
	firstLine := ""
	for _, pd := range items[0].BuildPretty() {
		name := pd.Name
		length, _ := maxLengthMap[pd.Name]
		p, _ := prettyMap[name]
		prettyName := fmtString(name, length, p.IsFillLeft)
		firstLine = firstLine + prettyName + " "

	}
	fmt.Fprintln(p.GetWriter(), firstLine)
	for _, pretty := range items {
		line := ""
		for _, pd := range pretty.BuildPretty() {
			length, _ := maxLengthMap[pd.Name]
			line = line + fmtString(pd.Value, length, pd.IsFillLeft) + " "
		}
		fmt.Fprintln(p.GetWriter(), line)
	}
	fmt.Fprintf(p.GetWriter(), "Total: %d\n", len(items))
}

func getPrettyMap(pretty Pretty) map[string]Field {
	m := make(map[string]Field, 0)
	for _, p := range pretty.BuildPretty() {
		m[p.Name] = p
	}
	return m
}

func getMaxLengthMap(items []Pretty) map[string]int {
	m := make(map[string]int, 0)
	for _, item := range items {
		data := item.BuildPretty()
		for _, d := range data {
			name := d.Name
			leng, exit := m[name]
			if exit {
				leng = int(math.Max(float64(runewidth.StringWidth(d.Value)), float64(leng)))
			} else {
				leng = runewidth.StringWidth(d.Value)
			}
			m[name] = leng
		}
	}
	for _, d := range items[0].BuildPretty() {
		name := d.Name
		leng, exit := m[name]
		if exit {
			leng = int(math.Max(float64(len(name)), float64(leng)))
		}
		m[name] = leng
	}
	return m
}

func fmtString(s string, length int, isFillLeft bool) string {
	var value string
	if isFillLeft {
		value = runewidth.FillLeft(s, length)
	} else {
		value = runewidth.FillRight(s, length)
	}
	return value
}
