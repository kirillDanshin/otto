package regexp

import (
	"sync"

	"github.com/kirillDanshin/go-pcre"
)

type regexpCache struct {
	cache map[string]*pcre.Regexp

	sync.RWMutex
}

var regs *regexpCache

func init() {
	regs = &regexpCache{
		cache: make(map[string]*pcre.Regexp),
	}
	presets := []string{
		"",
		".*",
		"(.*)",
		".+",
		"(.+)",
		`'`,
		`(')`,
		`"`,
		`(")`,
		`\+`,
		`(\+)`,
		`%20`,
		`(%20)`,
		`^https?:\/\/[^\/]*`,
		`[A-Z]`,
		`([A-Z])`,
		`^ms-`,
	}

	regs.Lock()
	defer regs.Unlock()
	for _, v := range presets {
		if r, err := pcre.CompileParse(v); err == nil {
			regs.cache[v] = &r
		}
	}
}
