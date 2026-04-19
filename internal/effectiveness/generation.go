package effectiveness

type Generation string

const (
	Current        Generation = ""
	GenerationI    Generation = "generation-i"
	GenerationII   Generation = "generation-ii"
	GenerationIII  Generation = "generation-iii"
	GenerationIV   Generation = "generation-iv"
	GenerationV    Generation = "generation-v"
	GenerationVI   Generation = "generation-vi"
	GenerationVII  Generation = "generation-vii"
	GenerationVIII Generation = "generation-viii"
	GenerationIX   Generation = "generation-ix"
)

func generationNumber(gen Generation) int {
	switch gen {
	case GenerationI:
		return 1
	case GenerationII:
		return 2
	case GenerationIII:
		return 3
	case GenerationIV:
		return 4
	case GenerationV:
		return 5
	case GenerationVI:
		return 6
	case GenerationVII:
		return 7
	case GenerationVIII:
		return 8
	case GenerationIX, Current:
		return 9
	default:
		return 0
	}
}
