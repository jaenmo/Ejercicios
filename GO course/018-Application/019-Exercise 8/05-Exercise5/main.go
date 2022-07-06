package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user
	func (a ByAge) Len()int				{ return len(a) }
	func (a ByAge) Swap(i, j int)		{ a[i], a[j] = a[j], a[i] }
	func (a ByAge) Less(i, j int) bool	{ return a[i].Age < a[j].Age }

type ByName []user
	func (n ByName) Len()int				{ return len(n) }
	func (n ByName) Swap(i, j int)		{ n[i], n[j] = n[j], n[i] }
	func (n ByName) Less(i, j int) bool	{ return n[i].First < n[j].First }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	sort.Sort(ByAge(users))
	for _, u := range users{
		fmt.Println("\n",u.First, u.Age)
		for _, say := range u.Sayings{
		fmt.Println("\t\t",say)
		}
	}
	fmt.Println("--------------")

	
	for _, u := range users{
		fmt.Println("\n",u.First, u.Age)
		sort.Strings(u.Sayings)
		for _, say := range u.Sayings{
		fmt.Println("\t\t",say)
		}
	}
	sort.Sort(ByName(users))
	for _, u := range users{
		fmt.Println("\n",u.First, u.Age)
		for _, say := range u.Sayings{
		fmt.Println("\t\t",say)
		}
	}


}
