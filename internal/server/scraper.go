package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
	"github.com/gocolly/colly"
)

var vlrBaseUrl = "https://www.vlr.gg"

func getMatches(in *pb.MatchesRequest) ([]*pb.Match, error) {
	c := colly.NewCollector()
	var matches = make([]*pb.Match, 0)
	stop := false

	c.OnHTML("a[href].match-item", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		status := buildMatchStatus(e.ChildText(".match-item-eta  .ml-status"))

		// Don't traverse to matches that don't match requested status
		if (in.MatchStatus != pb.MatchStatus_NONE && status != in.MatchStatus) || (in.TournamentId != 0 && status == pb.MatchStatus_FINISHED && in.MatchStatus != pb.MatchStatus_FINISHED) {
			return
		}

		if stop {
			return
		}

		match, err := getMatch(e.Request.AbsoluteURL(link))

		// Don't traverse to matches that don't match requested date
		if in.Date != nil && match.Date.AsTime().After(in.Date.AsTime()) {
			stop = true
		}

		if !stop && err == nil {
			matches = append(matches, match)
		}

	})

	if in.TournamentId != 0 {
		c.Visit(fmt.Sprintf("%s/event/matches/%d?series_id=all", vlrBaseUrl, in.TournamentId))
	} else {
		c.Visit(fmt.Sprintf("%s/matches", vlrBaseUrl))
	}

	c.Wait()

	return matches, nil
}

func getMatchFromId(id int64) (*pb.Match, error) {
	return getMatch(fmt.Sprintf("%s/%d", vlrBaseUrl, id))
}

func getTeam(in *pb.IdRequest) (*pb.TeamResponse, error) {
	c := colly.NewCollector()
	var team = &pb.TeamResponse{}
	var teamErr = error(nil)

	c.OnHTML(".col-container", func(e *colly.HTMLElement) {
		icon := strings.ReplaceAll(e.ChildAttr(".team-header-logo img", "src"), "//", "https://")
		name := e.ChildText(".team-header-name > h1.wf-title")
		country := strings.ToUpper(cleanText(e.ChildText(".team-header-country")))
		players := getPlayers(in.GetId())

		team = &pb.TeamResponse{
			Id:      in.GetId(),
			Icon:    icon,
			Name:    name,
			Country: country,
			Players: players,
		}
	})

	c.Visit(fmt.Sprintf("%v/team/%v", vlrBaseUrl, in.GetId()))
	c.Wait()

	return team, teamErr
}

func getPlayers(teamId int64) []*pb.Player {
	c := colly.NewCollector()
	var players = make([]*pb.Player, 0)

	c.OnHTML(".team-roster-item a[href]", func(e *colly.HTMLElement) {
		id, err := getIdFromUrl(e.Attr("href"))
		if err != nil {
			return
		}
		players = append(players, getPlayer(id, false))
	})

	c.Visit(fmt.Sprintf("%v/team/%v", vlrBaseUrl, teamId))
	c.Wait()
	return players
}

func getPlayer(pid int64, showTeams bool) *pb.Player {
	c := colly.NewCollector()
	player := &pb.Player{}
	var teams = make([]*pb.SimpleTeam, 0)

	c.OnHTML(".col-container", func(e *colly.HTMLElement) {
		typeText := strings.ToUpper(e.ChildText(".player-summary-container-1 > .wf-card:nth-of-type(4) > a  span"))
		var pType pb.PlayerType = pb.PlayerType_PLAYER
		fmt.Println(typeText)
		if typeText == "COACH" || typeText == "HEAD COACH" {
			pType = pb.PlayerType_HEAD_COACH
		} else if typeText == "ASSISTANT COACH" {
			pType = pb.PlayerType_ASSISTANT_COACH
		} else if typeText == "MANAGER" {
			pType = pb.PlayerType_MANAGER
		} else if typeText == "ANALYST" {
			pType = pb.PlayerType_ANALYST
		}

		player = &pb.Player{
			Id:      pid,
			Name:    e.ChildText(".player-real-name"),
			Handle:  e.ChildText(".wf-title"),
			Country: cleanText(e.ChildText("div.player-header div.ge-text-light")),
			Type:    pType,
		}
	})

	if showTeams == true {
		c.OnHTML("a[href*=\"/team\"]", func(e *colly.HTMLElement) {
			id, err := getIdFromUrl(e.Attr("href"))
			if err == nil {
				teams = append(teams, &pb.SimpleTeam{
					Name: cleanText(e.ChildText("div:nth-child(2) > div:nth-child(1)")),
					Id:   id,
				})
				player.Teams = teams

			} else {
				fmt.Println(err)
			}
		})
	}

	c.Visit(fmt.Sprintf("%v/player/%v", vlrBaseUrl, pid))
	c.Wait()

	return player
}

func getMatch(url string) (*pb.Match, error) {
	c := colly.NewCollector()
	var match = &pb.Match{}
	var matchErr = error(nil)

	c.OnHTML(".col-container", func(e *colly.HTMLElement) {
		// Id
		id, err := getIdFromUrl(e.Request.URL.Path)
		if err != nil {
			match = &pb.Match{}
			matchErr = err
			return
		}

		// Match Teams
		t1, err1 := buildMatchTeam(e, 1)
		t2, err2 := buildMatchTeam(e, 2)
		if err1 != nil {
			match = &pb.Match{}
			matchErr = err1
			return
		}

		if err2 != nil {
			match = &pb.Match{}
			matchErr = err2
			return
		}

		// Best of
		bestOf := int32(stringToInt64(e.ChildText(".match-header-vs-score > .match-header-vs-note:nth-child(3)")))

		// Status
		rawStatus := e.ChildText(".match-header-vs-score > .match-header-vs-note:nth-child(1)")
		status := buildMatchStatus(rawStatus)

		// Maps
		maps := make([]string, 0)
		if bestOf == 1 {
			m := cleanText(e.ChildText(".vm-stats-game-header .map > div:nth-child(1)"))
			if m != "TBD" {
				maps = append(maps, m)
			}
		} else {
			e.ForEach(".vm-stats-gamesnav > :not(div:nth-child(1)) > div", func(i int, e *colly.HTMLElement) {
				m := strings.ReplaceAll(cleanText(e.Text), strconv.Itoa(i+1), "")

				if m != "TBD" {
					maps = append(maps, m)
				}
			})
		}

		// Date
		rawDate := e.ChildAttr(".match-header-date > .moment-tz-convert:nth-child(1)", "data-utc-ts")
		loc, _ := time.LoadLocation("America/New_York")
		date, err := time.ParseInLocation("2006-01-02 15:04:05", rawDate, loc)
		if err != nil {
			log.Fatalf("Envountered error parsing date: %v\n", err)
		}

		// Tournament
		tournamentName := e.ChildText("a.match-header-event > div > div:not(.match-header-event-series)")
		tId, err := getIdFromUrl(e.ChildAttr("a.match-header-event", "href"))
		if err != nil {
			match = &pb.Match{}
			matchErr = err
			return
		}

		match = &pb.Match{
			Id:          int64(id),
			TeamOne:     t1,
			TeamTwo:     t2,
			MatchStatus: status,
			Maps:        maps,
			Date:        timestamppb.New(date),
			Tournament: &pb.Tournament{
				Name: tournamentName,
				Id:   tId,
			},
			BestOf: bestOf,
		}
	})

	c.Visit(url)

	c.Wait()

	return match, matchErr
}

func buildMatchTeam(e *colly.HTMLElement, num int) (*pb.MatchTeam, error) {
	info := fmt.Sprintf(".match-header-vs > .mod-%d", num)
	name := cleanText(e.ChildText(info + " .wf-title-med"))
	icon := strings.ReplaceAll(e.ChildAttr(info+" > img", "src"), "//", "https://")

	id, err := getIdFromUrl((e.ChildAttr(info, "href")))
	if err != nil {
		return nil, err
	}

	scoreIndex := 1
	if num == 2 {
		scoreIndex = 3
	}

	scoreTemp, err := strconv.Atoi(e.ChildText(fmt.Sprintf(".match-header-vs-score > div > div > span:nth-child(%d)", scoreIndex)))
	var score = 0

	if err == nil {
		score = scoreTemp
	}

	return &pb.MatchTeam{
		Id:    id,
		Name:  name,
		Icon:  icon,
		Score: int32(score),
	}, nil
}

func cleanText(text string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(text, "\t", ""), "\n", ""))
}

func buildMatchStatus(text string) pb.MatchStatus {
	upperText := strings.ToUpper(text)

	if upperText == "LIVE" {
		return pb.MatchStatus_LIVE
	} else if upperText == "FINAL" || upperText == "COMPLETED" {
		return pb.MatchStatus_FINISHED
	}

	return pb.MatchStatus_NOT_STARTED
}

func stringToInt64(text string) int64 {
	m1 := regexp.MustCompile(`\d+`)
	number := m1.FindString(text)
	num, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("Could not convert string field %v to int %v\n", number, err)
		return int64(0)
	}

	return int64(num)
}

func getIdFromUrl(url string) (int64, error) {
	m1 := regexp.MustCompile(`\/?\d+\/?`)
	regexedString := strings.TrimSpace(strings.ReplaceAll(m1.FindString(url), "/", ""))
	id, err := strconv.Atoi(regexedString)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to get id from url %s, %v\n", url, err)
		fmt.Print(errMsg)
		return -1, status.Error(codes.Internal, errMsg)
	}

	return int64(id), nil
}
