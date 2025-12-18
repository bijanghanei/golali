# golali  
### Persian (Jalali / Shamsi / Solar Hijri) Calendar Library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/bijanghanei/golali.svg)](https://pkg.go.dev/github.com/bijanghanei/golali)
[![Go Report Card](https://goreportcard.com/badge/github.com/bijanghanei/golali)](https://goreportcard.com/report/github.com/bijanghanei/golali)
[![Tests](https://github.com/bijanghanei/golali/actions/workflows/go.yml/badge.svg)](https://github.com/bijanghanei/golali/actions)

**golali** is a comprehensive, production-ready Go library for working with the **Persian (Jalali / Shamsi / Solar Hijri) calendar** — the official calendar of Iran and Afghanistan.

The library provides accurate **bidirectional conversion** between Jalali and Gregorian dates, full **time zone support**, Persian month and weekday names, **custom formatting**, **strict parsing**, and **date arithmetic**, while integrating seamlessly with Go’s standard `time` package.

---

## Features

- Accurate Jalali ↔ Gregorian conversion  
  *(based on **Calendrical Calculations** by Dershowitz & Reingold)*
- Full `time.Time` interoperability  
  (`ToTime`, `Add`, `Sub`, `Unix`, etc.)
- Persian month and weekday names (full and abbreviated)
- Powerful custom formatting (strftime-like)
- Strict parsing with layout validation
- Date arithmetic:
  - `AddYears`
  - `AddMonths`
  - `AddDays`
  - `DaysInBetween`
- Correct leap year handling  
  *(Esfand has 30 days in leap years)*
- Time zone support with a convenient `IRST()` helper (Asia/Tehran)
- No external dependencies
- Thoroughly tested with **90%+ code coverage**

---

## Installation

```bash
go get github.com/bijanghanei/golali
```

---

## Quick Example

```go
package main

import (
	"fmt"

	"github.com/bijanghanei/golali"
)

func main() {
	// Current Jalali date and time
	now := golali.Now()
	fmt.Println("Today:", now.Format("%Y/%m/%d"))       // e.g. 1404/09/27
	fmt.Println("Full:", now.Format("%w، %B %d، %Y"))   // پنجشنبه، آذر 27، 1404

	// Create a specific date (Nowruz example)
	nowruz := golali.Date(
		1403,
		golali.Farvardin,
		1,
		0, 0, 0, 0,
		golali.IRST(),
	)
	fmt.Println("Nowruz 1403:", nowruz.Format("%B %d, %Y"))

	// Convert to Gregorian
	greg := nowruz.ToTime()
	fmt.Println("Gregorian:", greg.Format("2006-01-02")) // 2024-03-20

	// Parse from string
	parsed, _ := golali.Parse("YYYY/MM/DD HH:MM:SS", "1404/09/27 14:30:00")
	fmt.Println("Parsed:", parsed.Format("%Y/%m/%d %T"))
}
```

---

## Formatting Tokens

| Token | Description                          | Example        |
|------:|--------------------------------------|----------------|
| `%Y`  | 4-digit year                         | `1404`         |
| `%y`  | 2-digit year                         | `04`           |
| `%m`  | 2-digit month (01–12)                | `09`           |
| `%d`  | 2-digit day                          | `27`           |
| `%B`  | Full Persian month name              | `آذر`          |
| `%b`  | Abbreviated Persian month (3 chars)  | `آذ`           |
| `%w`  | Persian weekday name                 | `پنجشنبه`      |
| `%H`  | Hour (00–23)                         | `14`           |
| `%M`  | Minute (00–59)                       | `30`           |
| `%S`  | Second (00–59)                       | `45`           |
| `%R`  | `HH:MM`                              | `14:30`        |
| `%T`  | `HH:MM:SS`                           | `14:30:45`     |
| `%p`  | Day period                           | `صبح / عصر`   |
| `%z`  | Time zone offset                     | `+0330`        |
| `%Z`  | Time zone name                       | `Asia/Tehran`  |
| `%%`  | Literal `%`                          | `%`            |
| `%n`  | Newline                              | `\n`           |

---

## API Overview

### Constructors & Conversion

```go
golali.Date(
	year int,
	month Month,
	day, hour, min, sec, nsec int,
	loc *time.Location,
) JalaliDateTime

golali.Now() JalaliDateTime
golali.ToJalaliDateTime(t time.Time) JalaliDateTime
(j JalaliDateTime) ToTime() time.Time
```

### Formatting & Parsing

```go
(j JalaliDateTime) Format(layout string) string
(j JalaliDateTime) String() string         // "YYYY/MM/DD HH:MM:SS"
(j JalaliDateTime) FormatDateTime() string // "YYYY/MM/DD HH:MM"

golali.Parse(layout, value string) (JalaliDateTime, error)
golali.ParseInLocation(layout, value string, loc *time.Location) (JalaliDateTime, error)
```

### Date Arithmetic

```go
(j JalaliDateTime) Add(d time.Duration) JalaliDateTime
(j JalaliDateTime) AddYears(n int) JalaliDateTime
(j JalaliDateTime) AddMonths(n int) JalaliDateTime
(j JalaliDateTime) AddDays(n int) JalaliDateTime
(j JalaliDateTime) DaysInBetween(e JalaliDateTime) int
```

### Time Zones

```go
golali.IRST() *time.Location
```
