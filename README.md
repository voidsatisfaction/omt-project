# OMT(Oh My Teacher) project

![qr code](./qr.png)

You can san this qr code at Line app to add "Hiyoko teacher"

![hiyoko sensei profile](./hiyoko_sensei_profile.png)

This is line bot named "Hiyoko teacher" who is specialist of word exams

## Motivation

I am currently actively participating a internet english speaking class. Every day, my teacher types me a lot of words or sentences used in class.

As I am lazy man, I usually study words only once after finishing class therefore I could not absorb these words and sentences efficiently.

So I made this app to solve word memorizing problem.

It is funny to make and even solve my problem. Wow.

## Technology Stack

1. echo(golang)
2. Line Messaging API
3. AWS S3
4. Heroku
5. Travis CI

## Developement

1. Make sure install docker && Set environment variables(Line bot keys)
2. clone it to `$GOPATH/src/`
3. `docker-compose up`
4. enter `localhost:19000`
5. If `Worked` shown, server is running

Hot recompiling(gin) works. So, after change server-side codes, just reload it on web browser.
