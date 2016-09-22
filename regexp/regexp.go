package regexp

import (
	"log"

	"github.com/kirillDanshin/go-pcre"
	"github.com/kirillDanshin/myutils"
)

const delim = `/`

type Regexp struct {
	re *pcre.Regexp
}

type Matcher struct {
	*pcre.Matcher
}

func MustCompile(pattern string) *Regexp {
	r, err := CompileParse(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

func CompileParse(pattern string) (*Regexp, error) {
	regs.RLock()
	if re, ok := regs.cache[pattern]; ok {
		regs.RUnlock()
		return &Regexp{
			re: re,
		}, nil
	}
	regs.RUnlock()
	re, err := pcre.CompileParse(pattern)
	if err != nil {
		log.Printf("Otto Regexp error: %s", err)
		return nil, err
	}
	regs.Lock()
	regs.cache[pattern] = &re
	regs.Unlock()

	return &Regexp{
		re: &re,
	}, nil
}

func Compile(pattern, flags string) (*Regexp, error) {
	cacheKey := myutils.Concat(pattern, delim, flags)
	regs.RLock()
	if re, ok := regs.cache[cacheKey]; ok {
		regs.RUnlock()
		return &Regexp{
			re: re,
		}, nil
	}
	regs.RUnlock()
	_, flagC := pcre.ParseFlags(flags)
	re, err := pcre.Compile(pattern, flagC)
	if err != nil {
		p := UnescapeUnicode(pattern)
		re, unquotedErr := pcre.Compile(p, flagC)
		if unquotedErr == nil {
			return &Regexp{
				re: &re,
			}, nil
		} else {
			log.Printf("debug: pattern=[ %q ]\n\terr=[%s]\n\n", p, unquotedErr)
		}
		log.Printf("Otto Regexp error: %s\n\n", err)
		return nil, err
	}
	regs.Lock()
	regs.cache[cacheKey] = &re
	regs.Unlock()
	return &Regexp{
		re: &re,
	}, nil
}
