# mood

This is a Golang web application that allows a team to anonymously vote and get a sense of the "team mood".

The mood is represented by a lissajous gif that will be green and calm for teams in a good mood and red and chaotic for teams in a bad mood.
The mood will normalize over time without any input. 

To vote a user will select one of 5 emojis corresponding to happiness.
This will raise or lower the score and affect the gif when viewed.

Searching with the `/team` endpoint by name will redirect to the team's view page.
There will be a `Vote` link to go to the vote form.
