package main

import "github.com/oxodao/cardsagainstoverflow/model"

func GetAllBoosters() []model.Deck {
	return []model.Deck{
		{
			Title: "Cards Against Overflow",
			Cards: []model.Card{
				{
					Text:        "Le vendredi c'est ____",
					IsBlackCard: true,
				},
				{
					Text:        "____ ça c'est tellement Léo",
					IsBlackCard: true,
				},
				{
					Text:        "Tu sais qui d'autre aimait ____ ? Hitler.",
					IsBlackCard: true,
				},
				{
					Text:        "Être un charo",
					IsBlackCard: false,
				},
				{
					Text:        "Les portugais",
					IsBlackCard: false,
				},
				{
					Text:        "Du tang goût kebab",
					IsBlackCard: false,
				},
			},
		},
		{
			Title: "Blanc Manger Coco",
			Cards: []model.Card{
				{
					Text:        "____, j'y pense tous les jours en me rasant",
					IsBlackCard: true,
				},
				{
					Text:        "Quand j'ai rencontré ta mère, ça a tout de suite été ____",
					IsBlackCard: true,
				},
				{
					Text:        "____, ça c'est tellement ton père",
					IsBlackCard: true,
				},
				{
					Text:        "Un zizi-oreille au petit matin",
					IsBlackCard: false,
				},
				{
					Text:        "Une pièce, une cigarette ou un ticket de restaurant",
					IsBlackCard: false,
				},
				{
					Text:        "Le chantage sexuel",
					IsBlackCard: false,
				},
			},
		},
		{
			Title: "Blanc Manger Coco 2: Le déluge",
			Cards: []model.Card{
				{
					Text:        "____ avec ____, c'est bien mais pas top",
					IsBlackCard: true,
				},
				{
					Text:        "Fais preuve de charité, donne donc ____ au vieux monsieur",
					IsBlackCard: true,
				},
				{
					Text:        "____ à Aulnay-sous-Bois, c'est risqué",
					IsBlackCard: true,
				},
				{
					Text:        "Une commode de 123cm x 45cm x 76cm",
					IsBlackCard: false,
				},
				{
					Text:        "Un doigt d'honneur",
					IsBlackCard: false,
				},
				{
					Text:        "2 heures de sexe non-stop",
					IsBlackCard: false,
				},
			},
		},
	}
}
