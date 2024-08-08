# mood

This is a Golang web application that allows a team to anonymously vote and get a sense of the "team mood".

The mood is represented by a lissajous gif that will be green and calm for teams in a good mood and red and chaotic for teams in a bad mood.
The mood will normalize over time without any input. 

There will be a `Vote` link on the team view page to go to the vote form.
To vote a user will select one of 5 emojis corresponding to happiness.
This will raise or lower the score and affect the gif when viewed.

## How to setup
1. While in the `mood` code directory, log in to MySQL/MariaDB and create a username/password combination you intend to use for the application.
2. Run `source scripts/mood-mysql.sql`
3. Set the username/password as ENV variables `DATABASE_USERNAME` and `DATABASE_PASSWORD`.

## To Do
- Unique IP cannot vote twice within an hour
- Ability to create a team that requires a password to vote
- Testing
- Unique team names
- Implement search by team name for view
