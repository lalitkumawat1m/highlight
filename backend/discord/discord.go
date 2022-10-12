package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type HighlightBot struct {
	Session *discordgo.Session
	GuildID string
}

func InitBot(guildId string) (*HighlightBot, error) {
	var (
		ok               bool
		DiscordBotSecret string
	)
	if DiscordBotSecret, ok = os.LookupEnv("DISCORD_BOT_SECRET"); !ok || DiscordBotSecret == "" {
		return nil, errors.New("DISCORD_BOT_SECRET not set")
	}

	session, err := discordgo.New("Bot " + DiscordBotSecret)

	if err != nil {
		return nil, errors.Wrap(err, "error creating Discord session")
	}

	return &HighlightBot{
		Session: session,
		GuildID: guildId,
	}, nil
}
