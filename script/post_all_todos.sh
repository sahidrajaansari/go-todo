#!/bin/bash

# Define the API endpoint
url="http://localhost:8080/api/v1/todos/"

# Array of JSON payloads (dummy data for various todos)
json_data_1='{
  "title": "Buy Groceries",
  "description": "Include fruits and vegetables",
  "status": "PENDING"
}'

json_data_2='{
  "title": "Complete Homework",
  "description": "Finish math and science assignments",
  "status": "PENDING"
}'

json_data_3='{
  "title": "Meeting with Boss",
  "description": "Discuss quarterly results and goals",
  "status": "PENDING"
}'

json_data_4='{
  "title": "Clean the House",
  "description": "Vacuum and mop the floors, clean the windows",
  "status": "IN_PROGRESS"
}'

json_data_5='{
  "title": "Plan Weekend Trip",
  "description": "Book flights and hotels for weekend getaway",
  "status": "PENDING"
}'

json_data_6='{
  "title": "Write Blog Post",
  "description": "Write a new post on personal development",
  "status": "COMPLETED"
}'

json_data_7='{
  "title": "Grocery Shopping",
  "description": "Buy ingredients for dinner recipes",
  "status": "PENDING"
}'

json_data_8='{
  "title": "Call Mom",
  "description": "Check in and catch up with mom",
  "status": "PENDING"
}'

json_data_9='{
  "title": "Finish Reading Book",
  "description": "Complete the last chapters of the novel",
  "status": "COMPLETED"
}'

json_data_10='{
  "title": "Prepare Presentation",
  "description": "Create slides for upcoming presentation on project progress",
  "status": "IN_PROGRESS"
}'

json_data_11='{
  "title": "Attend Yoga Class",
  "description": "Go to the evening yoga session",
  "status": "COMPLETED"
}'

# Send POST requests with curl for each prototype
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_1"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_2"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_3"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_4"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_5"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_6"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_7"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_8"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_9"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_10"
curl -X POST $url -H "Content-Type: application/json" -d "$json_data_11"

echo "All prototype data sent to $url"
