package parsedmvresp

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// GetAppointmentTime will parse DMV returned HTML content string to extract time. it returns time.Time, error
func GetAppointmentTime(s string, now time.Time) (time.Time, error) {

	oneYearLater := now.AddDate(1, 0, 0)

	if strings.Contains(s, "Sorry, all appointments at this office are currently taken") || strings.Contains(s, "no appointment is available") {
		return oneYearLater, nil
	}

	r, _ := regexp.Compile("(\\w+), ((\\w+) (\\d{1,2}), (20\\d{2}) at (\\d{1,2}:\\d{2}) (AM|PM))")
	r2d, _ := regexp.Compile("(\\w+) (\\d{1,2}), (20\\d{2})")
	r2h, _ := regexp.Compile("(\\d{1,2}:\\d{2}) (AM|PM)")

	match := r.FindStringSubmatch(s)
	m2d := r2d.FindStringSubmatch(s)
	m2h := r2h.FindStringSubmatch(s)

	ds := ""
	if len(match) > 2 {
		ds = match[2]
	}
	if ds == "" && len(m2d) > 0 && len(m2h) > 0 {
		ds = fmt.Sprintf("%s at %s", m2d[0], m2h[0])
	}

	if len(ds) == 0 {
		_ = logReponse(s, "debugReponse.html")
		log.Println("error:", "No datetime string found in return")
		return oneYearLater, errors.New("No datetime string found in return")
	}

	t, err := time.Parse("January 2, 2006 at 15:04 PM -0700", ds+" -0700")
	if err != nil {
		t, err = time.Parse("Jan 2, 2006 at 15:04 PM -0700", ds+" -0700")
	}

	if err != nil {
		log.Printf("ds: %v error: %v", ds, err)
		return oneYearLater, err
	}

	return t, nil
}

func logReponse(s string, fn string) error {
	if f, errCreate := os.Create(fn); errCreate != nil {
		log.Println(errCreate)
		return errCreate
	} else if _, errWrite := f.WriteString(s); errWrite != nil {
		log.Println(errWrite)
		f.Close()
		return errWrite
	} else {
		f.Close()
		return nil
	}
}
