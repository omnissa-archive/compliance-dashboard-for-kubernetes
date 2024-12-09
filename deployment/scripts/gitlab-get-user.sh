#!/bin/bash

source .env

ACCESS_TOKEN=$(source auth-gitlab.sh | jq -r '.access_token')

echo $ACCESS_TOKEN

curl -s -H "Authorization: Bearer $ACCESS_TOKEN" https://gitlab.eng.omnissa.com/api/v4/user

