install-air:
	go install github.com/air-verse/air@latest

# Run simulations
simulate:
	go run . --simulate $(ARGS)

# Simulation to decide how to play
guess:
	go run . --guess $(ARGS)

# Print all bets for a given number
bets:
	go run . --bets $(ARGS)

# Build the roulette map
build-map:
	go run . --build-map $(ARGS)

# Play the game with the robot
run:
	go run . --play-auto $(ARGS)

# Print the mouse position
mouse:
	go run . --mouse $(ARGS)

# Test capturing the drawn number
number:
	go run . --number $(ARGS)

# Find duplicated bets
duplicates:
	go run . --find-duplicated-bets $(ARGS)

# Print the roulette with the number areas
print-map:
	go run . --print-map $(ARGS)
