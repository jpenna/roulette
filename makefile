install-air:
	go install github.com/air-verse/air@latest

# Run simulations
simulate:
	air -- --simulate

# Play the game in the terminal
terminal:
	go run . --play-terminal

# Print all bets for a given number
bets:
	go run . --bets

# Build the roulette map
build-map:
	go run . --build-map

# Play the game with the robot
auto:
	go run . --play-auto

# Print the mouse position
mouse:
	go run . --mouse

# Test capturing the drawn number
number:
	go run . --number

# Find duplicated bets
duplicates:
	go run . --find-duplicated-bets
