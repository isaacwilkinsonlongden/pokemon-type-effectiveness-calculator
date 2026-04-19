package main

const (
	reset = "\033[0m"
	bold  = "\033[1m"
)

func styleBold(text string) string {
	return bold + text + reset
}

func colorForType(typeName string) string {
	switch typeName {
	case "normal":
		return "\033[37m"
	case "fire":
		return "\033[31m"
	case "water":
		return "\033[34m"
	case "electric":
		return "\033[33m"
	case "grass":
		return "\033[32m"
	case "ice":
		return "\033[36m"
	case "fighting":
		return "\033[31m"
	case "poison":
		return "\033[35m"
	case "ground":
		return "\033[33m"
	case "flying":
		return "\033[36m"
	case "psychic":
		return "\033[35m"
	case "bug":
		return "\033[32m"
	case "rock":
		return "\033[33m"
	case "ghost":
		return "\033[35m"
	case "dragon":
		return "\033[34m"
	case "dark":
		return "\033[90m"
	case "steel":
		return "\033[37m"
	case "fairy":
		return "\033[35m"
	default:
		return ""
	}
}

func colorizeType(typeName string) string {
	color := colorForType(typeName)
	if color == "" {
		return typeName
	}
	return color + typeName + reset
}
