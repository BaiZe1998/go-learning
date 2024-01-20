package main

type basic struct {
	maxLife int
	life    int
	attack  int
	defense int
}

func (b *basic) isAlive() bool {
	return b.life > 0
}

func (b *basic) attacked(attackPower int) int {
	if b.defense > attackPower {
		return 0
	}

	deduct := attackPower - b.defense
	b.life -= deduct
	if b.life < 0 {
		b.life = 0
	}
	return deduct
}
