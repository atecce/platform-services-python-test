# Objective
* Create RESTful endpoint(s) to expose Rewards Data.
* Create an algorithm to calculate a user's rewards points.
* Create a Customer Rewards Dashboard Application.

# Solution
This solution has deviated a bit from the instructions given in the _how_ but for the most part not the _what_. `svc` contains a Go service which implements all the business requirements for the datastore and `ui` contains a Vue app which displays a very simple dashboard.

I hope it's okay that I got a little creative with this, and I appreciate the interesting exercise! I had a lot of fun with it. Kimberly mentioned both Go and Vue as technologies you're exploring and that the team you're recruiting for has an opportunity to think about the software architecture from first principles.

She mentioned Monday as a deadline, so there are definitely a few loose ends I'd want to tie up with more time (in particular, getting docker compose to properly orchestrate the `ui` with the `svc`, which I stalled out on, and therefore am simply running on the host). I had a bunch of stuff over the weekend and would totally welcome an extension to clean some of it up, as this is mostly the result of one long hacking session to get a proof of concept up.

Otherwise I again appreciate the opportunity to solve an interesting problem, as it's still primarily the thing that motivates me to work, it's the kind of thing you don't have to pay me for, I'd be doing it anyway.

# Dependencies
### mongo 
```sh
brew install mongo
brew services start mongo
```
### go and dep
```sh
brew install dep
```
### yarn
```sh
brew install yarn
```