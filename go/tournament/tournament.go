package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"text/tabwriter"
)

// Standing keeps track of a teams matches and points
type Standing struct {
	Team string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

// ByScore is a struct that helps sort team standings based on score.
// it is deterministing. Two teams with the same amount of points will be
// ordered alphabetically.
type ByScore []Standing

func (s *Standing) String() string {
	return fmt.Sprintf("%s       \t  %d\t  %d\t  %d\t  %d\t  %d",
		s.Team, s.MP, s.W, s.D, s.L, s.P)
}

func (a ByScore) Len() int      { return len(a) }
func (a ByScore) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool {
	if a[i].P == a[j].P {
		return strings.Compare(a[i].Team, a[j].Team) < 1
	}

	return a[i].P > a[j].P
}

// Tally returns a print friendly representation of a match
func Tally(reader io.Reader, buffer io.Writer) error {
	s, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	standings := make(map[string]*Standing)
	lines := strings.Split(string(s), "\n")
	for _, line := range lines {
		// fmt.Printf("line: [%s]\n", line)
		if len(line) > 3 {
			err := analyze(line, standings)
			if err != nil {
				return err
			}
		}
	}

	return format(standings, &buffer)
}

func analyze(line string, standings map[string]*Standing) error {
	parts := strings.Split(line, ";")
	if len(parts) != 3 {
		return nil
	}

	homeTeam := strings.TrimSpace(parts[0])
	awayTeam := strings.TrimSpace(parts[1])
	outcome := parts[2]

	if standings[homeTeam] == nil {
		standings[homeTeam] = &Standing{Team: homeTeam}
	}

	if standings[awayTeam] == nil {
		standings[awayTeam] = &Standing{Team: awayTeam}
	}

	return score(homeTeam, awayTeam, outcome, standings)
}

func score(homeTeam string, awayTeam string, outcome string,
	standings map[string]*Standing) error {
	home := standings[homeTeam]
	away := standings[awayTeam]
	home.MP++
	away.MP++

	switch {
	case outcome == "win":
		home.W++
		away.L++
	case outcome == "loss":
		home.L++
		away.W++
	case outcome == "draw":
		home.D++
		away.D++
	default:
		return errors.New("win, loss, or draw not specified")
	}

	home.P = home.W*3 + home.D
	away.P = away.W*3 + away.D
	return nil
}

func format(standings map[string]*Standing, buffer *io.Writer) error {
	if len(standings) == 0 {
		return errors.New("No teams to score")
	}

	var sortedStandings []Standing
	for _, v := range standings {
		sortedStandings = append(sortedStandings, *v)
	}
	sort.Sort(ByScore(sortedStandings))

	w := tabwriter.NewWriter(*buffer, 0, 4, 1, ' ', tabwriter.TabIndent|tabwriter.Debug)
	fmt.Fprintln(w, "Team\t MP\t  W\t  D\t  L\t  P")

	for _, standing := range sortedStandings {
		fmt.Fprintln(w, standing.String())
	}
	w.Flush()

	return nil
}
