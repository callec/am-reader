// Package html has the responsibility to find and execute
// templates associated with each page.
//
// E.g. viewer.go contains the struct ViewPageParams which
// can be used with the function ViewPage which then parses
// and executes the html template.
package html

type Parameters interface {
	MainPageParams | ViewPageParams
}
