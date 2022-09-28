package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`(?m)ir(\d{1})(?P<named_label>\d{0,})(iu|il)(\d{0,})`)
	var carVinNumber = `ir63485458iu7664858`

	fmt.Printf("%#v\n", re.MatchString(carVinNumber))
	fmt.Printf("%#v\n", re.FindAllString(carVinNumber, -1))
	fmt.Printf("%#v\n", re.SubexpNames())
	fmt.Printf("%#v\n", re.SubexpIndex("named_label"))
	fmt.Printf("%#v\n", re.FindStringSubmatch(carVinNumber))

	var phoneRegexPattern = regexp.MustCompile(`\+(\d{0,2})(\d{3})(\d{7})`)
	var samplePhone = `+989124184801`

	fmt.Printf("%#v\n", phoneRegexPattern.MatchString(samplePhone))
	fmt.Printf("%#v\n", phoneRegexPattern.FindStringSubmatch(samplePhone))
}
