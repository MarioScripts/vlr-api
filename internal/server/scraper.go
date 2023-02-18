package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
	"github.com/gocolly/colly"
)

func getMatches(in *pb.MatchesRequest) []*pb.Match {
	c := colly.NewCollector()
	var matches = make([]*pb.Match, 0)
	stop := false

	c.OnHTML("a[href].match-item", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		status := buildMatchStatus(e.ChildText(".match-item-eta  .ml-status"))

		// Don't traverse to matches that don't match requested status
		if in.MatchStatus != pb.MatchStatus_NONE && status != in.MatchStatus {
			return
		}

		if stop {
			return
		}

		match := getMatch(e.Request.AbsoluteURL(link))

		// Don't traverse to matches that don't match requested date
		if in.Date != nil && match.Date.AsTime().After(in.Date.AsTime()) {
			stop = true
		}

		if !stop {
			matches = append(matches, match)
		}

	})

	c.Visit("https://www.vlr.gg/matches")

	c.Wait()

	return matches
}

func getMatchFromId(id int64) *pb.Match {
	return getMatch(fmt.Sprintf("https://www.vlr.gg/%d", id))
}

// TODO: Handle BO1s and BO5s
func getMatch(url string) *pb.Match {
	c := colly.NewCollector()
	var match = &pb.Match{}

	c.OnHTML(".col-container", func(e *colly.HTMLElement) {
		t1 := buildTeam(e, 1)
		t2 := buildTeam(e, 2)

		// Id
		id := getIdFromUrl(e.Request.URL.Path)

		// Status
		rawStatus := e.ChildText(".match-header-vs-score > .match-header-vs-note:nth-child(1)")
		status := buildMatchStatus(rawStatus)

		// Maps
		maps := make([]string, 0)
		e.ForEach(".vm-stats-gamesnav > :not(div:nth-child(1)) > div", func(i int, e *colly.HTMLElement) {
			m := strings.ReplaceAll(cleanText(e.Text), strconv.Itoa(i+1), "")

			if m != "TBD" {
				maps = append(maps, m)
			}
		})

		// Date
		rawDate := e.ChildAttr(".match-header-date > .moment-tz-convert:nth-child(1)", "data-utc-ts")
		loc, _ := time.LoadLocation("America/New_York")
		date, err := time.ParseInLocation("2006-01-02 15:04:05", rawDate, loc)
		if err != nil {
			log.Fatalf("Envountered error parsing date: %v\n", err)
		}

		// Tournament
		tournament := e.ChildText("a.match-header-event > div > div:not(.match-header-event-series)")
		tId := getIdFromUrl(e.ChildAttr("a.match-header-event", "href"))

		match = &pb.Match{
			Id:          int64(id),
			TeamOne:     t1,
			TeamTwo:     t2,
			MatchStatus: status,
			Maps:        maps,
			Date:        timestamppb.New(date),
			Tournament: &pb.Tournament{
				Id:   tId,
				Name: tournament,
			},
		}
	})

	c.Visit(url)

	c.Wait()

	return match
}

func buildTeam(e *colly.HTMLElement, num int) *pb.Team {
	info := fmt.Sprintf(".match-header-vs > .mod-%d", num)
	name := cleanText(e.ChildText(info + " .wf-title-med"))
	icon := e.ChildAttr(info+" > img", "src")

	m1 := regexp.MustCompile(`[^0-9]`)
	link, _ := strconv.Atoi(m1.ReplaceAllString(e.ChildAttr(info, "href"), ""))

	scoreIndex := 1
	if num == 2 {
		scoreIndex = 3
	}

	scoreTemp, err := strconv.Atoi(e.ChildText(fmt.Sprintf(".match-header-vs-score > div > div > span:nth-child(%d)", scoreIndex)))
	var score = 0

	if err == nil {
		score = scoreTemp
	}

	return &pb.Team{
		Id:    int64(link),
		Name:  name,
		Icon:  icon,
		Score: int32(score),
	}
}

func cleanText(text string) string {
	return strings.ReplaceAll(strings.ReplaceAll(text, "\t", ""), "\n", "")
}

func buildMatchStatus(text string) pb.MatchStatus {
	upperText := strings.ToUpper(text)

	if upperText == "LIVE" {
		return pb.MatchStatus_LIVE
	} else if upperText == "FINAL" {
		return pb.MatchStatus_FINISHED
	}

	return pb.MatchStatus_NOT_STARTED
}

func getIdFromUrl(url string) int64 {
	m1 := regexp.MustCompile(`\/?\d+\/?`)
	regexedString := strings.TrimSpace(strings.ReplaceAll(m1.FindString(url), "/", ""))
	id, err := strconv.Atoi(regexedString)
	if err != nil {
		log.Fatalf("Failed to get id from url %s, %v\n", url, err)
	}

	return int64(id)
}
