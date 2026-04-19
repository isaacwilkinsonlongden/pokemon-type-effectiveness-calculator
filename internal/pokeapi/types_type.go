package pokeapi

type Type struct {
	DamageRelations struct {
		DoubleDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"double_damage_from"`
		DoubleDamageTo []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"double_damage_to"`
		HalfDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"half_damage_from"`
		HalfDamageTo []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"half_damage_to"`
		NoDamageFrom []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"no_damage_from"`
		NoDamageTo []any `json:"no_damage_to"`
	} `json:"damage_relations"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	ID              int `json:"id"`
	MoveDamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_damage_class"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PastDamageRelations []struct {
		DamageRelations struct {
			DoubleDamageFrom []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"double_damage_from"`
			DoubleDamageTo []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"double_damage_to"`
			HalfDamageFrom []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"half_damage_from"`
			HalfDamageTo []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"half_damage_to"`
			NoDamageFrom []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"no_damage_from"`
			NoDamageTo []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"no_damage_to"`
		} `json:"damage_relations"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"past_damage_relations"`
	Pokemon []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		Slot int `json:"slot"`
	} `json:"pokemon"`
	Sprites struct {
		GenerationIii struct {
			Colosseum struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"colosseum"`
			Emerald struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"emerald"`
			FireredLeafgreen struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"firered-leafgreen"`
			RubySapphire struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"ruby-sapphire"`
			Xd struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"xd"`
		} `json:"generation-iii"`
		GenerationIv struct {
			DiamondPearl struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"diamond-pearl"`
			HeartgoldSoulsilver struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"heartgold-soulsilver"`
			Platinum struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"platinum"`
		} `json:"generation-iv"`
		GenerationIx struct {
			ScarletViolet struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon string `json:"symbol_icon"`
			} `json:"scarlet-violet"`
		} `json:"generation-ix"`
		GenerationV struct {
			Black2White2 struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"black-2-white-2"`
			BlackWhite struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"black-white"`
		} `json:"generation-v"`
		GenerationVi struct {
			OmegaRubyAlphaSapphire struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"omega-ruby-alpha-sapphire"`
			XY struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"x-y"`
		} `json:"generation-vi"`
		GenerationVii struct {
			LetsGoPikachuLetsGoEevee struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon string `json:"symbol_icon"`
			} `json:"lets-go-pikachu-lets-go-eevee"`
			SunMoon struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"sun-moon"`
			UltraSunUltraMoon struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon any    `json:"symbol_icon"`
			} `json:"ultra-sun-ultra-moon"`
		} `json:"generation-vii"`
		GenerationViii struct {
			BrilliantDiamondShiningPearl struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon string `json:"symbol_icon"`
			} `json:"brilliant-diamond-shining-pearl"`
			LegendsArceus struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon string `json:"symbol_icon"`
			} `json:"legends-arceus"`
			SwordShield struct {
				NameIcon   string `json:"name_icon"`
				SymbolIcon string `json:"symbol_icon"`
			} `json:"sword-shield"`
		} `json:"generation-viii"`
	} `json:"sprites"`
}
