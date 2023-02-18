package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
	"github.com/gocolly/colly"
)

func (s *Server) Matches(ctx context.Context, in *pb.MatchesRequest) (*pb.MatchesResponse, error) {
	return &pb.MatchesResponse{
		Matches: getMatches(),
	}, nil
}

func getMatches() []*pb.Match {
	c := colly.NewCollector()
	var matches = make([]*pb.Match, 0)

	c.OnHTML("a[href].match-item", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// fmt.Printf("Found match: %v\n", link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML(".col-container", func(e *colly.HTMLElement) {
		t1 := buildTeam(e, 1)
		t2 := buildTeam(e, 2)

		// Id
		m1 := regexp.MustCompile(`^\/(\d*)\/`)
		idString := strings.ReplaceAll(m1.FindString(e.Request.URL.Path), "/", "")
		id, _ := strconv.Atoi(idString)

		// Maps
		maps := make([]string, 0)
		e.ForEach(".vm-stats-gamesnav > :not(div:nth-child(1)) > div", func(i int, e *colly.HTMLElement) {
			m := strings.ReplaceAll(cleanText(e.Text), strconv.Itoa(i+1), "")

			if m != "TBD" {
				maps = append(maps, m)
			}
		})

		match := &pb.Match{
			Id:             int64(id),
			TeamOne:        t1,
			TeamTwo:        t2,
			Maps:           maps,
			TournamentName: "Test",
		}

		matches = append(matches, match)
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	c.Visit("https://www.vlr.gg/matches")

	c.Wait()

	return matches
}

func buildTeam(e *colly.HTMLElement, num int) *pb.Team {
	info := fmt.Sprintf(".match-header-vs > .mod-%d", num)
	name := e.ChildText(info)
	icon := e.ChildAttr(info+" > img", "src")

	m1 := regexp.MustCompile(`[^0-9]`)
	link, _ := strconv.Atoi(m1.ReplaceAllString(e.ChildAttr(info, "href"), ""))

	scoreIndex := 1
	if num == 2 {
		scoreIndex = 3
	}

	scoreTemp, err := strconv.Atoi(e.ChildText(fmt.Sprintf(".match-header-vs-score > div > span:nth-child(%d)", scoreIndex)))
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
