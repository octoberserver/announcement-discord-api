package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func ServeAPI() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("歡迎來到十月模組伺服器公告API"))
	})

	r.Get("/announcements", func(w http.ResponseWriter, r *http.Request) {
		type Result struct {
			ID            string
			Content       string
			TimePublished time.Time
			TimeEdited    time.Time
			EmojiID       string
			EmojiName     string
			ReactionCount uint
		}
		type ResReaction struct {
			EmojiID       string
			EmojiName     string
			ReactionCount uint
		}
		type Response struct {
			Content       string
			TimePublished time.Time
			TimeEdited    time.Time
			Reactions     []ResReaction
		}

		var results []Result
		db.Raw(`
			SELECT a.*, r.emoji_id, r.emoji_name, COUNT(r.emoji_name OR r.emoji_id) FROM announcements a
			LEFT JOIN reactions r ON r.message_id = a.id
			GROUP BY a.id, a.time_published, r.emoji_name, r.emoji_id
			ORDER BY a.time_published, a.id, COUNT(r.emoji_name OR r.emoji_id) DESC;
		`).Scan(&results)

		responseMap := make(map[string]Response)

		for i := range results {
			result := results[i]
			id := result.ID
			_, exists := responseMap[id]
			if !exists {
				responseMap[id] = Response{
					Content:       result.Content,
					TimePublished: result.TimePublished,
					TimeEdited:    result.TimeEdited,
					Reactions: []ResReaction{
						ResReaction{
							EmojiID:       result.EmojiID,
							EmojiName:     result.EmojiName,
							ReactionCount: result.ReactionCount,
						},
					},
				}
			} else {
				res, _ := responseMap[id]
				res.Reactions = append(res.Reactions, ResReaction{
					EmojiID:       result.EmojiID,
					EmojiName:     result.EmojiName,
					ReactionCount: result.ReactionCount,
				})
				responseMap[id] = res
			}
		}

		output, err := json.Marshal(responseMap)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(output)
	})

	//r.Get("/announcements/all", func(w http.ResponseWriter, r *http.Request) {
	//	var announcements []Announcement
	//	db.Find(&announcements)
	//
	//	output, err := json.Marshal(announcements)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	w.Write(output)
	//})

	http.ListenAndServe(":3000", r)
}
