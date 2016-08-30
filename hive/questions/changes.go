package questions

import (
	"log"

	"github.com/black-banana/bee-hive/hive/hub"
)

const (
	QuestionChanged = "question-changed"
	QuestionNew     = "question-new"
)

func GetAllChanges(h *hub.Hub) {
	go func(h *hub.Hub) {
		for {
			res, err := repository.Table().Changes().Run(repository.Session)
			if err != nil {
				log.Fatal(err)
			}
			var questionChanges interface{}
			for res.Next(&questionChanges) {
				var msg = &hub.Message{
					Type:    QuestionChanged,
					Content: &questionChanges,
				}
				h.Broadcast(msg)
			}
			if res.Err() != nil {
				log.Println(res.Err())
			}
		}
	}(h)
}
