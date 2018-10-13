package main

import (
	"math/rand"
	"time"
)

func Swear() string {
	time.Sleep(5 * time.Millisecond)
	rand.Seed(time.Now().UTC().UnixNano())
	swears := []string{
		"cuntarse", "twatfuck", "dogknob", "fuck",
		"cum", "cunt", "arse", "felch",
		"bum", "shit", "motherfucker", "piss", "wank", "knob", "vomit", "penis", "twat",
		"cock", "dumb", "fucknut", "cuntface", "wipe", "bastard",
		"bugfucker", "assclown", "boner", "clitoris", "chode", "dickhead", "cuntrag",
		"dipshit", "nutsack", "pecker", "prick", "queef", "rimjob", "scrote", "shitbag",
		"skullfuck", "turd", "tit", "fart", "flaps", "anus", "brownhole",
		"dogfuck", "turd", "bottysausage", "wazz", "asshole",
		"taint", "barse", "bollocks", "jerk", "scrotum", "fist", "puke",
		"pustule", "mom", "vagina", "testes",
		"sweaty ballsacks", "pisspiece", "godwanger",
		"pus", "balls", "gobshite", "bong", "stink", "stench", "fuckarse",
		"sewage", "asshat", "shyster",
		"shitface", "ballbags", "Shitfunnel", "pubes",
		"cthulu-arse", "mothafuck",
		"shite", "punk-arse", "Custard", "Dead", "Cunty", "Fucked",
		"Shitted", "Arsecake", "Fcuk", "fuk", "fucker", "fuck",
		"Cock", "Slippy", "Fuckaroo", "cocks", "fucks",
		"ass-fucker", "asses", "assfucker", "assfukka", "asshole", "assholes", "asswhole",
		"Mini", "Funky", "Moo", "Weepy", "squelching", "Arsey",
		"Diahorretic", "Cockripping ", "Jerkwad",
		"Arseweeping", "Pissy", "ooze", "snot",
		"Cacky", "Fuck", "Skate", "Dirty", "FUCK",
		"King", "Queen", "Princess", "Shitty", "Rat", "fuck",
		"Penis", "Knob", "Horse", "Hepatitis",
		"Shite", "Turd", "Spit", "Fuckknocking", "fucko", "cunts", "shits",
	}
	return swears[(rand.Intn(len(swears)))]
}

func Punc() string {
	marks := []string{"?", "!", ".", "!!!", "!?"}
	return marks[(rand.Intn(len(marks)))]
}
