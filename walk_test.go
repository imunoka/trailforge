package trailforge

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{Name string}{"imunoka"},
			[]string{"imunoka"},
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"imunoka", "Seoul"},
			[]string{"imunoka", "Seoul"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age int
			}{"imunoka", 17},
			[]string{"imunoka"},
		},
		{
			"nested fields",
			Person{
				"imunoka",
				Profile{17, "Seoul"},
			},
			[]string{"imunoka", "Seoul"},
		},
		{
			"pointers to things",
			&Person{
				"imunoka",
				Profile{33, "Seoul"},
			},
			[]string{"imunoka", "Seoul"},
		},
		{
			"slices",
			[]Profile{
				{17, "Seoul"},
				{24, "Incheon"},
			},
			[]string{"Seoul", "Incheon"},
		},
		{
			"arrays",
			[2]Profile{
				{17, "Seoul"},
				{24, "Incheon"},
			},
			[]string{"Seoul", "Incheon"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{17, "Seoul"}
			aChannel <- Profile{24, "Incheon"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Seoul", "Incheon"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
