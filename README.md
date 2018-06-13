# http-dialer

---

http-dialer is an open source golang lib for building http dialer and managing own http connections.

When using standard golang http client, you can hardly control underlying tcp connections.
Http-dialer is for senior golang users to control http underlying connections.

This is copied from [kubernetes](https://github.com/kubernetes/client-go/tree/master/util/connrota)

---

## To start using http-dialer


```
$ go get github.com/hzxuzhonghu/http-dialer
```

Using the http-dialer is very simple: see [example](https://github.com/hzxuzhonghu/http-dialer/example)
