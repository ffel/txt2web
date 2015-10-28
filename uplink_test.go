package txt2web

import "testing"

/*
what's next .... (after a long period of silence)

this is handy: pythia -http=:8081 github.com/ffel/txt2web

volgens mij is er een werkend principe in uplink, het enige is dat deze
alleen nog maar naar console print en niet de content in chunk wijzigt
(kan ik dat bevestigen?).

Dan zou ik na moeten gaan hoe de structuur van de section header struct
wijzigt door er een link in te stoppen.

Is het een idee om een json methode toe te voegen aan chunk? Voor nu
kunnen we eenvoudigweg testen hoe een header met link er uit ziet met,
.... we hadden toch een tooltje??

`t2tree tmp/header.txt` levert onderstaande tree op. Er zitten twee
headers in, op 1,0 en op 1,1. Het verschil zit m in 1,x,c,2. Voor `x==0`
is dit een lange list met de elementen van de titel. Voor `x==1` is dit
element een korte lijs met in 1,1,c,2,0 een Link CT object.

    + "": list:
        + "0": map:
            + "unMeta": map:
        + "1": list:
            + "0": map:
                + "c": list:
                    + "0": value: 1
                    + "1": list:
                        + "0": value: dit-is-een-normale-header
                        + "1": list:
                        + "2": list:
                    + "2": list:
                        + "0" - Str: "dit"
                        + "1" - Space
                        + "2" - Str: "is"
                        + "3" - Space
                        + "4" - Str: "een"
                        + "5" - Space
                        + "6" - Str: "normale"
                        + "7" - Space
                        + "8" - Str: "header"
                + "t": value: Header
            + "1": map:
                + "c": list:
                    + "0": value: 1
                    + "1": list:
                        + "0": value: header-in-link
                        + "1": list:
                        + "2": list:
                    + "2": list:
                        + "0": map:
                            + "c": list:
                                + "0": list:
                                    + "0" - Str: "header"
                                    + "1" - Space
                                    + "2" - Str: "in"
                                    + "3" - Space
                                    + "4" - Str: "link"
                                + "1": list:
                                    + "0": value: http://example.com
                                    + "1": value:
                            + "t": value: Link
                + "t": value: Header

De manier om een object te vervangen zou `deepreader.SetObject` moeten
zijn, maar ik vind geen voorbeeld van gebruik van deze code. (SetString
wordt wel gebruikt).

Maar ergens wordt de link van de header toch gezet (op z'n minst in de
ouder code).

Bestand `index.go` rond regel 130 volgt een andere benadering: ik
schrijf daar markdown en laat daar pandoc de json bouwen.

> oh wacht, index.go genereert complete markdown bestanden. Maar ik kan
> op z'n minst zoeken naar andere voorbeelden van pandoc gebruik waarbij
> een gedeelte van de contents wordt ingevoegd.

Er is een voorbeeld (een oud voorbeeld) van een wijziging in de
structuur in `ffel/pandocfilter/cmd/smath`. Er is daar een functie
`WrapTMath` waarin een wrapper wordt gedefinieerd. Niet echt fijne code,
maar voorlopig kunnen we het zo even doen. (werk nog wel!)

Toch is een ander alternatief dat ik van de json weer terug ga naar
markdown: hier doe ik een aanpassing die ik vervolgens weer in json
converteer - een beetje een omweg, maar de operatie is wellicht
gemakkelijk.

*/

var uplink_tests []string = []string{`
main
====

sub A
-----

### Sub A.A

sub B
-----

### Sub B.A

### Sub B.B

main II
=======
`, `
foo
`, `
### weird a

# weird b

###### weird c

###### weird c2

### weird d
`, `
bar
`, `
# h1
# [h2](#h1)
`, `
baz
`,
}

// for one test only
// go test -run TestUpLinks

func TestUpLinks(t *testing.T) {
	inout := []struct{ in, out string }{
		{uplink_tests[0], uplink_tests[1]},
		// {uplink_tests[2], uplink_tests[3]},
		// {uplink_tests[4], uplink_tests[5]}, // bedoeld om verschil in structuur te vinden
	}

	for _, tt := range inout {
		c := markdownChan(UpLinkNode(setFiles(contentGen(tt.in), "file.txt")))

		got := string(<-c)

		if got != tt.out {
			t.Errorf("---- expected:\n%q\n---- got:\n\n%q", tt.out, got)
		}
	}
}
