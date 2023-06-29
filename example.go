package main

/// This file provides the simple implementation
/// by which the reach data was collected (via dragonfly)

//// handler.go

// type handler struct {
// 	p *player.Player

// 	hits []float64

// 	player.NopHandler
// }

// func (h *handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
// 	*force, *height = 0.393, 0.398
// 	if a, ok := e.(*player.Player); ok {
// 		a.Heal(20, entity.FoodHealingSource{})
// 	}
// 	pos := entity.EyePosition(h.p)

// 	r, ok := trace.EntityIntercept(e, pos, pos.Add(h.p.Rotation().Vec3().Mul(10)))
// 	if !ok {
// 		return
// 	}

// 	dist := r.Position().Sub(pos).Len()

// 	var msg string

// 	if dist > 3.00 {
// 		msg = "<redstone>Reach (%dms): %.3f</redstone>"
// 	} else {
// 		msg = "<green>Reach (%dms): %.3f</green>"
// 	}

// 	h.hits = append(h.hits, dist)

// 	fmt.Println(h.hits)

// 	if len(h.hits) == 50 {
// 		var list []string
// 		for _, r := range h.hits {
// 			list = append(list, fmt.Sprintf("%.3f", r))
// 		}
// 		os.WriteFile(fmt.Sprintf("%d-hits.txt", (time.Now().UnixMilli())), []byte(strings.Join(list, "\n")), 0644)
// 		h.hits = []float64{}
// 		chat.Global.WriteString(text.Colourf("<gold>Hits reset</gold>"))
// 	}

// 	chat.Global.WriteString(text.Colourf(msg, h.p.Latency().Milliseconds()*2, dist))
// }
