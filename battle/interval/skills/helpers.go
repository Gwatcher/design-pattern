package skills

import "designer/battle/interval/core"

func pickLowestHPEnemy(self core.Unit, all []core.Unit) core.Unit {
	var best core.Unit
	for _, u := range all {
		if !u.IsAlive() || u.Team() == self.Team() {
			continue
		}
		if best == nil || u.HP() < best.HP() {
			best = u
		}
	}
	return best
}

func pickLowestHPAlly(self core.Unit, all []core.Unit) core.Unit {
	var best core.Unit
	for _, u := range all {
		if !u.IsAlive() || u.Team() != self.Team() {
			continue
		}
		if best == nil || u.HP() < best.HP() {
			best = u
		}
	}
	return best
}

func itoa(x int) string {
	if x == 0 {
		return "0"
	}
	neg := x < 0
	if neg {
		x = -x
	}
	var buf [20]byte
	i := len(buf)
	for x > 0 {
		i--
		buf[i] = byte('0' + x%10)
		x /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
