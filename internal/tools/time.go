package tools

import "time"

func TimeFixRange(start, end string, startFixed, endFixed time.Duration) (s, e time.Time) {
	s, err := time.Parse(time.RFC3339, start)
	if err != nil {
		s = time.Now()
		s.Add(startFixed)
	}

	e, err = time.Parse(time.RFC3339, end)
	if err != nil {
		e = time.Now()
		e.Add(endFixed)
	}

	return
}
