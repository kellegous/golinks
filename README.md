# GoLinks

## Status

This is not even closed to being finished, move along for now.

## What is this?

Many years ago, I threw together an extremely simple go-link server, [kellegous/go](https://github.com/kellegous/go/). To my surprise, a number of people used it and we used it for almost 9 years at Mailchimp. I always had one big regret about the hurried nature of that implementation. I wanted to be able to create go links that had some form of replacement. As an example, many companies have some kind of staff directory that is widely used and, thus, warrants a go-link, maybe `go/staff` that redirects to `https://staff.company.com/`. But what if I wanted to be able to give someone a go-link that goes directly to a specific person's profile? Sure would be nice if the go-link server could handle that. For instance, `go/staff/knorton` might redirect to `https://staff.company.com/person/knorton`. This is one of the primary reasons this project exists. Also, I made a horrible mistake naming the previous project **"go"** because the compiled binary is named **go** just like a much different command that is part of the Go Programming Language.
