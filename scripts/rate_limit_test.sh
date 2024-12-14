#!/bin/bash

# URL to send POST requests to
URL="http://localhost:8080/orders"

# Number of parallel requests
NUM_REQUESTS=10

# Optional: Data to send with each POST request
# Replace with actual data as needed, e.g., JSON payload

# Optional: Headers, e.g., Content-Type
HEADERS=(
  -H "Content-Type: application/json"
)

echo "Sending $NUM_REQUESTS parallel POST requests to $URL..."

# Function to send a single POST request
send_post_request() {
  curl -X POST "${HEADERS[@]}" "$URL"
}

# Export the function and variables for subshells
export -f send_post_request
export URL POST_DATA HEADERS

# Launch NUM_REQUESTS background processes
for i in $(seq 1 $NUM_REQUESTS); do
  send_post_request &
done

# Wait for all background processes to finish
wait

echo "All $NUM_REQUESTS POST requests have been sent."
