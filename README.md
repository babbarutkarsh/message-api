# message-api
A golang messaging api

Message data structure:
{
“account_id”: “<id>”
“message_id”: “<random-uuid>”,
“sender_number”: “<PHONE_NUMBER>”,
“receiver_number”: “<PHONE_NUMBER>”
}

● /get/messages/<account_id> → Returns all the messages with the sender and
receiver details pertaining to the provided account id.
● /create → Post call which saves the message with the sender and receiver
details
● /search → Search for keys using the following filters
○ Assume you have keys: message_id, sender_number, receiver_number
○ /search?message_id=”1,2” would return messages with the given
message ids.
○ /search?sender_number=”1,2” would return messages with the given
sender_number - can single sender number or multiple sender numbers
under a given account ID.
○ /search?receiver_number=”1,2” would return messages with the given
receiver_number - can single sender number or multiple receiver
numbers


Quality and coverages
● Test coverage 


● Dockerfile with security standards


● Error handling and logging


● Databases MySQL/PostgreSQL



