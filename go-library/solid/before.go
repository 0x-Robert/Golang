type Player struct {
	name string 
}

type Monster struct{
	hp int 
}

type (p *Player) Name() string {
	return p.name 
}

type (p *Player) Attack(m *Monster){
	m.DealDamage(p, 100)
}

func (m *Monster) DealDamage(attacker *Player, damage int){
	m.hp -= damage 
	if m.hp < 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다.")
	}
}