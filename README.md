# time_iso

A time formatting module for golang with iso 8601 in mind.

## Motivation

In order to format time in golang you have to give a layout defined by a specific time, e.g.

```go
now := time.Now()
now.Format("Mon, 02 Jan 2006 15:04:05 MST")
```

And yeah, it **has** to be January second, 2006 at 23:04:05 UTC (but in MST)

## Usage / Example

To do the same thing in the [Motivation](#motivation) we'll do:

```go
now := time_iso.Time{time.Now()} // or time_iso.Now()
now.Format("DDD, DD MMM YYYY hh:mm:ss ZZZ")
```

### Formats

- DD: A two-digit day of that month, 01 through 31
- DDD: A three-letter abbreviation for day of the week, e.g. Fri
- DDDD: A day of the week spelled out in full, e.g. Friday
- MM: A two-digit month of the year, 01 through 12
- MMM: A three-letter abbreviation for month, e.g. Mar
- YY: A two-digit year, e.g. 21
- YYYY: A four-digit year, 0000 through 9999
- Z: The zone designator for the zero UTC offset in the form ±\[hh\]\[mm\]
- ZZZ: The zone designator name, e.g. MST
- \_D: One-digit day of the month for days below 10, e.g. \_2
- hh: A zero-padded hour between 00 and 23
- mm: A zero-padded minute between 00 and 59
- ss: A zero-padded second between 00 and 60 (where 60 is only used to denote an added leap second
- sss: A zero-padded millisecond between 000 and 999
