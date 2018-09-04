package bot

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/fuel-/gideon/config"

	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session

//Start will start our bot a runnin'
func Start() {
	fmt.Println("howdy")
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	botID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	missionTimeGet()
}

func missionTimeGet() (hours, minutes int) {
	loc, _ := time.LoadLocation("America/Chicago")
	now := time.Now().In(loc)
	fmt.Println("\nToday : ", loc, " Time : ", now)

	guildMissionDay := "Saturday"
	weekday := now.Weekday()
	daysToAdd := 0
	if weekday.String() != guildMissionDay {
		weekdayInt := int(weekday)
		daysToAdd = 6 - weekdayInt
	}
	future := now.AddDate(0, 0, daysToAdd)
	year := now.Year()
	month := now.Month()
	day := future.Day()

	fmt.Printf("current weekday is %v \n\n\n", weekday)
	fmt.Printf("###############################################################\n\n\n")

	futureDate := time.Date(year, month, day, 20, 00, 00, 000, loc)
	fmt.Println("Future  : ", loc, " Time : ", futureDate) //
	fmt.Printf("###############################################################\n")
	diff := futureDate.Sub(now)

	hrs := int(diff.Hours())
	fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)

	mins := int(diff.Minutes())
	fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)

	sec := int(diff.Seconds())
	fmt.Printf("Difference in Seconds: %d Seconds\n", sec)

	return hrs, mins
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := ""
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == botID {
			return
		}

		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong!")
		} else if m.Content == "!gm" {
			hrs, mins := missionTimeGet()
			preMinutesLeft := mins / 60
			minutesLeft := (preMinutesLeft * 60) - mins

			newMinutesLeft := math.Abs(float64(minutesLeft))
			if math.Signbit(float64(hrs)) {
				message = fmt.Sprintf("Sorry! Guild misssions started %d and %.0f ago.", hrs, newMinutesLeft)

			} else {
				message = fmt.Sprintf("Guild missions start in %d hours and %.0f minutes.", hrs, newMinutesLeft)
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, message)
		} else if m.Content == "!help" {
			message = "Currently my list of commands are !ping, !gm and !brad"
			_, _ = s.ChannelMessageSend(m.ChannelID, message)
		} else if m.Content == "!brad" {
			message = "is dad."
			_, _ = s.ChannelMessageSend(m.ChannelID, message)
		} else if m.Content == "!goodbot" {
			message = "Aww, Thank you! https://media.giphy.com/media/1gbQIeNzZxcSk/giphy.gif"
			_, _ = s.ChannelMessageSend(m.ChannelID, message)
		}
		fmt.Println(m.Content)
	}

}
