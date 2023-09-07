package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func BotInit() {
	session, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		fmt.Println("error creating discord session,", err)
		return
	}

	session.AddHandler(onMsgCreate)
	session.AddHandler(onReactionAdd)
	session.AddHandler(onReactionRemove)
	session.AddHandler(onReactionRemoveAll)
	session.AddHandler(onMsgUpdate)
	session.AddHandler(onMsgDelete)

	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageReactions
	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

func onMsgCreate(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.ChannelID != *ChannelID {
		return
	}
	//fmt.Println("Create")
	//fmt.Println(e.Message.Timestamp)
	fmt.Println(e.Attachments)
	for i := range e.Attachments {
		fmt.Println(e.Attachments[i])
	}
	db.Create(&Announcement{ID: e.Message.ID, Content: e.Content, TimePublished: e.Message.Timestamp})
}

func onReactionAdd(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
	//fmt.Println("AddReaction")
	//fmt.Println(e.MessageID)
	//fmt.Println(e.MessageReaction.Emoji.ID)
	//fmt.Println(e.MessageReaction.Emoji.Name)
	db.Create(&Reaction{EmojiID: e.Emoji.ID, EmojiName: e.Emoji.Name, MessageID: e.MessageID, UserID: e.UserID})
}

func onReactionRemove(s *discordgo.Session, e *discordgo.MessageReactionRemove) {
	//fmt.Println("RemoveReaction")
	//fmt.Println(e.MessageID)
	//fmt.Println(e.Emoji.ID)
	//fmt.Println(e.Emoji.Name)
	if e.Emoji.ID == "" {
		db.Delete(&Reaction{}, "message_id = ? AND user_id = ? AND emoji_name = ?", e.MessageID, e.UserID, e.Emoji.Name)
		return
	}
	db.Delete(&Reaction{}, "message_id = ? AND user_id = ? AND emoji_id = ?", e.MessageID, e.UserID, e.Emoji.ID)
}

func onReactionRemoveAll(s *discordgo.Session, e *discordgo.MessageReactionRemoveAll) {
	//fmt.Println("RemoveAllReaction")
	//fmt.Println(e.MessageID)
	db.Delete(&Reaction{}, "message_id = ?", e.MessageID)
}

func onMsgUpdate(s *discordgo.Session, e *discordgo.MessageUpdate) {
	if e.ChannelID != *ChannelID {
		return
	}
	//fmt.Println("Update")
	//fmt.Println(e.Message.ID)
	//fmt.Println(e.Message.EditedTimestamp)
	//db.Save(&Announcement{ID: e.Message.ID, Content: e.Message.Content, TimeEdited: *e.Message.EditedTimestamp})
	db.Model(&Announcement{ID: e.Message.ID}).Updates(Announcement{Content: e.Message.Content, TimeEdited: *e.Message.EditedTimestamp})
}

func onMsgDelete(s *discordgo.Session, e *discordgo.MessageDelete) {
	if e.ChannelID != *ChannelID {
		return
	}
	//fmt.Println("Delete")
	//fmt.Println(e.Message.ID)
	//db.Save(&Announcement{ID: e.Message.ID, Deleted: true})
	db.Delete(&Reaction{}, "message_id = ?", e.Message.ID)
	db.Delete(&Announcement{}, "id = ?", e.Message.ID)
}
