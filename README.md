# Blurt

## A micro-tweeting platform.

This is a satirical website based on a Tom Scott video. He was going over some prior predictions of the future made tennish years ago, and he had thought that blogs were going to be dominant in the present. He made note that arguably Twitter and Tweets were a more commonly used medium.

So I have decided to extrapolate from there. So ten years ago you would write 2800 characters, today you write 280 characters, so in the future you will have a maximum of 14 characters to express yourself. Or more aptly, to _blurt_ out what you have to say.

## Security

The audience for this project is purposely being kept small, since a decent-sized userbase would be chaos. I've purposely ignored security, so that anyone can login as anyone. Yes, feel free to log in as Blurt CTO. Or Tom Scott. Once logged in, you should be pretty pleased with yourself for also being verified! Everyone is verified so no one feels left out.

So yes, you can also log in as 800 different people and have flame wars with yourself. Have fun.

My rule is don't be terrible. I expect some nonsense, but remember there will be other humans reading it. Maybe.

## Lik?

Instead of like. Artificially keeping the production value low.

## Coming Soon

Instructions on how to host your own Blurt server!

If you're fairly savvy, you can probably do this already! Clone/fork this repository, run `npm install` in the new directory, look into how to initiate a `prisma.io` sqlite database and you're ready to `npm run dev`!

## Stack

The [actual Blurt](https://letsblurt.duckdns.org) is running in my basement behind an nginx load-balancer/reverse proxy. The front and back ends are currently written in Javascript/Sveltekit with plans to migrate it to Typescript at some point. Prettiness thanks to TailwindCSS. Using and quite enjoying [prisma.io](https://prisma.io) as an ORM which interfaces with an SQLite database, since I don't really plan to have enough people use it to bother with anything bigger.

I do plan to write tests someday.
