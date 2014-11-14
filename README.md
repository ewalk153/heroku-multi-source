== Ruby and Rails plus Golang code

Proof of Concept to run a Rails app and go programs together in the same buildpack.

Note: the `godep` is particular. You'll need to do some path-mangling to make it happy because it wants your code to live in $GOPATH/src.  If you save this git repo this into $GOPATH/src, you should be all set.

=== Heroku deploy

```bash
heroku create -b https://github.com/ddollar/heroku-buildpack-multi.git --region eu
heroku addons:add heroku-postgresql:hobby-dev
heroku addons:add scheduler:standard
```

Then, you can run the go app via the scheduler by adding: `bin/nav-multi`
